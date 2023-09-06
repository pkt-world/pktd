// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: watchtowerrpc/watchtower.proto

package watchtowerrpc

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
	Watchtower_GetInfo_FullMethodName = "/watchtowerrpc.Watchtower/GetInfo"
)

// WatchtowerClient is the client API for Watchtower service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatchtowerClient interface {
	// lncli: tower info
	// GetInfo returns general information concerning the companion watchtower
	// including its public key and URIs where the server is currently
	// listening for clients.
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
}

type watchtowerClient struct {
	cc grpc.ClientConnInterface
}

func NewWatchtowerClient(cc grpc.ClientConnInterface) WatchtowerClient {
	return &watchtowerClient{cc}
}

func (c *watchtowerClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, Watchtower_GetInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WatchtowerServer is the server API for Watchtower service.
// All implementations should embed UnimplementedWatchtowerServer
// for forward compatibility
type WatchtowerServer interface {
	// lncli: tower info
	// GetInfo returns general information concerning the companion watchtower
	// including its public key and URIs where the server is currently
	// listening for clients.
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
}

// UnimplementedWatchtowerServer should be embedded to have forward compatible implementations.
type UnimplementedWatchtowerServer struct {
}

func (UnimplementedWatchtowerServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}

// UnsafeWatchtowerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatchtowerServer will
// result in compilation errors.
type UnsafeWatchtowerServer interface {
	mustEmbedUnimplementedWatchtowerServer()
}

func RegisterWatchtowerServer(s grpc.ServiceRegistrar, srv WatchtowerServer) {
	s.RegisterService(&Watchtower_ServiceDesc, srv)
}

func _Watchtower_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchtowerServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Watchtower_GetInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchtowerServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Watchtower_ServiceDesc is the grpc.ServiceDesc for Watchtower service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Watchtower_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "watchtowerrpc.Watchtower",
	HandlerType: (*WatchtowerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Watchtower_GetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "watchtowerrpc/watchtower.proto",
}
