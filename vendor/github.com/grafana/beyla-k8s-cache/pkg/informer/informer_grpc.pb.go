// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/informer.proto

package informer

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

// EventStreamServiceClient is the client API for EventStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventStreamServiceClient interface {
	// Subscribe allows clients to subscribe to a stream of events.
	Subscribe(ctx context.Context, in *SubscribeMessage, opts ...grpc.CallOption) (EventStreamService_SubscribeClient, error)
}

type eventStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventStreamServiceClient(cc grpc.ClientConnInterface) EventStreamServiceClient {
	return &eventStreamServiceClient{cc}
}

func (c *eventStreamServiceClient) Subscribe(ctx context.Context, in *SubscribeMessage, opts ...grpc.CallOption) (EventStreamService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventStreamService_ServiceDesc.Streams[0], "/informer.EventStreamService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventStreamServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventStreamService_SubscribeClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventStreamServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *eventStreamServiceSubscribeClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventStreamServiceServer is the server API for EventStreamService service.
// All implementations must embed UnimplementedEventStreamServiceServer
// for forward compatibility
type EventStreamServiceServer interface {
	// Subscribe allows clients to subscribe to a stream of events.
	Subscribe(*SubscribeMessage, EventStreamService_SubscribeServer) error
	mustEmbedUnimplementedEventStreamServiceServer()
}

// UnimplementedEventStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventStreamServiceServer struct {
}

func (UnimplementedEventStreamServiceServer) Subscribe(*SubscribeMessage, EventStreamService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedEventStreamServiceServer) mustEmbedUnimplementedEventStreamServiceServer() {}

// UnsafeEventStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventStreamServiceServer will
// result in compilation errors.
type UnsafeEventStreamServiceServer interface {
	mustEmbedUnimplementedEventStreamServiceServer()
}

func RegisterEventStreamServiceServer(s grpc.ServiceRegistrar, srv EventStreamServiceServer) {
	s.RegisterService(&EventStreamService_ServiceDesc, srv)
}

func _EventStreamService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventStreamServiceServer).Subscribe(m, &eventStreamServiceSubscribeServer{stream})
}

type EventStreamService_SubscribeServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventStreamServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *eventStreamServiceSubscribeServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

// EventStreamService_ServiceDesc is the grpc.ServiceDesc for EventStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "informer.EventStreamService",
	HandlerType: (*EventStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _EventStreamService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/informer.proto",
}
