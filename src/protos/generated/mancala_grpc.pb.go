// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mancala.proto

package protos

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

// MancalaServiceClient is the client API for MancalaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MancalaServiceClient interface {
	GameHandshake(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*HandshakeResponse, error)
	MakeMove(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*MoveResponse, error)
	RequestUpdate(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	NodeGameStatus(ctx context.Context, in *GameStatusRequest, opts ...grpc.CallOption) (*GameStatusResponse, error)
}

type mancalaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMancalaServiceClient(cc grpc.ClientConnInterface) MancalaServiceClient {
	return &mancalaServiceClient{cc}
}

func (c *mancalaServiceClient) GameHandshake(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*HandshakeResponse, error) {
	out := new(HandshakeResponse)
	err := c.cc.Invoke(ctx, "/mancala.MancalaService/GameHandshake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mancalaServiceClient) MakeMove(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*MoveResponse, error) {
	out := new(MoveResponse)
	err := c.cc.Invoke(ctx, "/mancala.MancalaService/MakeMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mancalaServiceClient) RequestUpdate(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/mancala.MancalaService/RequestUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mancalaServiceClient) NodeGameStatus(ctx context.Context, in *GameStatusRequest, opts ...grpc.CallOption) (*GameStatusResponse, error) {
	out := new(GameStatusResponse)
	err := c.cc.Invoke(ctx, "/mancala.MancalaService/NodeGameStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MancalaServiceServer is the server API for MancalaService service.
// All implementations must embed UnimplementedMancalaServiceServer
// for forward compatibility
type MancalaServiceServer interface {
	GameHandshake(context.Context, *HandshakeRequest) (*HandshakeResponse, error)
	MakeMove(context.Context, *MoveRequest) (*MoveResponse, error)
	RequestUpdate(context.Context, *UpdateRequest) (*UpdateResponse, error)
	NodeGameStatus(context.Context, *GameStatusRequest) (*GameStatusResponse, error)
	mustEmbedUnimplementedMancalaServiceServer()
}

// UnimplementedMancalaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMancalaServiceServer struct {
}

func (UnimplementedMancalaServiceServer) GameHandshake(context.Context, *HandshakeRequest) (*HandshakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GameHandshake not implemented")
}
func (UnimplementedMancalaServiceServer) MakeMove(context.Context, *MoveRequest) (*MoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeMove not implemented")
}
func (UnimplementedMancalaServiceServer) RequestUpdate(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestUpdate not implemented")
}
func (UnimplementedMancalaServiceServer) NodeGameStatus(context.Context, *GameStatusRequest) (*GameStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGameStatus not implemented")
}
func (UnimplementedMancalaServiceServer) mustEmbedUnimplementedMancalaServiceServer() {}

// UnsafeMancalaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MancalaServiceServer will
// result in compilation errors.
type UnsafeMancalaServiceServer interface {
	mustEmbedUnimplementedMancalaServiceServer()
}

func RegisterMancalaServiceServer(s grpc.ServiceRegistrar, srv MancalaServiceServer) {
	s.RegisterService(&MancalaService_ServiceDesc, srv)
}

func _MancalaService_GameHandshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandshakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MancalaServiceServer).GameHandshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mancala.MancalaService/GameHandshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MancalaServiceServer).GameHandshake(ctx, req.(*HandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MancalaService_MakeMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MancalaServiceServer).MakeMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mancala.MancalaService/MakeMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MancalaServiceServer).MakeMove(ctx, req.(*MoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MancalaService_RequestUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MancalaServiceServer).RequestUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mancala.MancalaService/RequestUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MancalaServiceServer).RequestUpdate(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MancalaService_NodeGameStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MancalaServiceServer).NodeGameStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mancala.MancalaService/NodeGameStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MancalaServiceServer).NodeGameStatus(ctx, req.(*GameStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MancalaService_ServiceDesc is the grpc.ServiceDesc for MancalaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MancalaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mancala.MancalaService",
	HandlerType: (*MancalaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GameHandshake",
			Handler:    _MancalaService_GameHandshake_Handler,
		},
		{
			MethodName: "MakeMove",
			Handler:    _MancalaService_MakeMove_Handler,
		},
		{
			MethodName: "RequestUpdate",
			Handler:    _MancalaService_RequestUpdate_Handler,
		},
		{
			MethodName: "NodeGameStatus",
			Handler:    _MancalaService_NodeGameStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mancala.proto",
}
