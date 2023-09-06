// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: routerrpc/router.proto

package routerrpc

import (
	context "context"
	lnrpc "github.com/pkt-cash/pktd/lnd/lnrpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Router_SendPaymentV2_FullMethodName       = "/routerrpc.Router/SendPaymentV2"
	Router_TrackPaymentV2_FullMethodName      = "/routerrpc.Router/TrackPaymentV2"
	Router_EstimateRouteFee_FullMethodName    = "/routerrpc.Router/EstimateRouteFee"
	Router_SendToRoute_FullMethodName         = "/routerrpc.Router/SendToRoute"
	Router_SendToRouteV2_FullMethodName       = "/routerrpc.Router/SendToRouteV2"
	Router_ResetMissionControl_FullMethodName = "/routerrpc.Router/ResetMissionControl"
	Router_QueryMissionControl_FullMethodName = "/routerrpc.Router/QueryMissionControl"
	Router_QueryProbability_FullMethodName    = "/routerrpc.Router/QueryProbability"
	Router_BuildRoute_FullMethodName          = "/routerrpc.Router/BuildRoute"
	Router_SubscribeHtlcEvents_FullMethodName = "/routerrpc.Router/SubscribeHtlcEvents"
	Router_SendPayment_FullMethodName         = "/routerrpc.Router/SendPayment"
	Router_TrackPayment_FullMethodName        = "/routerrpc.Router/TrackPayment"
	Router_HtlcInterceptor_FullMethodName     = "/routerrpc.Router/HtlcInterceptor"
)

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouterClient interface {
	// SendPaymentV2 attempts to route a payment described by the passed
	// PaymentRequest to the final destination. The call returns a stream of
	// payment updates.
	SendPaymentV2(ctx context.Context, in *SendPaymentRequest, opts ...grpc.CallOption) (Router_SendPaymentV2Client, error)
	// TrackPaymentV2 returns an update stream for the payment identified by the
	// payment hash.
	TrackPaymentV2(ctx context.Context, in *TrackPaymentRequest, opts ...grpc.CallOption) (Router_TrackPaymentV2Client, error)
	// EstimateRouteFee allows callers to obtain a lower bound w.r.t how much it
	// may cost to send an HTLC to the target end destination.
	EstimateRouteFee(ctx context.Context, in *RouteFeeRequest, opts ...grpc.CallOption) (*RouteFeeResponse, error)
	// Deprecated: Do not use.
	//
	// Deprecated, use SendToRouteV2. SendToRoute attempts to make a payment via
	// the specified route. This method differs from SendPayment in that it
	// allows users to specify a full route manually. This can be used for
	// things like rebalancing, and atomic swaps. It differs from the newer
	// SendToRouteV2 in that it doesn't return the full HTLC information.
	SendToRoute(ctx context.Context, in *SendToRouteRequest, opts ...grpc.CallOption) (*SendToRouteResponse, error)
	// SendToRouteV2 attempts to make a payment via the specified route. This
	// method differs from SendPayment in that it allows users to specify a full
	// route manually. This can be used for things like rebalancing, and atomic
	// swaps.
	SendToRouteV2(ctx context.Context, in *SendToRouteRequest, opts ...grpc.CallOption) (*lnrpc.HTLCAttempt, error)
	// ResetMissionControl clears all mission control state and starts with a clean
	// slate.
	ResetMissionControl(ctx context.Context, in *ResetMissionControlRequest, opts ...grpc.CallOption) (*ResetMissionControlResponse, error)
	// QueryMissionControl exposes the internal mission control state to callers.
	// It is a development feature.
	QueryMissionControl(ctx context.Context, in *QueryMissionControlRequest, opts ...grpc.CallOption) (*QueryMissionControlResponse, error)
	// QueryProbability returns the current success probability estimate for a
	// given node pair and amount.
	QueryProbability(ctx context.Context, in *QueryProbabilityRequest, opts ...grpc.CallOption) (*QueryProbabilityResponse, error)
	// BuildRoute builds a fully specified route based on a list of hop public
	// keys. It retrieves the relevant channel policies from the graph in order to
	// calculate the correct fees and time locks.
	BuildRoute(ctx context.Context, in *BuildRouteRequest, opts ...grpc.CallOption) (*BuildRouteResponse, error)
	// SubscribeHtlcEvents creates a uni-directional stream from the server to
	// the client which delivers a stream of htlc events.
	SubscribeHtlcEvents(ctx context.Context, in *SubscribeHtlcEventsRequest, opts ...grpc.CallOption) (Router_SubscribeHtlcEventsClient, error)
	// Deprecated: Do not use.
	//
	// Deprecated, use SendPaymentV2. SendPayment attempts to route a payment
	// described by the passed PaymentRequest to the final destination. The call
	// returns a stream of payment status updates.
	SendPayment(ctx context.Context, in *SendPaymentRequest, opts ...grpc.CallOption) (Router_SendPaymentClient, error)
	// Deprecated: Do not use.
	//
	// Deprecated, use TrackPaymentV2. TrackPayment returns an update stream for
	// the payment identified by the payment hash.
	TrackPayment(ctx context.Context, in *TrackPaymentRequest, opts ...grpc.CallOption) (Router_TrackPaymentClient, error)
	// *
	// HtlcInterceptor dispatches a bi-directional streaming RPC in which
	// Forwarded HTLC requests are sent to the client and the client responds with
	// a boolean that tells LND if this htlc should be intercepted.
	// In case of interception, the htlc can be either settled, cancelled or
	// resumed later by using the ResolveHoldForward endpoint.
	HtlcInterceptor(ctx context.Context, opts ...grpc.CallOption) (Router_HtlcInterceptorClient, error)
}

type routerClient struct {
	cc grpc.ClientConnInterface
}

func NewRouterClient(cc grpc.ClientConnInterface) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) SendPaymentV2(ctx context.Context, in *SendPaymentRequest, opts ...grpc.CallOption) (Router_SendPaymentV2Client, error) {
	stream, err := c.cc.NewStream(ctx, &Router_ServiceDesc.Streams[0], Router_SendPaymentV2_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &routerSendPaymentV2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Router_SendPaymentV2Client interface {
	Recv() (*lnrpc.Payment, error)
	grpc.ClientStream
}

type routerSendPaymentV2Client struct {
	grpc.ClientStream
}

func (x *routerSendPaymentV2Client) Recv() (*lnrpc.Payment, error) {
	m := new(lnrpc.Payment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routerClient) TrackPaymentV2(ctx context.Context, in *TrackPaymentRequest, opts ...grpc.CallOption) (Router_TrackPaymentV2Client, error) {
	stream, err := c.cc.NewStream(ctx, &Router_ServiceDesc.Streams[1], Router_TrackPaymentV2_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &routerTrackPaymentV2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Router_TrackPaymentV2Client interface {
	Recv() (*lnrpc.Payment, error)
	grpc.ClientStream
}

type routerTrackPaymentV2Client struct {
	grpc.ClientStream
}

func (x *routerTrackPaymentV2Client) Recv() (*lnrpc.Payment, error) {
	m := new(lnrpc.Payment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routerClient) EstimateRouteFee(ctx context.Context, in *RouteFeeRequest, opts ...grpc.CallOption) (*RouteFeeResponse, error) {
	out := new(RouteFeeResponse)
	err := c.cc.Invoke(ctx, Router_EstimateRouteFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *routerClient) SendToRoute(ctx context.Context, in *SendToRouteRequest, opts ...grpc.CallOption) (*SendToRouteResponse, error) {
	out := new(SendToRouteResponse)
	err := c.cc.Invoke(ctx, Router_SendToRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) SendToRouteV2(ctx context.Context, in *SendToRouteRequest, opts ...grpc.CallOption) (*lnrpc.HTLCAttempt, error) {
	out := new(lnrpc.HTLCAttempt)
	err := c.cc.Invoke(ctx, Router_SendToRouteV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) ResetMissionControl(ctx context.Context, in *ResetMissionControlRequest, opts ...grpc.CallOption) (*ResetMissionControlResponse, error) {
	out := new(ResetMissionControlResponse)
	err := c.cc.Invoke(ctx, Router_ResetMissionControl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) QueryMissionControl(ctx context.Context, in *QueryMissionControlRequest, opts ...grpc.CallOption) (*QueryMissionControlResponse, error) {
	out := new(QueryMissionControlResponse)
	err := c.cc.Invoke(ctx, Router_QueryMissionControl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) QueryProbability(ctx context.Context, in *QueryProbabilityRequest, opts ...grpc.CallOption) (*QueryProbabilityResponse, error) {
	out := new(QueryProbabilityResponse)
	err := c.cc.Invoke(ctx, Router_QueryProbability_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) BuildRoute(ctx context.Context, in *BuildRouteRequest, opts ...grpc.CallOption) (*BuildRouteResponse, error) {
	out := new(BuildRouteResponse)
	err := c.cc.Invoke(ctx, Router_BuildRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) SubscribeHtlcEvents(ctx context.Context, in *SubscribeHtlcEventsRequest, opts ...grpc.CallOption) (Router_SubscribeHtlcEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Router_ServiceDesc.Streams[2], Router_SubscribeHtlcEvents_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &routerSubscribeHtlcEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Router_SubscribeHtlcEventsClient interface {
	Recv() (*HtlcEvent, error)
	grpc.ClientStream
}

type routerSubscribeHtlcEventsClient struct {
	grpc.ClientStream
}

func (x *routerSubscribeHtlcEventsClient) Recv() (*HtlcEvent, error) {
	m := new(HtlcEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Deprecated: Do not use.
func (c *routerClient) SendPayment(ctx context.Context, in *SendPaymentRequest, opts ...grpc.CallOption) (Router_SendPaymentClient, error) {
	stream, err := c.cc.NewStream(ctx, &Router_ServiceDesc.Streams[3], Router_SendPayment_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &routerSendPaymentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Router_SendPaymentClient interface {
	Recv() (*PaymentStatus, error)
	grpc.ClientStream
}

type routerSendPaymentClient struct {
	grpc.ClientStream
}

func (x *routerSendPaymentClient) Recv() (*PaymentStatus, error) {
	m := new(PaymentStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Deprecated: Do not use.
func (c *routerClient) TrackPayment(ctx context.Context, in *TrackPaymentRequest, opts ...grpc.CallOption) (Router_TrackPaymentClient, error) {
	stream, err := c.cc.NewStream(ctx, &Router_ServiceDesc.Streams[4], Router_TrackPayment_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &routerTrackPaymentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Router_TrackPaymentClient interface {
	Recv() (*PaymentStatus, error)
	grpc.ClientStream
}

type routerTrackPaymentClient struct {
	grpc.ClientStream
}

func (x *routerTrackPaymentClient) Recv() (*PaymentStatus, error) {
	m := new(PaymentStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routerClient) HtlcInterceptor(ctx context.Context, opts ...grpc.CallOption) (Router_HtlcInterceptorClient, error) {
	stream, err := c.cc.NewStream(ctx, &Router_ServiceDesc.Streams[5], Router_HtlcInterceptor_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &routerHtlcInterceptorClient{stream}
	return x, nil
}

type Router_HtlcInterceptorClient interface {
	Send(*ForwardHtlcInterceptResponse) error
	Recv() (*ForwardHtlcInterceptRequest, error)
	grpc.ClientStream
}

type routerHtlcInterceptorClient struct {
	grpc.ClientStream
}

func (x *routerHtlcInterceptorClient) Send(m *ForwardHtlcInterceptResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routerHtlcInterceptorClient) Recv() (*ForwardHtlcInterceptRequest, error) {
	m := new(ForwardHtlcInterceptRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouterServer is the server API for Router service.
// All implementations should embed UnimplementedRouterServer
// for forward compatibility
type RouterServer interface {
	// SendPaymentV2 attempts to route a payment described by the passed
	// PaymentRequest to the final destination. The call returns a stream of
	// payment updates.
	SendPaymentV2(*SendPaymentRequest, Router_SendPaymentV2Server) error
	// TrackPaymentV2 returns an update stream for the payment identified by the
	// payment hash.
	TrackPaymentV2(*TrackPaymentRequest, Router_TrackPaymentV2Server) error
	// EstimateRouteFee allows callers to obtain a lower bound w.r.t how much it
	// may cost to send an HTLC to the target end destination.
	EstimateRouteFee(context.Context, *RouteFeeRequest) (*RouteFeeResponse, error)
	// Deprecated: Do not use.
	//
	// Deprecated, use SendToRouteV2. SendToRoute attempts to make a payment via
	// the specified route. This method differs from SendPayment in that it
	// allows users to specify a full route manually. This can be used for
	// things like rebalancing, and atomic swaps. It differs from the newer
	// SendToRouteV2 in that it doesn't return the full HTLC information.
	SendToRoute(context.Context, *SendToRouteRequest) (*SendToRouteResponse, error)
	// SendToRouteV2 attempts to make a payment via the specified route. This
	// method differs from SendPayment in that it allows users to specify a full
	// route manually. This can be used for things like rebalancing, and atomic
	// swaps.
	SendToRouteV2(context.Context, *SendToRouteRequest) (*lnrpc.HTLCAttempt, error)
	// ResetMissionControl clears all mission control state and starts with a clean
	// slate.
	ResetMissionControl(context.Context, *ResetMissionControlRequest) (*ResetMissionControlResponse, error)
	// QueryMissionControl exposes the internal mission control state to callers.
	// It is a development feature.
	QueryMissionControl(context.Context, *QueryMissionControlRequest) (*QueryMissionControlResponse, error)
	// QueryProbability returns the current success probability estimate for a
	// given node pair and amount.
	QueryProbability(context.Context, *QueryProbabilityRequest) (*QueryProbabilityResponse, error)
	// BuildRoute builds a fully specified route based on a list of hop public
	// keys. It retrieves the relevant channel policies from the graph in order to
	// calculate the correct fees and time locks.
	BuildRoute(context.Context, *BuildRouteRequest) (*BuildRouteResponse, error)
	// SubscribeHtlcEvents creates a uni-directional stream from the server to
	// the client which delivers a stream of htlc events.
	SubscribeHtlcEvents(*SubscribeHtlcEventsRequest, Router_SubscribeHtlcEventsServer) error
	// Deprecated: Do not use.
	//
	// Deprecated, use SendPaymentV2. SendPayment attempts to route a payment
	// described by the passed PaymentRequest to the final destination. The call
	// returns a stream of payment status updates.
	SendPayment(*SendPaymentRequest, Router_SendPaymentServer) error
	// Deprecated: Do not use.
	//
	// Deprecated, use TrackPaymentV2. TrackPayment returns an update stream for
	// the payment identified by the payment hash.
	TrackPayment(*TrackPaymentRequest, Router_TrackPaymentServer) error
	// *
	// HtlcInterceptor dispatches a bi-directional streaming RPC in which
	// Forwarded HTLC requests are sent to the client and the client responds with
	// a boolean that tells LND if this htlc should be intercepted.
	// In case of interception, the htlc can be either settled, cancelled or
	// resumed later by using the ResolveHoldForward endpoint.
	HtlcInterceptor(Router_HtlcInterceptorServer) error
}

// UnimplementedRouterServer should be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (UnimplementedRouterServer) SendPaymentV2(*SendPaymentRequest, Router_SendPaymentV2Server) error {
	return status.Errorf(codes.Unimplemented, "method SendPaymentV2 not implemented")
}
func (UnimplementedRouterServer) TrackPaymentV2(*TrackPaymentRequest, Router_TrackPaymentV2Server) error {
	return status.Errorf(codes.Unimplemented, "method TrackPaymentV2 not implemented")
}
func (UnimplementedRouterServer) EstimateRouteFee(context.Context, *RouteFeeRequest) (*RouteFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimateRouteFee not implemented")
}
func (UnimplementedRouterServer) SendToRoute(context.Context, *SendToRouteRequest) (*SendToRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToRoute not implemented")
}
func (UnimplementedRouterServer) SendToRouteV2(context.Context, *SendToRouteRequest) (*lnrpc.HTLCAttempt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToRouteV2 not implemented")
}
func (UnimplementedRouterServer) ResetMissionControl(context.Context, *ResetMissionControlRequest) (*ResetMissionControlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetMissionControl not implemented")
}
func (UnimplementedRouterServer) QueryMissionControl(context.Context, *QueryMissionControlRequest) (*QueryMissionControlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMissionControl not implemented")
}
func (UnimplementedRouterServer) QueryProbability(context.Context, *QueryProbabilityRequest) (*QueryProbabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProbability not implemented")
}
func (UnimplementedRouterServer) BuildRoute(context.Context, *BuildRouteRequest) (*BuildRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildRoute not implemented")
}
func (UnimplementedRouterServer) SubscribeHtlcEvents(*SubscribeHtlcEventsRequest, Router_SubscribeHtlcEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeHtlcEvents not implemented")
}
func (UnimplementedRouterServer) SendPayment(*SendPaymentRequest, Router_SendPaymentServer) error {
	return status.Errorf(codes.Unimplemented, "method SendPayment not implemented")
}
func (UnimplementedRouterServer) TrackPayment(*TrackPaymentRequest, Router_TrackPaymentServer) error {
	return status.Errorf(codes.Unimplemented, "method TrackPayment not implemented")
}
func (UnimplementedRouterServer) HtlcInterceptor(Router_HtlcInterceptorServer) error {
	return status.Errorf(codes.Unimplemented, "method HtlcInterceptor not implemented")
}

// UnsafeRouterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouterServer will
// result in compilation errors.
type UnsafeRouterServer interface {
	mustEmbedUnimplementedRouterServer()
}

func RegisterRouterServer(s grpc.ServiceRegistrar, srv RouterServer) {
	s.RegisterService(&Router_ServiceDesc, srv)
}

func _Router_SendPaymentV2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SendPaymentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouterServer).SendPaymentV2(m, &routerSendPaymentV2Server{stream})
}

type Router_SendPaymentV2Server interface {
	Send(*lnrpc.Payment) error
	grpc.ServerStream
}

type routerSendPaymentV2Server struct {
	grpc.ServerStream
}

func (x *routerSendPaymentV2Server) Send(m *lnrpc.Payment) error {
	return x.ServerStream.SendMsg(m)
}

func _Router_TrackPaymentV2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TrackPaymentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouterServer).TrackPaymentV2(m, &routerTrackPaymentV2Server{stream})
}

type Router_TrackPaymentV2Server interface {
	Send(*lnrpc.Payment) error
	grpc.ServerStream
}

type routerTrackPaymentV2Server struct {
	grpc.ServerStream
}

func (x *routerTrackPaymentV2Server) Send(m *lnrpc.Payment) error {
	return x.ServerStream.SendMsg(m)
}

func _Router_EstimateRouteFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).EstimateRouteFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_EstimateRouteFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).EstimateRouteFee(ctx, req.(*RouteFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_SendToRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendToRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).SendToRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_SendToRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).SendToRoute(ctx, req.(*SendToRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_SendToRouteV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendToRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).SendToRouteV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_SendToRouteV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).SendToRouteV2(ctx, req.(*SendToRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_ResetMissionControl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetMissionControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).ResetMissionControl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_ResetMissionControl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).ResetMissionControl(ctx, req.(*ResetMissionControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_QueryMissionControl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMissionControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).QueryMissionControl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_QueryMissionControl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).QueryMissionControl(ctx, req.(*QueryMissionControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_QueryProbability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProbabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).QueryProbability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_QueryProbability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).QueryProbability(ctx, req.(*QueryProbabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_BuildRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).BuildRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_BuildRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).BuildRoute(ctx, req.(*BuildRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_SubscribeHtlcEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeHtlcEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouterServer).SubscribeHtlcEvents(m, &routerSubscribeHtlcEventsServer{stream})
}

type Router_SubscribeHtlcEventsServer interface {
	Send(*HtlcEvent) error
	grpc.ServerStream
}

type routerSubscribeHtlcEventsServer struct {
	grpc.ServerStream
}

func (x *routerSubscribeHtlcEventsServer) Send(m *HtlcEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _Router_SendPayment_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SendPaymentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouterServer).SendPayment(m, &routerSendPaymentServer{stream})
}

type Router_SendPaymentServer interface {
	Send(*PaymentStatus) error
	grpc.ServerStream
}

type routerSendPaymentServer struct {
	grpc.ServerStream
}

func (x *routerSendPaymentServer) Send(m *PaymentStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _Router_TrackPayment_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TrackPaymentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouterServer).TrackPayment(m, &routerTrackPaymentServer{stream})
}

type Router_TrackPaymentServer interface {
	Send(*PaymentStatus) error
	grpc.ServerStream
}

type routerTrackPaymentServer struct {
	grpc.ServerStream
}

func (x *routerTrackPaymentServer) Send(m *PaymentStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _Router_HtlcInterceptor_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouterServer).HtlcInterceptor(&routerHtlcInterceptorServer{stream})
}

type Router_HtlcInterceptorServer interface {
	Send(*ForwardHtlcInterceptRequest) error
	Recv() (*ForwardHtlcInterceptResponse, error)
	grpc.ServerStream
}

type routerHtlcInterceptorServer struct {
	grpc.ServerStream
}

func (x *routerHtlcInterceptorServer) Send(m *ForwardHtlcInterceptRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routerHtlcInterceptorServer) Recv() (*ForwardHtlcInterceptResponse, error) {
	m := new(ForwardHtlcInterceptResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Router_ServiceDesc is the grpc.ServiceDesc for Router service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Router_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routerrpc.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EstimateRouteFee",
			Handler:    _Router_EstimateRouteFee_Handler,
		},
		{
			MethodName: "SendToRoute",
			Handler:    _Router_SendToRoute_Handler,
		},
		{
			MethodName: "SendToRouteV2",
			Handler:    _Router_SendToRouteV2_Handler,
		},
		{
			MethodName: "ResetMissionControl",
			Handler:    _Router_ResetMissionControl_Handler,
		},
		{
			MethodName: "QueryMissionControl",
			Handler:    _Router_QueryMissionControl_Handler,
		},
		{
			MethodName: "QueryProbability",
			Handler:    _Router_QueryProbability_Handler,
		},
		{
			MethodName: "BuildRoute",
			Handler:    _Router_BuildRoute_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendPaymentV2",
			Handler:       _Router_SendPaymentV2_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TrackPaymentV2",
			Handler:       _Router_TrackPaymentV2_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeHtlcEvents",
			Handler:       _Router_SubscribeHtlcEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendPayment",
			Handler:       _Router_SendPayment_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TrackPayment",
			Handler:       _Router_TrackPayment_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HtlcInterceptor",
			Handler:       _Router_HtlcInterceptor_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "routerrpc/router.proto",
}
