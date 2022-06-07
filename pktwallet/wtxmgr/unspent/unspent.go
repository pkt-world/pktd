package unspent

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/pktlog/log"
	"github.com/pkt-cash/pktd/pktwallet/walletdb"
	"github.com/pkt-cash/pktd/pktwallet/wtxmgr/dbstructs"
	"github.com/pkt-cash/pktd/pktwallet/wtxmgr/utilfun"
	"github.com/pkt-cash/pktd/wire"
)

var bucketUnspentOld = []byte("u")
var bucketUnspent = []byte("u2")

var UnspentErr = er.GenericErrorType.Code("UnspentErr")

// The unspent index records all outpoints for mined credits which are not spent
// by any other mined transaction records (but may be spent by a mempool
// transaction).
//
// Keys are use the canonical outpoint serialization:
//
//   [0:32]  Transaction hash (32 bytes)
//   [32:36] Output index (4 bytes)
//
// Values WERE serialized as such, now are json.
//
//   [0:4]   Block height (4 bytes)
//   [4:36]  Block hash (32 bytes)

func Put(ns walletdb.ReadWriteBucket, u *dbstructs.Unspent) er.R {
	k := utilfun.CanonicalOutPoint(&u.OutPoint.Hash, u.OutPoint.Index)
	if v, err := jsoniter.Marshal(u); err != nil {
		return UnspentErr.New("cannot marshal unspent", er.E(err))
	} else if err := ns.NestedReadWriteBucket(bucketUnspent).Put(k, v); err != nil {
		return UnspentErr.New("cannot put unspent", err)
	}
	return nil
}

// existsUnspent returns the key for the unspent output and the corresponding
// key for the credits bucket.  If there is no unspent output recorded, the
// credit key is nil.
func Exists(ns walletdb.ReadBucket, outPoint *wire.OutPoint) (k, credKey []byte) {
	k = utilfun.CanonicalOutPoint(&outPoint.Hash, outPoint.Index)
	credKey = ExistsRaw(ns, k)
	return k, credKey
}

// existsRawUnspent returns the credit key if there exists an output recorded
// for the raw unspent key.  It returns nil if the k/v pair does not exist.
func ExistsRaw(ns walletdb.ReadBucket, k []byte) (credKey []byte) {
	if len(k) < 36 {
		return nil
	}
	v := ns.NestedReadBucket(bucketUnspent).Get(k)
	if len(v) < 36 {
		return nil
	}
	credKey = make([]byte, 72)
	copy(credKey, k[:32])
	copy(credKey[32:68], v)
	copy(credKey[68:72], k[32:36])
	return credKey
}

func DeleteRaw(ns walletdb.ReadWriteBucket, k []byte) er.R {
	err := ns.NestedReadWriteBucket(bucketUnspent).Delete(k)
	if err != nil {
		return UnspentErr.New("failed to delete unspent", err)
	}
	return nil
}

func ForEachUnspentOutput(
	ns walletdb.ReadBucket,
	beginKey []byte,
	visitor func(key []byte, unspent *dbstructs.Unspent) er.R,
) er.R {
	bu := ns.NestedReadBucket(bucketUnspent)
	return bu.ForEachBeginningWith(beginKey, func(k, v []byte) er.R {
		var unspent dbstructs.Unspent
		if err := jsoniter.Unmarshal(v, &unspent); err != nil {
			return UnspentErr.New(fmt.Sprintf("Unable to parse JSON [%s]", string(v)), er.E(err))
		}
		if err := visitor(k, &unspent); err != nil {
			return err
		}
		return nil
	})
}

func CreateBuckets(ns walletdb.ReadWriteBucket) er.R {
	_, err := ns.CreateBucket(bucketUnspent)
	return err
}
func DeleteBuckets(ns walletdb.ReadWriteBucket) er.R {
	return ns.DeleteNestedBucket(bucketUnspent)
}

func ExtendUnspents(ns walletdb.ReadWriteBucket, extend func(u *dbstructs.Unspent) er.R) er.R {
	if _, err := ns.CreateBucket(bucketUnspent); err != nil {
		return err
	}
	bu := ns.NestedReadBucket(bucketUnspentOld)
	if bucketUnspentOld == nil {
		log.Warn("There is no bucketUnspentOld bucket")
	}
	i := 0
	if err := bu.ForEach(func(k, v []byte) er.R {
		var unspent dbstructs.Unspent
		if err := utilfun.ReadCanonicalOutPoint(k, &unspent.OutPoint); err != nil {
			return err
		}
		if err := utilfun.ReadUnspentBlock(v, &unspent.Block); err != nil {
			return err
		}
		if err := extend(&unspent); err != nil {
			return err
		}
		if err := Put(ns, &unspent); err != nil {
			return err
		}
		if i%1000 == 0 {
			log.Infof("Migrating UTXO [%d] ([%d]%%)", i, int(k[0])*100/256)
		}
		i += 1
		return nil
	}); err != nil {
		return err
	}
	log.Info("Deleting old UTXO bucket")
	return ns.DeleteNestedBucket(bucketUnspentOld)
}
