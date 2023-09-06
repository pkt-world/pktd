// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: chainrpc/chainnotifier.proto

package chainrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChainNotifier_RegisterConfirmationsNtfn_FullMethodName = "/chainrpc.ChainNotifier/RegisterConfirmationsNtfn"
	ChainNotifier_RegisterSpendNtfn_FullMethodName         = "/chainrpc.ChainNotifier/RegisterSpendNtfn"
	ChainNotifier_RegisterBlockEpochNtfn_FullMethodName    = "/chainrpc.ChainNotifier/RegisterBlockEpochNtfn"
)

// ChainNotifierClient is the client API for ChainNotifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChainNotifierClient interface {
	// RegisterConfirmationsNtfn is a synchronous response-streaming RPC that
	// registers an intent for a client to be notified once a confirmation request
	// has reached its required number of confirmations on-chain.
	//
	// A client can specify whether the confirmation request should be for a
	// particular transaction by its hash or for an output script by specifying a
	// zero hash.
	RegisterConfirmationsNtfn(ctx context.Context, in *ConfRequest, opts ...grpc.CallOption) (ChainNotifier_RegisterConfirmationsNtfnClient, error)
	// RegisterSpendNtfn is a synchronous response-streaming RPC that registers an
	// intent for a client to be notification once a spend request has been spent
	// by a transaction that has confirmed on-chain.
	//
	// A client can specify whether the spend request should be for a particular
	// outpoint  or for an output script by specifying a zero outpoint.
	RegisterSpendNtfn(ctx context.Context, in *SpendRequest, opts ...grpc.CallOption) (ChainNotifier_RegisterSpendNtfnClient, error)
	// RegisterBlockEpochNtfn is a synchronous response-streaming RPC that
	// registers an intent for a client to be notified of blocks in the chain. The
	// stream will return a hash and height tuple of a block for each new/stale
	// block in the chain. It is the client's responsibility to determine whether
	// the tuple returned is for a new or stale block in the chain.
	//
	// A client can also request a historical backlog of blocks from a particular
	// point. This allows clients to be idempotent by ensuring that they do not
	// missing processing a single block within the chain.
	RegisterBlockEpochNtfn(ctx context.Context, in *BlockEpoch, opts ...grpc.CallOption) (ChainNotifier_RegisterBlockEpochNtfnClient, error)
}

type chainNotifierClient struct {
	cc grpc.ClientConnInterface
}

func NewChainNotifierClient(cc grpc.ClientConnInterface) ChainNotifierClient {
	return &chainNotifierClient{cc}
}

func (c *chainNotifierClient) RegisterConfirmationsNtfn(ctx context.Context, in *ConfRequest, opts ...grpc.CallOption) (ChainNotifier_RegisterConfirmationsNtfnClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChainNotifier_ServiceDesc.Streams[0], ChainNotifier_RegisterConfirmationsNtfn_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chainNotifierRegisterConfirmationsNtfnClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChainNotifier_RegisterConfirmationsNtfnClient interface {
	Recv() (*ConfEvent, error)
	grpc.ClientStream
}

type chainNotifierRegisterConfirmationsNtfnClient struct {
	grpc.ClientStream
}

func (x *chainNotifierRegisterConfirmationsNtfnClient) Recv() (*ConfEvent, error) {
	m := new(ConfEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chainNotifierClient) RegisterSpendNtfn(ctx context.Context, in *SpendRequest, opts ...grpc.CallOption) (ChainNotifier_RegisterSpendNtfnClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChainNotifier_ServiceDesc.Streams[1], ChainNotifier_RegisterSpendNtfn_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chainNotifierRegisterSpendNtfnClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChainNotifier_RegisterSpendNtfnClient interface {
	Recv() (*SpendEvent, error)
	grpc.ClientStream
}

type chainNotifierRegisterSpendNtfnClient struct {
	grpc.ClientStream
}

func (x *chainNotifierRegisterSpendNtfnClient) Recv() (*SpendEvent, error) {
	m := new(SpendEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chainNotifierClient) RegisterBlockEpochNtfn(ctx context.Context, in *BlockEpoch, opts ...grpc.CallOption) (ChainNotifier_RegisterBlockEpochNtfnClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChainNotifier_ServiceDesc.Streams[2], ChainNotifier_RegisterBlockEpochNtfn_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chainNotifierRegisterBlockEpochNtfnClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChainNotifier_RegisterBlockEpochNtfnClient interface {
	Recv() (*BlockEpoch, error)
	grpc.ClientStream
}

type chainNotifierRegisterBlockEpochNtfnClient struct {
	grpc.ClientStream
}

func (x *chainNotifierRegisterBlockEpochNtfnClient) Recv() (*BlockEpoch, error) {
	m := new(BlockEpoch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChainNotifierServer is the server API for ChainNotifier service.
// All implementations should embed UnimplementedChainNotifierServer
// for forward compatibility
type ChainNotifierServer interface {
	// RegisterConfirmationsNtfn is a synchronous response-streaming RPC that
	// registers an intent for a client to be notified once a confirmation request
	// has reached its required number of confirmations on-chain.
	//
	// A client can specify whether the confirmation request should be for a
	// particular transaction by its hash or for an output script by specifying a
	// zero hash.
	RegisterConfirmationsNtfn(*ConfRequest, ChainNotifier_RegisterConfirmationsNtfnServer) error
	// RegisterSpendNtfn is a synchronous response-streaming RPC that registers an
	// intent for a client to be notification once a spend request has been spent
	// by a transaction that has confirmed on-chain.
	//
	// A client can specify whether the spend request should be for a particular
	// outpoint  or for an output script by specifying a zero outpoint.
	RegisterSpendNtfn(*SpendRequest, ChainNotifier_RegisterSpendNtfnServer) error
	// RegisterBlockEpochNtfn is a synchronous response-streaming RPC that
	// registers an intent for a client to be notified of blocks in the chain. The
	// stream will return a hash and height tuple of a block for each new/stale
	// block in the chain. It is the client's responsibility to determine whether
	// the tuple returned is for a new or stale block in the chain.
	//
	// A client can also request a historical backlog of blocks from a particular
	// point. This allows clients to be idempotent by ensuring that they do not
	// missing processing a single block within the chain.
	RegisterBlockEpochNtfn(*BlockEpoch, ChainNotifier_RegisterBlockEpochNtfnServer) error
}

// UnimplementedChainNotifierServer should be embedded to have forward compatible implementations.
type UnimplementedChainNotifierServer struct {
}

func (UnimplementedChainNotifierServer) RegisterConfirmationsNtfn(*ConfRequest, ChainNotifier_RegisterConfirmationsNtfnServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterConfirmationsNtfn not implemented")
}
func (UnimplementedChainNotifierServer) RegisterSpendNtfn(*SpendRequest, ChainNotifier_RegisterSpendNtfnServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterSpendNtfn not implemented")
}
func (UnimplementedChainNotifierServer) RegisterBlockEpochNtfn(*BlockEpoch, ChainNotifier_RegisterBlockEpochNtfnServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterBlockEpochNtfn not implemented")
}

// UnsafeChainNotifierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChainNotifierServer will
// result in compilation errors.
type UnsafeChainNotifierServer interface {
	mustEmbedUnimplementedChainNotifierServer()
}

func RegisterChainNotifierServer(s grpc.ServiceRegistrar, srv ChainNotifierServer) {
	s.RegisterService(&ChainNotifier_ServiceDesc, srv)
}

func _ChainNotifier_RegisterConfirmationsNtfn_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConfRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainNotifierServer).RegisterConfirmationsNtfn(m, &chainNotifierRegisterConfirmationsNtfnServer{stream})
}

type ChainNotifier_RegisterConfirmationsNtfnServer interface {
	Send(*ConfEvent) error
	grpc.ServerStream
}

type chainNotifierRegisterConfirmationsNtfnServer struct {
	grpc.ServerStream
}

func (x *chainNotifierRegisterConfirmationsNtfnServer) Send(m *ConfEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ChainNotifier_RegisterSpendNtfn_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SpendRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainNotifierServer).RegisterSpendNtfn(m, &chainNotifierRegisterSpendNtfnServer{stream})
}

type ChainNotifier_RegisterSpendNtfnServer interface {
	Send(*SpendEvent) error
	grpc.ServerStream
}

type chainNotifierRegisterSpendNtfnServer struct {
	grpc.ServerStream
}

func (x *chainNotifierRegisterSpendNtfnServer) Send(m *SpendEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ChainNotifier_RegisterBlockEpochNtfn_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockEpoch)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainNotifierServer).RegisterBlockEpochNtfn(m, &chainNotifierRegisterBlockEpochNtfnServer{stream})
}

type ChainNotifier_RegisterBlockEpochNtfnServer interface {
	Send(*BlockEpoch) error
	grpc.ServerStream
}

type chainNotifierRegisterBlockEpochNtfnServer struct {
	grpc.ServerStream
}

func (x *chainNotifierRegisterBlockEpochNtfnServer) Send(m *BlockEpoch) error {
	return x.ServerStream.SendMsg(m)
}

// ChainNotifier_ServiceDesc is the grpc.ServiceDesc for ChainNotifier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChainNotifier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainrpc.ChainNotifier",
	HandlerType: (*ChainNotifierServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterConfirmationsNtfn",
			Handler:       _ChainNotifier_RegisterConfirmationsNtfn_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RegisterSpendNtfn",
			Handler:       _ChainNotifier_RegisterSpendNtfn_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RegisterBlockEpochNtfn",
			Handler:       _ChainNotifier_RegisterBlockEpochNtfn_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chainrpc/chainnotifier.proto",
}
