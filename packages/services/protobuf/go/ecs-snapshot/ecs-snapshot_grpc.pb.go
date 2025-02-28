// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: proto/ecs-snapshot.proto

package ecs_snapshot

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

// ECSStateSnapshotServiceClient is the client API for ECSStateSnapshotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ECSStateSnapshotServiceClient interface {
	// Requests the latest ECS state.
	GetStateLatest(ctx context.Context, in *ECSStateRequestLatest, opts ...grpc.CallOption) (*ECSStateReply, error)
	// Requests the latest ECS state in stream format, which will chunk the state.
	GetStateLatestStream(ctx context.Context, in *ECSStateRequestLatestStream, opts ...grpc.CallOption) (ECSStateSnapshotService_GetStateLatestStreamClient, error)
	// Requests the latest ECS state in stream format, which will chunk the state.
	//
	// V2 version optimized to return entities as raw bytes.
	GetStateLatestStreamV2(ctx context.Context, in *ECSStateRequestLatestStream, opts ...grpc.CallOption) (ECSStateSnapshotService_GetStateLatestStreamV2Client, error)
	// Requests the latest ECS state, with aditional pruning.
	GetStateLatestStreamPruned(ctx context.Context, in *ECSStateRequestLatestStreamPruned, opts ...grpc.CallOption) (ECSStateSnapshotService_GetStateLatestStreamPrunedClient, error)
	// Requests the latest ECS state, with aditional pruning.
	//
	// V2 version optimized to return entities as raw bytes.
	GetStateLatestStreamPrunedV2(ctx context.Context, in *ECSStateRequestLatestStreamPruned, opts ...grpc.CallOption) (ECSStateSnapshotService_GetStateLatestStreamPrunedV2Client, error)
	// Requests the latest block number based on the latest ECS state.
	GetStateBlockLatest(ctx context.Context, in *ECSStateBlockRequestLatest, opts ...grpc.CallOption) (*ECSStateBlockReply, error)
	// Requests the ECS state at specific block.
	GetStateAtBlock(ctx context.Context, in *ECSStateRequestAtBlock, opts ...grpc.CallOption) (*ECSStateReply, error)
	// Requests a list of known worlds based on chain state.
	GetWorlds(ctx context.Context, in *WorldsRequest, opts ...grpc.CallOption) (*Worlds, error)
}

type eCSStateSnapshotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewECSStateSnapshotServiceClient(cc grpc.ClientConnInterface) ECSStateSnapshotServiceClient {
	return &eCSStateSnapshotServiceClient{cc}
}

func (c *eCSStateSnapshotServiceClient) GetStateLatest(ctx context.Context, in *ECSStateRequestLatest, opts ...grpc.CallOption) (*ECSStateReply, error) {
	out := new(ECSStateReply)
	err := c.cc.Invoke(ctx, "/ecssnapshot.ECSStateSnapshotService/GetStateLatest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eCSStateSnapshotServiceClient) GetStateLatestStream(ctx context.Context, in *ECSStateRequestLatestStream, opts ...grpc.CallOption) (ECSStateSnapshotService_GetStateLatestStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ECSStateSnapshotService_ServiceDesc.Streams[0], "/ecssnapshot.ECSStateSnapshotService/GetStateLatestStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &eCSStateSnapshotServiceGetStateLatestStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ECSStateSnapshotService_GetStateLatestStreamClient interface {
	Recv() (*ECSStateReply, error)
	grpc.ClientStream
}

type eCSStateSnapshotServiceGetStateLatestStreamClient struct {
	grpc.ClientStream
}

func (x *eCSStateSnapshotServiceGetStateLatestStreamClient) Recv() (*ECSStateReply, error) {
	m := new(ECSStateReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eCSStateSnapshotServiceClient) GetStateLatestStreamV2(ctx context.Context, in *ECSStateRequestLatestStream, opts ...grpc.CallOption) (ECSStateSnapshotService_GetStateLatestStreamV2Client, error) {
	stream, err := c.cc.NewStream(ctx, &ECSStateSnapshotService_ServiceDesc.Streams[1], "/ecssnapshot.ECSStateSnapshotService/GetStateLatestStreamV2", opts...)
	if err != nil {
		return nil, err
	}
	x := &eCSStateSnapshotServiceGetStateLatestStreamV2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ECSStateSnapshotService_GetStateLatestStreamV2Client interface {
	Recv() (*ECSStateReplyV2, error)
	grpc.ClientStream
}

type eCSStateSnapshotServiceGetStateLatestStreamV2Client struct {
	grpc.ClientStream
}

func (x *eCSStateSnapshotServiceGetStateLatestStreamV2Client) Recv() (*ECSStateReplyV2, error) {
	m := new(ECSStateReplyV2)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eCSStateSnapshotServiceClient) GetStateLatestStreamPruned(ctx context.Context, in *ECSStateRequestLatestStreamPruned, opts ...grpc.CallOption) (ECSStateSnapshotService_GetStateLatestStreamPrunedClient, error) {
	stream, err := c.cc.NewStream(ctx, &ECSStateSnapshotService_ServiceDesc.Streams[2], "/ecssnapshot.ECSStateSnapshotService/GetStateLatestStreamPruned", opts...)
	if err != nil {
		return nil, err
	}
	x := &eCSStateSnapshotServiceGetStateLatestStreamPrunedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ECSStateSnapshotService_GetStateLatestStreamPrunedClient interface {
	Recv() (*ECSStateReply, error)
	grpc.ClientStream
}

type eCSStateSnapshotServiceGetStateLatestStreamPrunedClient struct {
	grpc.ClientStream
}

func (x *eCSStateSnapshotServiceGetStateLatestStreamPrunedClient) Recv() (*ECSStateReply, error) {
	m := new(ECSStateReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eCSStateSnapshotServiceClient) GetStateLatestStreamPrunedV2(ctx context.Context, in *ECSStateRequestLatestStreamPruned, opts ...grpc.CallOption) (ECSStateSnapshotService_GetStateLatestStreamPrunedV2Client, error) {
	stream, err := c.cc.NewStream(ctx, &ECSStateSnapshotService_ServiceDesc.Streams[3], "/ecssnapshot.ECSStateSnapshotService/GetStateLatestStreamPrunedV2", opts...)
	if err != nil {
		return nil, err
	}
	x := &eCSStateSnapshotServiceGetStateLatestStreamPrunedV2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ECSStateSnapshotService_GetStateLatestStreamPrunedV2Client interface {
	Recv() (*ECSStateReplyV2, error)
	grpc.ClientStream
}

type eCSStateSnapshotServiceGetStateLatestStreamPrunedV2Client struct {
	grpc.ClientStream
}

func (x *eCSStateSnapshotServiceGetStateLatestStreamPrunedV2Client) Recv() (*ECSStateReplyV2, error) {
	m := new(ECSStateReplyV2)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eCSStateSnapshotServiceClient) GetStateBlockLatest(ctx context.Context, in *ECSStateBlockRequestLatest, opts ...grpc.CallOption) (*ECSStateBlockReply, error) {
	out := new(ECSStateBlockReply)
	err := c.cc.Invoke(ctx, "/ecssnapshot.ECSStateSnapshotService/GetStateBlockLatest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eCSStateSnapshotServiceClient) GetStateAtBlock(ctx context.Context, in *ECSStateRequestAtBlock, opts ...grpc.CallOption) (*ECSStateReply, error) {
	out := new(ECSStateReply)
	err := c.cc.Invoke(ctx, "/ecssnapshot.ECSStateSnapshotService/GetStateAtBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eCSStateSnapshotServiceClient) GetWorlds(ctx context.Context, in *WorldsRequest, opts ...grpc.CallOption) (*Worlds, error) {
	out := new(Worlds)
	err := c.cc.Invoke(ctx, "/ecssnapshot.ECSStateSnapshotService/GetWorlds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ECSStateSnapshotServiceServer is the server API for ECSStateSnapshotService service.
// All implementations must embed UnimplementedECSStateSnapshotServiceServer
// for forward compatibility
type ECSStateSnapshotServiceServer interface {
	// Requests the latest ECS state.
	GetStateLatest(context.Context, *ECSStateRequestLatest) (*ECSStateReply, error)
	// Requests the latest ECS state in stream format, which will chunk the state.
	GetStateLatestStream(*ECSStateRequestLatestStream, ECSStateSnapshotService_GetStateLatestStreamServer) error
	// Requests the latest ECS state in stream format, which will chunk the state.
	//
	// V2 version optimized to return entities as raw bytes.
	GetStateLatestStreamV2(*ECSStateRequestLatestStream, ECSStateSnapshotService_GetStateLatestStreamV2Server) error
	// Requests the latest ECS state, with aditional pruning.
	GetStateLatestStreamPruned(*ECSStateRequestLatestStreamPruned, ECSStateSnapshotService_GetStateLatestStreamPrunedServer) error
	// Requests the latest ECS state, with aditional pruning.
	//
	// V2 version optimized to return entities as raw bytes.
	GetStateLatestStreamPrunedV2(*ECSStateRequestLatestStreamPruned, ECSStateSnapshotService_GetStateLatestStreamPrunedV2Server) error
	// Requests the latest block number based on the latest ECS state.
	GetStateBlockLatest(context.Context, *ECSStateBlockRequestLatest) (*ECSStateBlockReply, error)
	// Requests the ECS state at specific block.
	GetStateAtBlock(context.Context, *ECSStateRequestAtBlock) (*ECSStateReply, error)
	// Requests a list of known worlds based on chain state.
	GetWorlds(context.Context, *WorldsRequest) (*Worlds, error)
	mustEmbedUnimplementedECSStateSnapshotServiceServer()
}

// UnimplementedECSStateSnapshotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedECSStateSnapshotServiceServer struct {
}

func (UnimplementedECSStateSnapshotServiceServer) GetStateLatest(context.Context, *ECSStateRequestLatest) (*ECSStateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStateLatest not implemented")
}
func (UnimplementedECSStateSnapshotServiceServer) GetStateLatestStream(*ECSStateRequestLatestStream, ECSStateSnapshotService_GetStateLatestStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStateLatestStream not implemented")
}
func (UnimplementedECSStateSnapshotServiceServer) GetStateLatestStreamV2(*ECSStateRequestLatestStream, ECSStateSnapshotService_GetStateLatestStreamV2Server) error {
	return status.Errorf(codes.Unimplemented, "method GetStateLatestStreamV2 not implemented")
}
func (UnimplementedECSStateSnapshotServiceServer) GetStateLatestStreamPruned(*ECSStateRequestLatestStreamPruned, ECSStateSnapshotService_GetStateLatestStreamPrunedServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStateLatestStreamPruned not implemented")
}
func (UnimplementedECSStateSnapshotServiceServer) GetStateLatestStreamPrunedV2(*ECSStateRequestLatestStreamPruned, ECSStateSnapshotService_GetStateLatestStreamPrunedV2Server) error {
	return status.Errorf(codes.Unimplemented, "method GetStateLatestStreamPrunedV2 not implemented")
}
func (UnimplementedECSStateSnapshotServiceServer) GetStateBlockLatest(context.Context, *ECSStateBlockRequestLatest) (*ECSStateBlockReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStateBlockLatest not implemented")
}
func (UnimplementedECSStateSnapshotServiceServer) GetStateAtBlock(context.Context, *ECSStateRequestAtBlock) (*ECSStateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStateAtBlock not implemented")
}
func (UnimplementedECSStateSnapshotServiceServer) GetWorlds(context.Context, *WorldsRequest) (*Worlds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorlds not implemented")
}
func (UnimplementedECSStateSnapshotServiceServer) mustEmbedUnimplementedECSStateSnapshotServiceServer() {
}

// UnsafeECSStateSnapshotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ECSStateSnapshotServiceServer will
// result in compilation errors.
type UnsafeECSStateSnapshotServiceServer interface {
	mustEmbedUnimplementedECSStateSnapshotServiceServer()
}

func RegisterECSStateSnapshotServiceServer(s grpc.ServiceRegistrar, srv ECSStateSnapshotServiceServer) {
	s.RegisterService(&ECSStateSnapshotService_ServiceDesc, srv)
}

func _ECSStateSnapshotService_GetStateLatest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ECSStateRequestLatest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ECSStateSnapshotServiceServer).GetStateLatest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecssnapshot.ECSStateSnapshotService/GetStateLatest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ECSStateSnapshotServiceServer).GetStateLatest(ctx, req.(*ECSStateRequestLatest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ECSStateSnapshotService_GetStateLatestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ECSStateRequestLatestStream)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ECSStateSnapshotServiceServer).GetStateLatestStream(m, &eCSStateSnapshotServiceGetStateLatestStreamServer{stream})
}

type ECSStateSnapshotService_GetStateLatestStreamServer interface {
	Send(*ECSStateReply) error
	grpc.ServerStream
}

type eCSStateSnapshotServiceGetStateLatestStreamServer struct {
	grpc.ServerStream
}

func (x *eCSStateSnapshotServiceGetStateLatestStreamServer) Send(m *ECSStateReply) error {
	return x.ServerStream.SendMsg(m)
}

func _ECSStateSnapshotService_GetStateLatestStreamV2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ECSStateRequestLatestStream)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ECSStateSnapshotServiceServer).GetStateLatestStreamV2(m, &eCSStateSnapshotServiceGetStateLatestStreamV2Server{stream})
}

type ECSStateSnapshotService_GetStateLatestStreamV2Server interface {
	Send(*ECSStateReplyV2) error
	grpc.ServerStream
}

type eCSStateSnapshotServiceGetStateLatestStreamV2Server struct {
	grpc.ServerStream
}

func (x *eCSStateSnapshotServiceGetStateLatestStreamV2Server) Send(m *ECSStateReplyV2) error {
	return x.ServerStream.SendMsg(m)
}

func _ECSStateSnapshotService_GetStateLatestStreamPruned_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ECSStateRequestLatestStreamPruned)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ECSStateSnapshotServiceServer).GetStateLatestStreamPruned(m, &eCSStateSnapshotServiceGetStateLatestStreamPrunedServer{stream})
}

type ECSStateSnapshotService_GetStateLatestStreamPrunedServer interface {
	Send(*ECSStateReply) error
	grpc.ServerStream
}

type eCSStateSnapshotServiceGetStateLatestStreamPrunedServer struct {
	grpc.ServerStream
}

func (x *eCSStateSnapshotServiceGetStateLatestStreamPrunedServer) Send(m *ECSStateReply) error {
	return x.ServerStream.SendMsg(m)
}

func _ECSStateSnapshotService_GetStateLatestStreamPrunedV2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ECSStateRequestLatestStreamPruned)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ECSStateSnapshotServiceServer).GetStateLatestStreamPrunedV2(m, &eCSStateSnapshotServiceGetStateLatestStreamPrunedV2Server{stream})
}

type ECSStateSnapshotService_GetStateLatestStreamPrunedV2Server interface {
	Send(*ECSStateReplyV2) error
	grpc.ServerStream
}

type eCSStateSnapshotServiceGetStateLatestStreamPrunedV2Server struct {
	grpc.ServerStream
}

func (x *eCSStateSnapshotServiceGetStateLatestStreamPrunedV2Server) Send(m *ECSStateReplyV2) error {
	return x.ServerStream.SendMsg(m)
}

func _ECSStateSnapshotService_GetStateBlockLatest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ECSStateBlockRequestLatest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ECSStateSnapshotServiceServer).GetStateBlockLatest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecssnapshot.ECSStateSnapshotService/GetStateBlockLatest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ECSStateSnapshotServiceServer).GetStateBlockLatest(ctx, req.(*ECSStateBlockRequestLatest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ECSStateSnapshotService_GetStateAtBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ECSStateRequestAtBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ECSStateSnapshotServiceServer).GetStateAtBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecssnapshot.ECSStateSnapshotService/GetStateAtBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ECSStateSnapshotServiceServer).GetStateAtBlock(ctx, req.(*ECSStateRequestAtBlock))
	}
	return interceptor(ctx, in, info, handler)
}

func _ECSStateSnapshotService_GetWorlds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ECSStateSnapshotServiceServer).GetWorlds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecssnapshot.ECSStateSnapshotService/GetWorlds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ECSStateSnapshotServiceServer).GetWorlds(ctx, req.(*WorldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ECSStateSnapshotService_ServiceDesc is the grpc.ServiceDesc for ECSStateSnapshotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ECSStateSnapshotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecssnapshot.ECSStateSnapshotService",
	HandlerType: (*ECSStateSnapshotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStateLatest",
			Handler:    _ECSStateSnapshotService_GetStateLatest_Handler,
		},
		{
			MethodName: "GetStateBlockLatest",
			Handler:    _ECSStateSnapshotService_GetStateBlockLatest_Handler,
		},
		{
			MethodName: "GetStateAtBlock",
			Handler:    _ECSStateSnapshotService_GetStateAtBlock_Handler,
		},
		{
			MethodName: "GetWorlds",
			Handler:    _ECSStateSnapshotService_GetWorlds_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStateLatestStream",
			Handler:       _ECSStateSnapshotService_GetStateLatestStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetStateLatestStreamV2",
			Handler:       _ECSStateSnapshotService_GetStateLatestStreamV2_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetStateLatestStreamPruned",
			Handler:       _ECSStateSnapshotService_GetStateLatestStreamPruned_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetStateLatestStreamPrunedV2",
			Handler:       _ECSStateSnapshotService_GetStateLatestStreamPrunedV2_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/ecs-snapshot.proto",
}
