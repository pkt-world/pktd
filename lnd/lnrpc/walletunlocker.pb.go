// Code generated by protoc-gen-go. DO NOT EDIT.
// source: walletunlocker.proto

package lnrpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GenSeedRequest struct {
	//
	//seed_passphrase is the optional user specified passphrase that will be used
	//to encrypt the generated seed.
	SeedPassphrase string `protobuf:"bytes,1,opt,name=seed_passphrase,json=seedPassphrase,proto3" json:"seed_passphrase,omitempty"`
	//
	//seed_passphrase_bin overrides seed_passphrase if specified, for binary
	//representation of the passphrase. If using JSON then this field must be base64
	//encoded.
	SeedPassphraseBin []byte `protobuf:"bytes,2,opt,name=seed_passphrase_bin,json=seedPassphraseBin,proto3" json:"seed_passphrase_bin,omitempty"`
	//
	//seed_entropy is an optional 16-bytes generated via CSPRNG. If not
	//specified, then a fresh set of randomness will be used to create the seed.
	//When using REST, this field must be encoded as base64.
	SeedEntropy          []byte   `protobuf:"bytes,3,opt,name=seed_entropy,json=seedEntropy,proto3" json:"seed_entropy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenSeedRequest) Reset()         { *m = GenSeedRequest{} }
func (m *GenSeedRequest) String() string { return proto.CompactTextString(m) }
func (*GenSeedRequest) ProtoMessage()    {}
func (*GenSeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e3ed10ed53e4fd, []int{0}
}

func (m *GenSeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenSeedRequest.Unmarshal(m, b)
}
func (m *GenSeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenSeedRequest.Marshal(b, m, deterministic)
}
func (m *GenSeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenSeedRequest.Merge(m, src)
}
func (m *GenSeedRequest) XXX_Size() int {
	return xxx_messageInfo_GenSeedRequest.Size(m)
}
func (m *GenSeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenSeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenSeedRequest proto.InternalMessageInfo

func (m *GenSeedRequest) GetSeedPassphrase() string {
	if m != nil {
		return m.SeedPassphrase
	}
	return ""
}

func (m *GenSeedRequest) GetSeedPassphraseBin() []byte {
	if m != nil {
		return m.SeedPassphraseBin
	}
	return nil
}

func (m *GenSeedRequest) GetSeedEntropy() []byte {
	if m != nil {
		return m.SeedEntropy
	}
	return nil
}

type GenSeedResponse struct {
	//
	//seed is a 15-word mnemonic that encodes a secret root seed used to generate
	//all private keys of the wallet.
	Seed                 []string `protobuf:"bytes,1,rep,name=seed,proto3" json:"seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenSeedResponse) Reset()         { *m = GenSeedResponse{} }
func (m *GenSeedResponse) String() string { return proto.CompactTextString(m) }
func (*GenSeedResponse) ProtoMessage()    {}
func (*GenSeedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e3ed10ed53e4fd, []int{1}
}

func (m *GenSeedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenSeedResponse.Unmarshal(m, b)
}
func (m *GenSeedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenSeedResponse.Marshal(b, m, deterministic)
}
func (m *GenSeedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenSeedResponse.Merge(m, src)
}
func (m *GenSeedResponse) XXX_Size() int {
	return xxx_messageInfo_GenSeedResponse.Size(m)
}
func (m *GenSeedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenSeedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenSeedResponse proto.InternalMessageInfo

func (m *GenSeedResponse) GetSeed() []string {
	if m != nil {
		return m.Seed
	}
	return nil
}

type InitWalletRequest struct {
	//
	//wallet_passphrase is the passphrase that should be used to encrypt the
	//wallet. This MUST be at least 8 chars in length. After creation, this
	//password is required to unlock the daemon.
	WalletPassphrase string `protobuf:"bytes,1,opt,name=wallet_passphrase,json=walletPassphrase,proto3" json:"wallet_passphrase,omitempty"`
	//
	//If specified, will override wallet_passphrase, but is expressed in binary.
	//When using REST, this field must be encoded as base64.
	WalletPassphraseBin []byte `protobuf:"bytes,2,opt,name=wallet_passphrase_bin,json=walletPassphraseBin,proto3" json:"wallet_passphrase_bin,omitempty"`
	//
	//wallet_seed is a 15-word wallet seed. This may have been generated by the
	//GenSeed method, or be an existing seed.
	WalletSeed []string `protobuf:"bytes,3,rep,name=wallet_seed,json=walletSeed,proto3" json:"wallet_seed,omitempty"`
	//
	//seed_passphrase is an optional user provided passphrase that will be used
	//to encrypt the generated seed.
	SeedPassphrase string `protobuf:"bytes,4,opt,name=seed_passphrase,json=seedPassphrase,proto3" json:"seed_passphrase,omitempty"`
	//
	//If specified, will override seed_passphrase, but is expressed in binary.
	//When using REST, this field must be encoded as base64.
	SeedPassphraseBin []byte `protobuf:"bytes,5,opt,name=seed_passphrase_bin,json=seedPassphraseBin,proto3" json:"seed_passphrase_bin,omitempty"`
	//
	//recovery_window is an optional argument specifying the address lookahead
	//when restoring a wallet seed. The recovery window applies to each
	//individual branch of the BIP44 derivation paths. Supplying a recovery
	//window of zero indicates that no addresses should be recovered, such after
	//the first initialization of the wallet.
	RecoveryWindow int32 `protobuf:"varint,6,opt,name=recovery_window,json=recoveryWindow,proto3" json:"recovery_window,omitempty"`
	//
	//channel_backups is an optional argument that allows clients to recover the
	//settled funds within a set of channels. This should be populated if the
	//user was unable to close out all channels and sweep funds before partial or
	//total data loss occurred. If specified, then after on-chain recovery of
	//funds, lnd begin to carry out the data loss recovery protocol in order to
	//recover the funds in each channel from a remote force closed transaction.
	ChannelBackups *ChanBackupSnapshot `protobuf:"bytes,7,opt,name=channel_backups,json=channelBackups,proto3" json:"channel_backups,omitempty"`
	//
	//wallet_name is an optional argument that allows to define the
	//wallet filename other than the default wallet.db
	WalletName           string   `protobuf:"bytes,8,opt,name=wallet_name,json=walletName,proto3" json:"wallet_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitWalletRequest) Reset()         { *m = InitWalletRequest{} }
func (m *InitWalletRequest) String() string { return proto.CompactTextString(m) }
func (*InitWalletRequest) ProtoMessage()    {}
func (*InitWalletRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e3ed10ed53e4fd, []int{2}
}

func (m *InitWalletRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitWalletRequest.Unmarshal(m, b)
}
func (m *InitWalletRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitWalletRequest.Marshal(b, m, deterministic)
}
func (m *InitWalletRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitWalletRequest.Merge(m, src)
}
func (m *InitWalletRequest) XXX_Size() int {
	return xxx_messageInfo_InitWalletRequest.Size(m)
}
func (m *InitWalletRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitWalletRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitWalletRequest proto.InternalMessageInfo

func (m *InitWalletRequest) GetWalletPassphrase() string {
	if m != nil {
		return m.WalletPassphrase
	}
	return ""
}

func (m *InitWalletRequest) GetWalletPassphraseBin() []byte {
	if m != nil {
		return m.WalletPassphraseBin
	}
	return nil
}

func (m *InitWalletRequest) GetWalletSeed() []string {
	if m != nil {
		return m.WalletSeed
	}
	return nil
}

func (m *InitWalletRequest) GetSeedPassphrase() string {
	if m != nil {
		return m.SeedPassphrase
	}
	return ""
}

func (m *InitWalletRequest) GetSeedPassphraseBin() []byte {
	if m != nil {
		return m.SeedPassphraseBin
	}
	return nil
}

func (m *InitWalletRequest) GetRecoveryWindow() int32 {
	if m != nil {
		return m.RecoveryWindow
	}
	return 0
}

func (m *InitWalletRequest) GetChannelBackups() *ChanBackupSnapshot {
	if m != nil {
		return m.ChannelBackups
	}
	return nil
}

func (m *InitWalletRequest) GetWalletName() string {
	if m != nil {
		return m.WalletName
	}
	return ""
}

type InitWalletResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitWalletResponse) Reset()         { *m = InitWalletResponse{} }
func (m *InitWalletResponse) String() string { return proto.CompactTextString(m) }
func (*InitWalletResponse) ProtoMessage()    {}
func (*InitWalletResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e3ed10ed53e4fd, []int{3}
}

func (m *InitWalletResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitWalletResponse.Unmarshal(m, b)
}
func (m *InitWalletResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitWalletResponse.Marshal(b, m, deterministic)
}
func (m *InitWalletResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitWalletResponse.Merge(m, src)
}
func (m *InitWalletResponse) XXX_Size() int {
	return xxx_messageInfo_InitWalletResponse.Size(m)
}
func (m *InitWalletResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitWalletResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitWalletResponse proto.InternalMessageInfo

type UnlockWalletRequest struct {
	//
	//wallet_passphrase should be the current valid private passphrase for the daemon. This
	//will be required to decrypt on-disk material that the daemon requires to
	//function properly.
	WalletPassphrase string `protobuf:"bytes,1,opt,name=wallet_passphrase,json=walletPassphrase,proto3" json:"wallet_passphrase,omitempty"`
	//
	//If specified, will override wallet_passphrase, but is expressed in binary.
	//When using REST, this field must be encoded as base64.
	WalletPassphraseBin []byte `protobuf:"bytes,2,opt,name=wallet_passphrase_bin,json=walletPassphraseBin,proto3" json:"wallet_passphrase_bin,omitempty"`
	//
	//recovery_window is an optional argument specifying the address lookahead
	//when restoring a wallet seed. The recovery window applies to each
	//individual branch of the BIP44 derivation paths. Supplying a recovery
	//window of zero indicates that no addresses should be recovered, such after
	//the first initialization of the wallet.
	RecoveryWindow int32 `protobuf:"varint,3,opt,name=recovery_window,json=recoveryWindow,proto3" json:"recovery_window,omitempty"`
	//
	//channel_backups is an optional argument that allows clients to recover the
	//settled funds within a set of channels. This should be populated if the
	//user was unable to close out all channels and sweep funds before partial or
	//total data loss occurred. If specified, then after on-chain recovery of
	//funds, lnd begin to carry out the data loss recovery protocol in order to
	//recover the funds in each channel from a remote force closed transaction.
	ChannelBackups *ChanBackupSnapshot `protobuf:"bytes,4,opt,name=channel_backups,json=channelBackups,proto3" json:"channel_backups,omitempty"`
	//
	//wallet_name is optional when the user wants to load a specified wallet other
	//than the default wallet.db
	WalletName           string   `protobuf:"bytes,5,opt,name=wallet_name,json=walletName,proto3" json:"wallet_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnlockWalletRequest) Reset()         { *m = UnlockWalletRequest{} }
func (m *UnlockWalletRequest) String() string { return proto.CompactTextString(m) }
func (*UnlockWalletRequest) ProtoMessage()    {}
func (*UnlockWalletRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e3ed10ed53e4fd, []int{4}
}

func (m *UnlockWalletRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnlockWalletRequest.Unmarshal(m, b)
}
func (m *UnlockWalletRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnlockWalletRequest.Marshal(b, m, deterministic)
}
func (m *UnlockWalletRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnlockWalletRequest.Merge(m, src)
}
func (m *UnlockWalletRequest) XXX_Size() int {
	return xxx_messageInfo_UnlockWalletRequest.Size(m)
}
func (m *UnlockWalletRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnlockWalletRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnlockWalletRequest proto.InternalMessageInfo

func (m *UnlockWalletRequest) GetWalletPassphrase() string {
	if m != nil {
		return m.WalletPassphrase
	}
	return ""
}

func (m *UnlockWalletRequest) GetWalletPassphraseBin() []byte {
	if m != nil {
		return m.WalletPassphraseBin
	}
	return nil
}

func (m *UnlockWalletRequest) GetRecoveryWindow() int32 {
	if m != nil {
		return m.RecoveryWindow
	}
	return 0
}

func (m *UnlockWalletRequest) GetChannelBackups() *ChanBackupSnapshot {
	if m != nil {
		return m.ChannelBackups
	}
	return nil
}

func (m *UnlockWalletRequest) GetWalletName() string {
	if m != nil {
		return m.WalletName
	}
	return ""
}

type UnlockWalletResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnlockWalletResponse) Reset()         { *m = UnlockWalletResponse{} }
func (m *UnlockWalletResponse) String() string { return proto.CompactTextString(m) }
func (*UnlockWalletResponse) ProtoMessage()    {}
func (*UnlockWalletResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e3ed10ed53e4fd, []int{5}
}

func (m *UnlockWalletResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnlockWalletResponse.Unmarshal(m, b)
}
func (m *UnlockWalletResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnlockWalletResponse.Marshal(b, m, deterministic)
}
func (m *UnlockWalletResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnlockWalletResponse.Merge(m, src)
}
func (m *UnlockWalletResponse) XXX_Size() int {
	return xxx_messageInfo_UnlockWalletResponse.Size(m)
}
func (m *UnlockWalletResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnlockWalletResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnlockWalletResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenSeedRequest)(nil), "lnrpc.GenSeedRequest")
	proto.RegisterType((*GenSeedResponse)(nil), "lnrpc.GenSeedResponse")
	proto.RegisterType((*InitWalletRequest)(nil), "lnrpc.InitWalletRequest")
	proto.RegisterType((*InitWalletResponse)(nil), "lnrpc.InitWalletResponse")
	proto.RegisterType((*UnlockWalletRequest)(nil), "lnrpc.UnlockWalletRequest")
	proto.RegisterType((*UnlockWalletResponse)(nil), "lnrpc.UnlockWalletResponse")
}

func init() { proto.RegisterFile("walletunlocker.proto", fileDescriptor_76e3ed10ed53e4fd) }

var fileDescriptor_76e3ed10ed53e4fd = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0x26, 0xfd, 0xb1, 0x6b, 0x5f, 0x4b, 0x6a, 0xa7, 0xdd, 0x25, 0x1b, 0x0f, 0xc6, 0xa0, 0x6c,
	0x40, 0xcc, 0x42, 0xbd, 0x78, 0xb5, 0x22, 0x8b, 0x17, 0x91, 0x2c, 0xcb, 0x82, 0x97, 0x32, 0x4d,
	0x1f, 0x26, 0x34, 0x9d, 0x19, 0x67, 0x52, 0xcb, 0xde, 0x3d, 0xfa, 0xf7, 0xf9, 0xd7, 0x78, 0x90,
	0xcc, 0xcc, 0xae, 0x69, 0x1b, 0x41, 0x11, 0x3c, 0x04, 0x86, 0xef, 0x7d, 0x6f, 0xf8, 0xbe, 0xef,
	0xcd, 0x0b, 0x4c, 0xb6, 0xb4, 0x28, 0xb0, 0xdc, 0xb0, 0x82, 0xa7, 0x2b, 0x94, 0xb1, 0x90, 0xbc,
	0xe4, 0xa4, 0x5b, 0x30, 0x29, 0x52, 0xbf, 0x27, 0x45, 0x6a, 0x90, 0xf0, 0x9b, 0x03, 0xee, 0x25,
	0xb2, 0x2b, 0xc4, 0x65, 0x82, 0x9f, 0x37, 0xa8, 0x4a, 0x72, 0x0e, 0x43, 0x85, 0xb8, 0x9c, 0x0b,
	0xaa, 0x94, 0xc8, 0x24, 0x55, 0xe8, 0x39, 0x81, 0x13, 0xf5, 0x12, 0xb7, 0x82, 0x3f, 0xdc, 0xa3,
	0x24, 0x86, 0xf1, 0x1e, 0x71, 0xbe, 0xc8, 0x99, 0xd7, 0x0a, 0x9c, 0x68, 0x90, 0x8c, 0x76, 0xc9,
	0xb3, 0x9c, 0x91, 0x27, 0x30, 0xd0, 0x7c, 0x64, 0xa5, 0xe4, 0xe2, 0xd6, 0x6b, 0x6b, 0x62, 0xbf,
	0xc2, 0xde, 0x1a, 0x28, 0x7c, 0x06, 0xc3, 0x7b, 0x35, 0x4a, 0x70, 0xa6, 0x90, 0x10, 0xe8, 0x54,
	0x0c, 0xcf, 0x09, 0xda, 0x51, 0x2f, 0xd1, 0xe7, 0xf0, 0x47, 0x0b, 0x46, 0xef, 0x58, 0x5e, 0xde,
	0x68, 0x93, 0x77, 0xc2, 0x9f, 0xc3, 0xc8, 0xb8, 0x3e, 0x94, 0xfe, 0xd0, 0x14, 0x6a, 0xe2, 0xa7,
	0x70, 0x72, 0x40, 0xae, 0xc9, 0x1f, 0xef, 0x37, 0x54, 0x06, 0x1e, 0x43, 0xdf, 0xf6, 0x68, 0x45,
	0x6d, 0xad, 0x08, 0x0c, 0x54, 0x69, 0x6e, 0x8a, 0xae, 0xf3, 0x37, 0xd1, 0x75, 0x7f, 0x17, 0xdd,
	0x39, 0x0c, 0x25, 0xa6, 0xfc, 0x0b, 0xca, 0xdb, 0xf9, 0x36, 0x67, 0x4b, 0xbe, 0xf5, 0x8e, 0x02,
	0x27, 0xea, 0x26, 0xee, 0x1d, 0x7c, 0xa3, 0x51, 0x32, 0x83, 0x61, 0x9a, 0x51, 0xc6, 0xb0, 0x98,
	0x2f, 0x68, 0xba, 0xda, 0x08, 0xe5, 0x1d, 0x07, 0x4e, 0xd4, 0x9f, 0x9e, 0xc5, 0x7a, 0xf6, 0xf1,
	0x9b, 0x8c, 0xb2, 0x99, 0xae, 0x5c, 0x31, 0x2a, 0x54, 0xc6, 0xcb, 0xc4, 0xb5, 0x1d, 0x06, 0x56,
	0x35, 0x9b, 0x8c, 0xae, 0xd1, 0x7b, 0xa0, 0x1d, 0x58, 0x9b, 0xef, 0xe9, 0x1a, 0xc3, 0x09, 0x90,
	0x7a, 0xfa, 0x66, 0x50, 0xe1, 0xd7, 0x16, 0x8c, 0xaf, 0xf5, 0x7b, 0xfb, 0xcf, 0x63, 0x69, 0x08,
	0xa7, 0xfd, 0xa7, 0xe1, 0x74, 0xfe, 0x31, 0x9c, 0xee, 0x41, 0x38, 0xa7, 0x30, 0xd9, 0x4d, 0xc1,
	0xc4, 0x33, 0xfd, 0xee, 0x80, 0x6b, 0xa0, 0x6b, 0xbb, 0x94, 0xe4, 0x15, 0x1c, 0xdb, 0xd7, 0x4e,
	0x4e, 0xac, 0x82, 0xdd, 0x5d, 0xf4, 0x4f, 0xf7, 0x61, 0xbb, 0x14, 0xaf, 0x01, 0x7e, 0x4d, 0x80,
	0x78, 0x96, 0x75, 0xb0, 0x12, 0xfe, 0x59, 0x43, 0xc5, 0x5e, 0x71, 0x09, 0x83, 0xba, 0x4e, 0xe2,
	0x5b, 0x6a, 0xc3, 0x08, 0xfd, 0x47, 0x8d, 0x35, 0x73, 0xd1, 0xec, 0xe9, 0xc7, 0xf0, 0x53, 0x5e,
	0x66, 0x9b, 0x45, 0x9c, 0xf2, 0xf5, 0x85, 0x58, 0x95, 0x2f, 0x52, 0xaa, 0xb2, 0xea, 0xb0, 0xbc,
	0x28, 0x58, 0xf5, 0x49, 0x91, 0x2e, 0x8e, 0xf4, 0xff, 0xe6, 0xe5, 0xcf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfa, 0xc4, 0x92, 0x0a, 0x99, 0x04, 0x00, 0x00,
}
