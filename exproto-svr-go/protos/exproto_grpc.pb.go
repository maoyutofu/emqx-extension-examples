// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: protos/exproto.proto

package __

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

// ConnectionAdapterClient is the client API for ConnectionAdapter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionAdapterClient interface {
	Send(ctx context.Context, in *SendBytesRequest, opts ...grpc.CallOption) (*CodeResponse, error)
	Close(ctx context.Context, in *CloseSocketRequest, opts ...grpc.CallOption) (*CodeResponse, error)
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*CodeResponse, error)
	StartTimer(ctx context.Context, in *TimerRequest, opts ...grpc.CallOption) (*CodeResponse, error)
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*CodeResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*CodeResponse, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*CodeResponse, error)
}

type connectionAdapterClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionAdapterClient(cc grpc.ClientConnInterface) ConnectionAdapterClient {
	return &connectionAdapterClient{cc}
}

func (c *connectionAdapterClient) Send(ctx context.Context, in *SendBytesRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
	out := new(CodeResponse)
	err := c.cc.Invoke(ctx, "/emqx.exproto.v1.ConnectionAdapter/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionAdapterClient) Close(ctx context.Context, in *CloseSocketRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
	out := new(CodeResponse)
	err := c.cc.Invoke(ctx, "/emqx.exproto.v1.ConnectionAdapter/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionAdapterClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
	out := new(CodeResponse)
	err := c.cc.Invoke(ctx, "/emqx.exproto.v1.ConnectionAdapter/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionAdapterClient) StartTimer(ctx context.Context, in *TimerRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
	out := new(CodeResponse)
	err := c.cc.Invoke(ctx, "/emqx.exproto.v1.ConnectionAdapter/StartTimer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionAdapterClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
	out := new(CodeResponse)
	err := c.cc.Invoke(ctx, "/emqx.exproto.v1.ConnectionAdapter/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionAdapterClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
	out := new(CodeResponse)
	err := c.cc.Invoke(ctx, "/emqx.exproto.v1.ConnectionAdapter/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionAdapterClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
	out := new(CodeResponse)
	err := c.cc.Invoke(ctx, "/emqx.exproto.v1.ConnectionAdapter/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectionAdapterServer is the server API for ConnectionAdapter service.
// All implementations must embed UnimplementedConnectionAdapterServer
// for forward compatibility
type ConnectionAdapterServer interface {
	Send(context.Context, *SendBytesRequest) (*CodeResponse, error)
	Close(context.Context, *CloseSocketRequest) (*CodeResponse, error)
	Authenticate(context.Context, *AuthenticateRequest) (*CodeResponse, error)
	StartTimer(context.Context, *TimerRequest) (*CodeResponse, error)
	Publish(context.Context, *PublishRequest) (*CodeResponse, error)
	Subscribe(context.Context, *SubscribeRequest) (*CodeResponse, error)
	Unsubscribe(context.Context, *UnsubscribeRequest) (*CodeResponse, error)
	mustEmbedUnimplementedConnectionAdapterServer()
}

// UnimplementedConnectionAdapterServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionAdapterServer struct {
}

func (UnimplementedConnectionAdapterServer) Send(context.Context, *SendBytesRequest) (*CodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedConnectionAdapterServer) Close(context.Context, *CloseSocketRequest) (*CodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (UnimplementedConnectionAdapterServer) Authenticate(context.Context, *AuthenticateRequest) (*CodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedConnectionAdapterServer) StartTimer(context.Context, *TimerRequest) (*CodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartTimer not implemented")
}
func (UnimplementedConnectionAdapterServer) Publish(context.Context, *PublishRequest) (*CodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedConnectionAdapterServer) Subscribe(context.Context, *SubscribeRequest) (*CodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedConnectionAdapterServer) Unsubscribe(context.Context, *UnsubscribeRequest) (*CodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedConnectionAdapterServer) mustEmbedUnimplementedConnectionAdapterServer() {}

// UnsafeConnectionAdapterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionAdapterServer will
// result in compilation errors.
type UnsafeConnectionAdapterServer interface {
	mustEmbedUnimplementedConnectionAdapterServer()
}

func RegisterConnectionAdapterServer(s grpc.ServiceRegistrar, srv ConnectionAdapterServer) {
	s.RegisterService(&ConnectionAdapter_ServiceDesc, srv)
}

func _ConnectionAdapter_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendBytesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionAdapterServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emqx.exproto.v1.ConnectionAdapter/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionAdapterServer).Send(ctx, req.(*SendBytesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionAdapter_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSocketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionAdapterServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emqx.exproto.v1.ConnectionAdapter/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionAdapterServer).Close(ctx, req.(*CloseSocketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionAdapter_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionAdapterServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emqx.exproto.v1.ConnectionAdapter/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionAdapterServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionAdapter_StartTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionAdapterServer).StartTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emqx.exproto.v1.ConnectionAdapter/StartTimer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionAdapterServer).StartTimer(ctx, req.(*TimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionAdapter_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionAdapterServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emqx.exproto.v1.ConnectionAdapter/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionAdapterServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionAdapter_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionAdapterServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emqx.exproto.v1.ConnectionAdapter/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionAdapterServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionAdapter_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionAdapterServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emqx.exproto.v1.ConnectionAdapter/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionAdapterServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectionAdapter_ServiceDesc is the grpc.ServiceDesc for ConnectionAdapter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectionAdapter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emqx.exproto.v1.ConnectionAdapter",
	HandlerType: (*ConnectionAdapterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _ConnectionAdapter_Send_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _ConnectionAdapter_Close_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _ConnectionAdapter_Authenticate_Handler,
		},
		{
			MethodName: "StartTimer",
			Handler:    _ConnectionAdapter_StartTimer_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _ConnectionAdapter_Publish_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _ConnectionAdapter_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _ConnectionAdapter_Unsubscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/exproto.proto",
}

// ConnectionHandlerClient is the client API for ConnectionHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionHandlerClient interface {
	OnSocketCreated(ctx context.Context, opts ...grpc.CallOption) (ConnectionHandler_OnSocketCreatedClient, error)
	OnSocketClosed(ctx context.Context, opts ...grpc.CallOption) (ConnectionHandler_OnSocketClosedClient, error)
	OnReceivedBytes(ctx context.Context, opts ...grpc.CallOption) (ConnectionHandler_OnReceivedBytesClient, error)
	OnTimerTimeout(ctx context.Context, opts ...grpc.CallOption) (ConnectionHandler_OnTimerTimeoutClient, error)
	OnReceivedMessages(ctx context.Context, opts ...grpc.CallOption) (ConnectionHandler_OnReceivedMessagesClient, error)
}

type connectionHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionHandlerClient(cc grpc.ClientConnInterface) ConnectionHandlerClient {
	return &connectionHandlerClient{cc}
}

func (c *connectionHandlerClient) OnSocketCreated(ctx context.Context, opts ...grpc.CallOption) (ConnectionHandler_OnSocketCreatedClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConnectionHandler_ServiceDesc.Streams[0], "/emqx.exproto.v1.ConnectionHandler/OnSocketCreated", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectionHandlerOnSocketCreatedClient{stream}
	return x, nil
}

type ConnectionHandler_OnSocketCreatedClient interface {
	Send(*SocketCreatedRequest) error
	CloseAndRecv() (*EmptySuccess, error)
	grpc.ClientStream
}

type connectionHandlerOnSocketCreatedClient struct {
	grpc.ClientStream
}

func (x *connectionHandlerOnSocketCreatedClient) Send(m *SocketCreatedRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectionHandlerOnSocketCreatedClient) CloseAndRecv() (*EmptySuccess, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptySuccess)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *connectionHandlerClient) OnSocketClosed(ctx context.Context, opts ...grpc.CallOption) (ConnectionHandler_OnSocketClosedClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConnectionHandler_ServiceDesc.Streams[1], "/emqx.exproto.v1.ConnectionHandler/OnSocketClosed", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectionHandlerOnSocketClosedClient{stream}
	return x, nil
}

type ConnectionHandler_OnSocketClosedClient interface {
	Send(*SocketClosedRequest) error
	CloseAndRecv() (*EmptySuccess, error)
	grpc.ClientStream
}

type connectionHandlerOnSocketClosedClient struct {
	grpc.ClientStream
}

func (x *connectionHandlerOnSocketClosedClient) Send(m *SocketClosedRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectionHandlerOnSocketClosedClient) CloseAndRecv() (*EmptySuccess, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptySuccess)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *connectionHandlerClient) OnReceivedBytes(ctx context.Context, opts ...grpc.CallOption) (ConnectionHandler_OnReceivedBytesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConnectionHandler_ServiceDesc.Streams[2], "/emqx.exproto.v1.ConnectionHandler/OnReceivedBytes", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectionHandlerOnReceivedBytesClient{stream}
	return x, nil
}

type ConnectionHandler_OnReceivedBytesClient interface {
	Send(*ReceivedBytesRequest) error
	CloseAndRecv() (*EmptySuccess, error)
	grpc.ClientStream
}

type connectionHandlerOnReceivedBytesClient struct {
	grpc.ClientStream
}

func (x *connectionHandlerOnReceivedBytesClient) Send(m *ReceivedBytesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectionHandlerOnReceivedBytesClient) CloseAndRecv() (*EmptySuccess, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptySuccess)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *connectionHandlerClient) OnTimerTimeout(ctx context.Context, opts ...grpc.CallOption) (ConnectionHandler_OnTimerTimeoutClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConnectionHandler_ServiceDesc.Streams[3], "/emqx.exproto.v1.ConnectionHandler/OnTimerTimeout", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectionHandlerOnTimerTimeoutClient{stream}
	return x, nil
}

type ConnectionHandler_OnTimerTimeoutClient interface {
	Send(*TimerTimeoutRequest) error
	CloseAndRecv() (*EmptySuccess, error)
	grpc.ClientStream
}

type connectionHandlerOnTimerTimeoutClient struct {
	grpc.ClientStream
}

func (x *connectionHandlerOnTimerTimeoutClient) Send(m *TimerTimeoutRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectionHandlerOnTimerTimeoutClient) CloseAndRecv() (*EmptySuccess, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptySuccess)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *connectionHandlerClient) OnReceivedMessages(ctx context.Context, opts ...grpc.CallOption) (ConnectionHandler_OnReceivedMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConnectionHandler_ServiceDesc.Streams[4], "/emqx.exproto.v1.ConnectionHandler/OnReceivedMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectionHandlerOnReceivedMessagesClient{stream}
	return x, nil
}

type ConnectionHandler_OnReceivedMessagesClient interface {
	Send(*ReceivedMessagesRequest) error
	CloseAndRecv() (*EmptySuccess, error)
	grpc.ClientStream
}

type connectionHandlerOnReceivedMessagesClient struct {
	grpc.ClientStream
}

func (x *connectionHandlerOnReceivedMessagesClient) Send(m *ReceivedMessagesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectionHandlerOnReceivedMessagesClient) CloseAndRecv() (*EmptySuccess, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptySuccess)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConnectionHandlerServer is the server API for ConnectionHandler service.
// All implementations must embed UnimplementedConnectionHandlerServer
// for forward compatibility
type ConnectionHandlerServer interface {
	OnSocketCreated(ConnectionHandler_OnSocketCreatedServer) error
	OnSocketClosed(ConnectionHandler_OnSocketClosedServer) error
	OnReceivedBytes(ConnectionHandler_OnReceivedBytesServer) error
	OnTimerTimeout(ConnectionHandler_OnTimerTimeoutServer) error
	OnReceivedMessages(ConnectionHandler_OnReceivedMessagesServer) error
	mustEmbedUnimplementedConnectionHandlerServer()
}

// UnimplementedConnectionHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionHandlerServer struct {
}

func (UnimplementedConnectionHandlerServer) OnSocketCreated(ConnectionHandler_OnSocketCreatedServer) error {
	return status.Errorf(codes.Unimplemented, "method OnSocketCreated not implemented")
}
func (UnimplementedConnectionHandlerServer) OnSocketClosed(ConnectionHandler_OnSocketClosedServer) error {
	return status.Errorf(codes.Unimplemented, "method OnSocketClosed not implemented")
}
func (UnimplementedConnectionHandlerServer) OnReceivedBytes(ConnectionHandler_OnReceivedBytesServer) error {
	return status.Errorf(codes.Unimplemented, "method OnReceivedBytes not implemented")
}
func (UnimplementedConnectionHandlerServer) OnTimerTimeout(ConnectionHandler_OnTimerTimeoutServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTimerTimeout not implemented")
}
func (UnimplementedConnectionHandlerServer) OnReceivedMessages(ConnectionHandler_OnReceivedMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method OnReceivedMessages not implemented")
}
func (UnimplementedConnectionHandlerServer) mustEmbedUnimplementedConnectionHandlerServer() {}

// UnsafeConnectionHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionHandlerServer will
// result in compilation errors.
type UnsafeConnectionHandlerServer interface {
	mustEmbedUnimplementedConnectionHandlerServer()
}

func RegisterConnectionHandlerServer(s grpc.ServiceRegistrar, srv ConnectionHandlerServer) {
	s.RegisterService(&ConnectionHandler_ServiceDesc, srv)
}

func _ConnectionHandler_OnSocketCreated_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnectionHandlerServer).OnSocketCreated(&connectionHandlerOnSocketCreatedServer{stream})
}

type ConnectionHandler_OnSocketCreatedServer interface {
	SendAndClose(*EmptySuccess) error
	Recv() (*SocketCreatedRequest, error)
	grpc.ServerStream
}

type connectionHandlerOnSocketCreatedServer struct {
	grpc.ServerStream
}

func (x *connectionHandlerOnSocketCreatedServer) SendAndClose(m *EmptySuccess) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connectionHandlerOnSocketCreatedServer) Recv() (*SocketCreatedRequest, error) {
	m := new(SocketCreatedRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ConnectionHandler_OnSocketClosed_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnectionHandlerServer).OnSocketClosed(&connectionHandlerOnSocketClosedServer{stream})
}

type ConnectionHandler_OnSocketClosedServer interface {
	SendAndClose(*EmptySuccess) error
	Recv() (*SocketClosedRequest, error)
	grpc.ServerStream
}

type connectionHandlerOnSocketClosedServer struct {
	grpc.ServerStream
}

func (x *connectionHandlerOnSocketClosedServer) SendAndClose(m *EmptySuccess) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connectionHandlerOnSocketClosedServer) Recv() (*SocketClosedRequest, error) {
	m := new(SocketClosedRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ConnectionHandler_OnReceivedBytes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnectionHandlerServer).OnReceivedBytes(&connectionHandlerOnReceivedBytesServer{stream})
}

type ConnectionHandler_OnReceivedBytesServer interface {
	SendAndClose(*EmptySuccess) error
	Recv() (*ReceivedBytesRequest, error)
	grpc.ServerStream
}

type connectionHandlerOnReceivedBytesServer struct {
	grpc.ServerStream
}

func (x *connectionHandlerOnReceivedBytesServer) SendAndClose(m *EmptySuccess) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connectionHandlerOnReceivedBytesServer) Recv() (*ReceivedBytesRequest, error) {
	m := new(ReceivedBytesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ConnectionHandler_OnTimerTimeout_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnectionHandlerServer).OnTimerTimeout(&connectionHandlerOnTimerTimeoutServer{stream})
}

type ConnectionHandler_OnTimerTimeoutServer interface {
	SendAndClose(*EmptySuccess) error
	Recv() (*TimerTimeoutRequest, error)
	grpc.ServerStream
}

type connectionHandlerOnTimerTimeoutServer struct {
	grpc.ServerStream
}

func (x *connectionHandlerOnTimerTimeoutServer) SendAndClose(m *EmptySuccess) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connectionHandlerOnTimerTimeoutServer) Recv() (*TimerTimeoutRequest, error) {
	m := new(TimerTimeoutRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ConnectionHandler_OnReceivedMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnectionHandlerServer).OnReceivedMessages(&connectionHandlerOnReceivedMessagesServer{stream})
}

type ConnectionHandler_OnReceivedMessagesServer interface {
	SendAndClose(*EmptySuccess) error
	Recv() (*ReceivedMessagesRequest, error)
	grpc.ServerStream
}

type connectionHandlerOnReceivedMessagesServer struct {
	grpc.ServerStream
}

func (x *connectionHandlerOnReceivedMessagesServer) SendAndClose(m *EmptySuccess) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connectionHandlerOnReceivedMessagesServer) Recv() (*ReceivedMessagesRequest, error) {
	m := new(ReceivedMessagesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConnectionHandler_ServiceDesc is the grpc.ServiceDesc for ConnectionHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectionHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emqx.exproto.v1.ConnectionHandler",
	HandlerType: (*ConnectionHandlerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OnSocketCreated",
			Handler:       _ConnectionHandler_OnSocketCreated_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "OnSocketClosed",
			Handler:       _ConnectionHandler_OnSocketClosed_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "OnReceivedBytes",
			Handler:       _ConnectionHandler_OnReceivedBytes_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "OnTimerTimeout",
			Handler:       _ConnectionHandler_OnTimerTimeout_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "OnReceivedMessages",
			Handler:       _ConnectionHandler_OnReceivedMessages_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protos/exproto.proto",
}
