// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: stream.proto

package stream

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
	StreamEcho_RunServerStream_FullMethodName = "/stream.StreamEcho/RunServerStream"
	StreamEcho_RunClientStream_FullMethodName = "/stream.StreamEcho/RunClientStream"
	StreamEcho_RunDuplexStream_FullMethodName = "/stream.StreamEcho/RunDuplexStream"
)

// StreamEchoClient is the client API for StreamEcho service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamEchoClient interface {
	RunServerStream(ctx context.Context, in *RequestPayload, opts ...grpc.CallOption) (StreamEcho_RunServerStreamClient, error)
	RunClientStream(ctx context.Context, opts ...grpc.CallOption) (StreamEcho_RunClientStreamClient, error)
	RunDuplexStream(ctx context.Context, opts ...grpc.CallOption) (StreamEcho_RunDuplexStreamClient, error)
}

type streamEchoClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamEchoClient(cc grpc.ClientConnInterface) StreamEchoClient {
	return &streamEchoClient{cc}
}

func (c *streamEchoClient) RunServerStream(ctx context.Context, in *RequestPayload, opts ...grpc.CallOption) (StreamEcho_RunServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamEcho_ServiceDesc.Streams[0], StreamEcho_RunServerStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamEchoRunServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamEcho_RunServerStreamClient interface {
	Recv() (*ResponsePayload, error)
	grpc.ClientStream
}

type streamEchoRunServerStreamClient struct {
	grpc.ClientStream
}

func (x *streamEchoRunServerStreamClient) Recv() (*ResponsePayload, error) {
	m := new(ResponsePayload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamEchoClient) RunClientStream(ctx context.Context, opts ...grpc.CallOption) (StreamEcho_RunClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamEcho_ServiceDesc.Streams[1], StreamEcho_RunClientStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamEchoRunClientStreamClient{stream}
	return x, nil
}

type StreamEcho_RunClientStreamClient interface {
	Send(*RequestPayload) error
	CloseAndRecv() (*ArrResponsePayload, error)
	grpc.ClientStream
}

type streamEchoRunClientStreamClient struct {
	grpc.ClientStream
}

func (x *streamEchoRunClientStreamClient) Send(m *RequestPayload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamEchoRunClientStreamClient) CloseAndRecv() (*ArrResponsePayload, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ArrResponsePayload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamEchoClient) RunDuplexStream(ctx context.Context, opts ...grpc.CallOption) (StreamEcho_RunDuplexStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamEcho_ServiceDesc.Streams[2], StreamEcho_RunDuplexStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamEchoRunDuplexStreamClient{stream}
	return x, nil
}

type StreamEcho_RunDuplexStreamClient interface {
	Send(*RequestPayload) error
	Recv() (*ResponsePayload, error)
	grpc.ClientStream
}

type streamEchoRunDuplexStreamClient struct {
	grpc.ClientStream
}

func (x *streamEchoRunDuplexStreamClient) Send(m *RequestPayload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamEchoRunDuplexStreamClient) Recv() (*ResponsePayload, error) {
	m := new(ResponsePayload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamEchoServer is the server API for StreamEcho service.
// All implementations must embed UnimplementedStreamEchoServer
// for forward compatibility
type StreamEchoServer interface {
	RunServerStream(*RequestPayload, StreamEcho_RunServerStreamServer) error
	RunClientStream(StreamEcho_RunClientStreamServer) error
	RunDuplexStream(StreamEcho_RunDuplexStreamServer) error
	mustEmbedUnimplementedStreamEchoServer()
}

// UnimplementedStreamEchoServer must be embedded to have forward compatible implementations.
type UnimplementedStreamEchoServer struct {
}

func (UnimplementedStreamEchoServer) RunServerStream(*RequestPayload, StreamEcho_RunServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method RunServerStream not implemented")
}
func (UnimplementedStreamEchoServer) RunClientStream(StreamEcho_RunClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method RunClientStream not implemented")
}
func (UnimplementedStreamEchoServer) RunDuplexStream(StreamEcho_RunDuplexStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method RunDuplexStream not implemented")
}
func (UnimplementedStreamEchoServer) mustEmbedUnimplementedStreamEchoServer() {}

// UnsafeStreamEchoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamEchoServer will
// result in compilation errors.
type UnsafeStreamEchoServer interface {
	mustEmbedUnimplementedStreamEchoServer()
}

func RegisterStreamEchoServer(s grpc.ServiceRegistrar, srv StreamEchoServer) {
	s.RegisterService(&StreamEcho_ServiceDesc, srv)
}

func _StreamEcho_RunServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestPayload)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamEchoServer).RunServerStream(m, &streamEchoRunServerStreamServer{stream})
}

type StreamEcho_RunServerStreamServer interface {
	Send(*ResponsePayload) error
	grpc.ServerStream
}

type streamEchoRunServerStreamServer struct {
	grpc.ServerStream
}

func (x *streamEchoRunServerStreamServer) Send(m *ResponsePayload) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamEcho_RunClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamEchoServer).RunClientStream(&streamEchoRunClientStreamServer{stream})
}

type StreamEcho_RunClientStreamServer interface {
	SendAndClose(*ArrResponsePayload) error
	Recv() (*RequestPayload, error)
	grpc.ServerStream
}

type streamEchoRunClientStreamServer struct {
	grpc.ServerStream
}

func (x *streamEchoRunClientStreamServer) SendAndClose(m *ArrResponsePayload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamEchoRunClientStreamServer) Recv() (*RequestPayload, error) {
	m := new(RequestPayload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamEcho_RunDuplexStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamEchoServer).RunDuplexStream(&streamEchoRunDuplexStreamServer{stream})
}

type StreamEcho_RunDuplexStreamServer interface {
	Send(*ResponsePayload) error
	Recv() (*RequestPayload, error)
	grpc.ServerStream
}

type streamEchoRunDuplexStreamServer struct {
	grpc.ServerStream
}

func (x *streamEchoRunDuplexStreamServer) Send(m *ResponsePayload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamEchoRunDuplexStreamServer) Recv() (*RequestPayload, error) {
	m := new(RequestPayload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamEcho_ServiceDesc is the grpc.ServiceDesc for StreamEcho service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamEcho_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.StreamEcho",
	HandlerType: (*StreamEchoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunServerStream",
			Handler:       _StreamEcho_RunServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RunClientStream",
			Handler:       _StreamEcho_RunClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RunDuplexStream",
			Handler:       _StreamEcho_RunDuplexStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
