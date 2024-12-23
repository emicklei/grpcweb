// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: test/test.proto

package test

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TestService_PingEmpty_FullMethodName    = "/test.TestService/PingEmpty"
	TestService_Ping_FullMethodName         = "/test.TestService/Ping"
	TestService_PingError_FullMethodName    = "/test.TestService/PingError"
	TestService_PingList_FullMethodName     = "/test.TestService/PingList"
	TestService_PingPongBidi_FullMethodName = "/test.TestService/PingPongBidi"
	TestService_PingStream_FullMethodName   = "/test.TestService/PingStream"
	TestService_Echo_FullMethodName         = "/test.TestService/Echo"
)

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	PingEmpty(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PingResponse], error)
	PingPongBidi(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[PingRequest, PingResponse], error)
	PingStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[PingRequest, PingResponse], error)
	Echo(ctx context.Context, in *TextMessage, opts ...grpc.CallOption) (*TextMessage, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) PingEmpty(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, TestService_PingEmpty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, TestService_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TestService_PingError_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PingResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[0], TestService_PingList_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PingRequest, PingResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TestService_PingListClient = grpc.ServerStreamingClient[PingResponse]

func (c *testServiceClient) PingPongBidi(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[PingRequest, PingResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[1], TestService_PingPongBidi_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PingRequest, PingResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TestService_PingPongBidiClient = grpc.BidiStreamingClient[PingRequest, PingResponse]

func (c *testServiceClient) PingStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[PingRequest, PingResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[2], TestService_PingStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PingRequest, PingResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TestService_PingStreamClient = grpc.ClientStreamingClient[PingRequest, PingResponse]

func (c *testServiceClient) Echo(ctx context.Context, in *TextMessage, opts ...grpc.CallOption) (*TextMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TextMessage)
	err := c.cc.Invoke(ctx, TestService_Echo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility.
type TestServiceServer interface {
	PingEmpty(context.Context, *emptypb.Empty) (*PingResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	PingError(context.Context, *PingRequest) (*emptypb.Empty, error)
	PingList(*PingRequest, grpc.ServerStreamingServer[PingResponse]) error
	PingPongBidi(grpc.BidiStreamingServer[PingRequest, PingResponse]) error
	PingStream(grpc.ClientStreamingServer[PingRequest, PingResponse]) error
	Echo(context.Context, *TextMessage) (*TextMessage, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTestServiceServer struct{}

func (UnimplementedTestServiceServer) PingEmpty(context.Context, *emptypb.Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingEmpty not implemented")
}
func (UnimplementedTestServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedTestServiceServer) PingError(context.Context, *PingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingError not implemented")
}
func (UnimplementedTestServiceServer) PingList(*PingRequest, grpc.ServerStreamingServer[PingResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PingList not implemented")
}
func (UnimplementedTestServiceServer) PingPongBidi(grpc.BidiStreamingServer[PingRequest, PingResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PingPongBidi not implemented")
}
func (UnimplementedTestServiceServer) PingStream(grpc.ClientStreamingServer[PingRequest, PingResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PingStream not implemented")
}
func (UnimplementedTestServiceServer) Echo(context.Context, *TextMessage) (*TextMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}
func (UnimplementedTestServiceServer) testEmbeddedByValue()                     {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	// If the following call pancis, it indicates UnimplementedTestServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_PingEmpty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingEmpty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_PingEmpty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingEmpty(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_PingError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingError(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).PingList(m, &grpc.GenericServerStream[PingRequest, PingResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TestService_PingListServer = grpc.ServerStreamingServer[PingResponse]

func _TestService_PingPongBidi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).PingPongBidi(&grpc.GenericServerStream[PingRequest, PingResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TestService_PingPongBidiServer = grpc.BidiStreamingServer[PingRequest, PingResponse]

func _TestService_PingStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).PingStream(&grpc.GenericServerStream[PingRequest, PingResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TestService_PingStreamServer = grpc.ClientStreamingServer[PingRequest, PingResponse]

func _TestService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Echo(ctx, req.(*TextMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingEmpty",
			Handler:    _TestService_PingEmpty_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _TestService_Ping_Handler,
		},
		{
			MethodName: "PingError",
			Handler:    _TestService_PingError_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _TestService_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingList",
			Handler:       _TestService_PingList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PingPongBidi",
			Handler:       _TestService_PingPongBidi_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PingStream",
			Handler:       _TestService_PingStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "test/test.proto",
}

const (
	TestUtilService_ContinueStream_FullMethodName    = "/test.TestUtilService/ContinueStream"
	TestUtilService_CheckStreamClosed_FullMethodName = "/test.TestUtilService/CheckStreamClosed"
)

// TestUtilServiceClient is the client API for TestUtilService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestUtilServiceClient interface {
	ContinueStream(ctx context.Context, in *ContinueStreamRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckStreamClosed(ctx context.Context, in *CheckStreamClosedRequest, opts ...grpc.CallOption) (*CheckStreamClosedResponse, error)
}

type testUtilServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestUtilServiceClient(cc grpc.ClientConnInterface) TestUtilServiceClient {
	return &testUtilServiceClient{cc}
}

func (c *testUtilServiceClient) ContinueStream(ctx context.Context, in *ContinueStreamRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TestUtilService_ContinueStream_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testUtilServiceClient) CheckStreamClosed(ctx context.Context, in *CheckStreamClosedRequest, opts ...grpc.CallOption) (*CheckStreamClosedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckStreamClosedResponse)
	err := c.cc.Invoke(ctx, TestUtilService_CheckStreamClosed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestUtilServiceServer is the server API for TestUtilService service.
// All implementations must embed UnimplementedTestUtilServiceServer
// for forward compatibility.
type TestUtilServiceServer interface {
	ContinueStream(context.Context, *ContinueStreamRequest) (*emptypb.Empty, error)
	CheckStreamClosed(context.Context, *CheckStreamClosedRequest) (*CheckStreamClosedResponse, error)
	mustEmbedUnimplementedTestUtilServiceServer()
}

// UnimplementedTestUtilServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTestUtilServiceServer struct{}

func (UnimplementedTestUtilServiceServer) ContinueStream(context.Context, *ContinueStreamRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContinueStream not implemented")
}
func (UnimplementedTestUtilServiceServer) CheckStreamClosed(context.Context, *CheckStreamClosedRequest) (*CheckStreamClosedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStreamClosed not implemented")
}
func (UnimplementedTestUtilServiceServer) mustEmbedUnimplementedTestUtilServiceServer() {}
func (UnimplementedTestUtilServiceServer) testEmbeddedByValue()                         {}

// UnsafeTestUtilServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestUtilServiceServer will
// result in compilation errors.
type UnsafeTestUtilServiceServer interface {
	mustEmbedUnimplementedTestUtilServiceServer()
}

func RegisterTestUtilServiceServer(s grpc.ServiceRegistrar, srv TestUtilServiceServer) {
	// If the following call pancis, it indicates UnimplementedTestUtilServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TestUtilService_ServiceDesc, srv)
}

func _TestUtilService_ContinueStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContinueStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestUtilServiceServer).ContinueStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestUtilService_ContinueStream_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestUtilServiceServer).ContinueStream(ctx, req.(*ContinueStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestUtilService_CheckStreamClosed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckStreamClosedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestUtilServiceServer).CheckStreamClosed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestUtilService_CheckStreamClosed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestUtilServiceServer).CheckStreamClosed(ctx, req.(*CheckStreamClosedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestUtilService_ServiceDesc is the grpc.ServiceDesc for TestUtilService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestUtilService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.TestUtilService",
	HandlerType: (*TestUtilServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContinueStream",
			Handler:    _TestUtilService_ContinueStream_Handler,
		},
		{
			MethodName: "CheckStreamClosed",
			Handler:    _TestUtilService_CheckStreamClosed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test/test.proto",
}

const (
	FailService_NonExistant_FullMethodName = "/test.FailService/NonExistant"
)

// FailServiceClient is the client API for FailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FailServiceClient interface {
	NonExistant(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type failServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFailServiceClient(cc grpc.ClientConnInterface) FailServiceClient {
	return &failServiceClient{cc}
}

func (c *failServiceClient) NonExistant(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, FailService_NonExistant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FailServiceServer is the server API for FailService service.
// All implementations must embed UnimplementedFailServiceServer
// for forward compatibility.
type FailServiceServer interface {
	NonExistant(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedFailServiceServer()
}

// UnimplementedFailServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFailServiceServer struct{}

func (UnimplementedFailServiceServer) NonExistant(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NonExistant not implemented")
}
func (UnimplementedFailServiceServer) mustEmbedUnimplementedFailServiceServer() {}
func (UnimplementedFailServiceServer) testEmbeddedByValue()                     {}

// UnsafeFailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FailServiceServer will
// result in compilation errors.
type UnsafeFailServiceServer interface {
	mustEmbedUnimplementedFailServiceServer()
}

func RegisterFailServiceServer(s grpc.ServiceRegistrar, srv FailServiceServer) {
	// If the following call pancis, it indicates UnimplementedFailServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FailService_ServiceDesc, srv)
}

func _FailService_NonExistant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FailServiceServer).NonExistant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FailService_NonExistant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FailServiceServer).NonExistant(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FailService_ServiceDesc is the grpc.ServiceDesc for FailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.FailService",
	HandlerType: (*FailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NonExistant",
			Handler:    _FailService_NonExistant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test/test.proto",
}
