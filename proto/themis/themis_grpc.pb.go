// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: themis.proto

package themis

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
	Themis_Span_FullMethodName            = "/themis.Themis/Span"
	Themis_StateSyncEvents_FullMethodName = "/themis.Themis/StateSyncEvents"
)

// ThemisClient is the client API for Themis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThemisClient interface {
	Span(ctx context.Context, in *SpanRequest, opts ...grpc.CallOption) (*SpanResponse, error)
	StateSyncEvents(ctx context.Context, in *StateSyncEventsRequest, opts ...grpc.CallOption) (Themis_StateSyncEventsClient, error)
}

type themisClient struct {
	cc grpc.ClientConnInterface
}

func NewThemisClient(cc grpc.ClientConnInterface) ThemisClient {
	return &themisClient{cc}
}

func (c *themisClient) Span(ctx context.Context, in *SpanRequest, opts ...grpc.CallOption) (*SpanResponse, error) {
	out := new(SpanResponse)
	err := c.cc.Invoke(ctx, Themis_Span_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themisClient) StateSyncEvents(ctx context.Context, in *StateSyncEventsRequest, opts ...grpc.CallOption) (Themis_StateSyncEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Themis_ServiceDesc.Streams[0], Themis_StateSyncEvents_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &themisStateSyncEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Themis_StateSyncEventsClient interface {
	Recv() (*StateSyncEventsResponse, error)
	grpc.ClientStream
}

type themisStateSyncEventsClient struct {
	grpc.ClientStream
}

func (x *themisStateSyncEventsClient) Recv() (*StateSyncEventsResponse, error) {
	m := new(StateSyncEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ThemisServer is the server API for Themis service.
// All implementations must embed UnimplementedThemisServer
// for forward compatibility
type ThemisServer interface {
	Span(context.Context, *SpanRequest) (*SpanResponse, error)
	StateSyncEvents(*StateSyncEventsRequest, Themis_StateSyncEventsServer) error
	mustEmbedUnimplementedThemisServer()
}

// UnimplementedThemisServer must be embedded to have forward compatible implementations.
type UnimplementedThemisServer struct {
}

func (UnimplementedThemisServer) Span(context.Context, *SpanRequest) (*SpanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Span not implemented")
}
func (UnimplementedThemisServer) StateSyncEvents(*StateSyncEventsRequest, Themis_StateSyncEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StateSyncEvents not implemented")
}
func (UnimplementedThemisServer) mustEmbedUnimplementedThemisServer() {}

// UnsafeThemisServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThemisServer will
// result in compilation errors.
type UnsafeThemisServer interface {
	mustEmbedUnimplementedThemisServer()
}

func RegisterThemisServer(s grpc.ServiceRegistrar, srv ThemisServer) {
	s.RegisterService(&Themis_ServiceDesc, srv)
}

func _Themis_Span_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemisServer).Span(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Themis_Span_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemisServer).Span(ctx, req.(*SpanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Themis_StateSyncEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StateSyncEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ThemisServer).StateSyncEvents(m, &themisStateSyncEventsServer{stream})
}

type Themis_StateSyncEventsServer interface {
	Send(*StateSyncEventsResponse) error
	grpc.ServerStream
}

type themisStateSyncEventsServer struct {
	grpc.ServerStream
}

func (x *themisStateSyncEventsServer) Send(m *StateSyncEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Themis_ServiceDesc is the grpc.ServiceDesc for Themis service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Themis_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "themis.Themis",
	HandlerType: (*ThemisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Span",
			Handler:    _Themis_Span_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StateSyncEvents",
			Handler:       _Themis_StateSyncEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "themis.proto",
}