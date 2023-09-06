// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: signrpc/signer.proto

package signrpc

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
	Signer_SignOutputRaw_FullMethodName      = "/signrpc.Signer/SignOutputRaw"
	Signer_ComputeInputScript_FullMethodName = "/signrpc.Signer/ComputeInputScript"
	Signer_SignMessage_FullMethodName        = "/signrpc.Signer/SignMessage"
	Signer_VerifyMessage_FullMethodName      = "/signrpc.Signer/VerifyMessage"
	Signer_DeriveSharedKey_FullMethodName    = "/signrpc.Signer/DeriveSharedKey"
)

// SignerClient is the client API for Signer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignerClient interface {
	// SignOutputRaw is a method that can be used to generated a signature for a
	// set of inputs/outputs to a transaction. Each request specifies details
	// concerning how the outputs should be signed, which keys they should be
	// signed with, and also any optional tweaks. The return value is a fixed
	// 64-byte signature (the same format as we use on the wire in Lightning).
	//
	// If we are  unable to sign using the specified keys, then an error will be
	// returned.
	SignOutputRaw(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*SignResp, error)
	// ComputeInputScript generates a complete InputIndex for the passed
	// transaction with the signature as defined within the passed SignDescriptor.
	// This method should be capable of generating the proper input script for
	// both regular p2wkh output and p2wkh outputs nested within a regular p2sh
	// output.
	//
	// Note that when using this method to sign inputs belonging to the wallet,
	// the only items of the SignDescriptor that need to be populated are pkScript
	// in the TxOut field, the value in that same field, and finally the input
	// index.
	ComputeInputScript(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*InputScriptResp, error)
	// SignMessage signs a message with the key specified in the key locator. The
	// returned signature is fixed-size LN wire format encoded.
	//
	// The main difference to SignMessage in the main RPC is that a specific key is
	// used to sign the message instead of the node identity private key.
	SignMessage(ctx context.Context, in *SignMessageReq, opts ...grpc.CallOption) (*SignMessageResp, error)
	// VerifyMessage verifies a signature over a message using the public key
	// provided. The signature must be fixed-size LN wire format encoded.
	//
	// The main difference to VerifyMessage in the main RPC is that the public key
	// used to sign the message does not have to be a node known to the network.
	VerifyMessage(ctx context.Context, in *VerifyMessageReq, opts ...grpc.CallOption) (*VerifyMessageResp, error)
	// DeriveSharedKey returns a shared secret key by performing Diffie-Hellman key
	// derivation between the ephemeral public key in the request and the node's
	// key specified in the key_desc parameter. Either a key locator or a raw
	// public key is expected in the key_desc, if neither is supplied, defaults to
	// the node's identity private key:
	// P_shared = privKeyNode * ephemeralPubkey
	// The resulting shared public key is serialized in the compressed format and
	// hashed with sha256, resulting in the final key length of 256bit.
	DeriveSharedKey(ctx context.Context, in *SharedKeyRequest, opts ...grpc.CallOption) (*SharedKeyResponse, error)
}

type signerClient struct {
	cc grpc.ClientConnInterface
}

func NewSignerClient(cc grpc.ClientConnInterface) SignerClient {
	return &signerClient{cc}
}

func (c *signerClient) SignOutputRaw(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*SignResp, error) {
	out := new(SignResp)
	err := c.cc.Invoke(ctx, Signer_SignOutputRaw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signerClient) ComputeInputScript(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*InputScriptResp, error) {
	out := new(InputScriptResp)
	err := c.cc.Invoke(ctx, Signer_ComputeInputScript_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signerClient) SignMessage(ctx context.Context, in *SignMessageReq, opts ...grpc.CallOption) (*SignMessageResp, error) {
	out := new(SignMessageResp)
	err := c.cc.Invoke(ctx, Signer_SignMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signerClient) VerifyMessage(ctx context.Context, in *VerifyMessageReq, opts ...grpc.CallOption) (*VerifyMessageResp, error) {
	out := new(VerifyMessageResp)
	err := c.cc.Invoke(ctx, Signer_VerifyMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signerClient) DeriveSharedKey(ctx context.Context, in *SharedKeyRequest, opts ...grpc.CallOption) (*SharedKeyResponse, error) {
	out := new(SharedKeyResponse)
	err := c.cc.Invoke(ctx, Signer_DeriveSharedKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignerServer is the server API for Signer service.
// All implementations should embed UnimplementedSignerServer
// for forward compatibility
type SignerServer interface {
	// SignOutputRaw is a method that can be used to generated a signature for a
	// set of inputs/outputs to a transaction. Each request specifies details
	// concerning how the outputs should be signed, which keys they should be
	// signed with, and also any optional tweaks. The return value is a fixed
	// 64-byte signature (the same format as we use on the wire in Lightning).
	//
	// If we are  unable to sign using the specified keys, then an error will be
	// returned.
	SignOutputRaw(context.Context, *SignReq) (*SignResp, error)
	// ComputeInputScript generates a complete InputIndex for the passed
	// transaction with the signature as defined within the passed SignDescriptor.
	// This method should be capable of generating the proper input script for
	// both regular p2wkh output and p2wkh outputs nested within a regular p2sh
	// output.
	//
	// Note that when using this method to sign inputs belonging to the wallet,
	// the only items of the SignDescriptor that need to be populated are pkScript
	// in the TxOut field, the value in that same field, and finally the input
	// index.
	ComputeInputScript(context.Context, *SignReq) (*InputScriptResp, error)
	// SignMessage signs a message with the key specified in the key locator. The
	// returned signature is fixed-size LN wire format encoded.
	//
	// The main difference to SignMessage in the main RPC is that a specific key is
	// used to sign the message instead of the node identity private key.
	SignMessage(context.Context, *SignMessageReq) (*SignMessageResp, error)
	// VerifyMessage verifies a signature over a message using the public key
	// provided. The signature must be fixed-size LN wire format encoded.
	//
	// The main difference to VerifyMessage in the main RPC is that the public key
	// used to sign the message does not have to be a node known to the network.
	VerifyMessage(context.Context, *VerifyMessageReq) (*VerifyMessageResp, error)
	// DeriveSharedKey returns a shared secret key by performing Diffie-Hellman key
	// derivation between the ephemeral public key in the request and the node's
	// key specified in the key_desc parameter. Either a key locator or a raw
	// public key is expected in the key_desc, if neither is supplied, defaults to
	// the node's identity private key:
	// P_shared = privKeyNode * ephemeralPubkey
	// The resulting shared public key is serialized in the compressed format and
	// hashed with sha256, resulting in the final key length of 256bit.
	DeriveSharedKey(context.Context, *SharedKeyRequest) (*SharedKeyResponse, error)
}

// UnimplementedSignerServer should be embedded to have forward compatible implementations.
type UnimplementedSignerServer struct {
}

func (UnimplementedSignerServer) SignOutputRaw(context.Context, *SignReq) (*SignResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOutputRaw not implemented")
}
func (UnimplementedSignerServer) ComputeInputScript(context.Context, *SignReq) (*InputScriptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeInputScript not implemented")
}
func (UnimplementedSignerServer) SignMessage(context.Context, *SignMessageReq) (*SignMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignMessage not implemented")
}
func (UnimplementedSignerServer) VerifyMessage(context.Context, *VerifyMessageReq) (*VerifyMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyMessage not implemented")
}
func (UnimplementedSignerServer) DeriveSharedKey(context.Context, *SharedKeyRequest) (*SharedKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeriveSharedKey not implemented")
}

// UnsafeSignerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignerServer will
// result in compilation errors.
type UnsafeSignerServer interface {
	mustEmbedUnimplementedSignerServer()
}

func RegisterSignerServer(s grpc.ServiceRegistrar, srv SignerServer) {
	s.RegisterService(&Signer_ServiceDesc, srv)
}

func _Signer_SignOutputRaw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).SignOutputRaw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Signer_SignOutputRaw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).SignOutputRaw(ctx, req.(*SignReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signer_ComputeInputScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).ComputeInputScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Signer_ComputeInputScript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).ComputeInputScript(ctx, req.(*SignReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signer_SignMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).SignMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Signer_SignMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).SignMessage(ctx, req.(*SignMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signer_VerifyMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).VerifyMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Signer_VerifyMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).VerifyMessage(ctx, req.(*VerifyMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signer_DeriveSharedKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SharedKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).DeriveSharedKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Signer_DeriveSharedKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).DeriveSharedKey(ctx, req.(*SharedKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Signer_ServiceDesc is the grpc.ServiceDesc for Signer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Signer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "signrpc.Signer",
	HandlerType: (*SignerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignOutputRaw",
			Handler:    _Signer_SignOutputRaw_Handler,
		},
		{
			MethodName: "ComputeInputScript",
			Handler:    _Signer_ComputeInputScript_Handler,
		},
		{
			MethodName: "SignMessage",
			Handler:    _Signer_SignMessage_Handler,
		},
		{
			MethodName: "VerifyMessage",
			Handler:    _Signer_VerifyMessage_Handler,
		},
		{
			MethodName: "DeriveSharedKey",
			Handler:    _Signer_DeriveSharedKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signrpc/signer.proto",
}
