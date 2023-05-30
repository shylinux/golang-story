// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// SpaceServiceClient is the client API for SpaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpaceServiceClient interface {
	Create(ctx context.Context, in *SpaceCreateRequest, opts ...grpc.CallOption) (*SpaceCreateReply, error)
	Remove(ctx context.Context, in *SpaceRemoveRequest, opts ...grpc.CallOption) (*SpaceRemoveReply, error)
	Info(ctx context.Context, in *SpaceInfoRequest, opts ...grpc.CallOption) (*SpaceInfoReply, error)
	List(ctx context.Context, in *SpaceListRequest, opts ...grpc.CallOption) (*SpaceListReply, error)
}

type spaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpaceServiceClient(cc grpc.ClientConnInterface) SpaceServiceClient {
	return &spaceServiceClient{cc}
}

func (c *spaceServiceClient) Create(ctx context.Context, in *SpaceCreateRequest, opts ...grpc.CallOption) (*SpaceCreateReply, error) {
	out := new(SpaceCreateReply)
	err := c.cc.Invoke(ctx, "/demo.space.SpaceService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) Remove(ctx context.Context, in *SpaceRemoveRequest, opts ...grpc.CallOption) (*SpaceRemoveReply, error) {
	out := new(SpaceRemoveReply)
	err := c.cc.Invoke(ctx, "/demo.space.SpaceService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) Info(ctx context.Context, in *SpaceInfoRequest, opts ...grpc.CallOption) (*SpaceInfoReply, error) {
	out := new(SpaceInfoReply)
	err := c.cc.Invoke(ctx, "/demo.space.SpaceService/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) List(ctx context.Context, in *SpaceListRequest, opts ...grpc.CallOption) (*SpaceListReply, error) {
	out := new(SpaceListReply)
	err := c.cc.Invoke(ctx, "/demo.space.SpaceService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpaceServiceServer is the server API for SpaceService service.
// All implementations must embed UnimplementedSpaceServiceServer
// for forward compatibility
type SpaceServiceServer interface {
	Create(context.Context, *SpaceCreateRequest) (*SpaceCreateReply, error)
	Remove(context.Context, *SpaceRemoveRequest) (*SpaceRemoveReply, error)
	Info(context.Context, *SpaceInfoRequest) (*SpaceInfoReply, error)
	List(context.Context, *SpaceListRequest) (*SpaceListReply, error)
	mustEmbedUnimplementedSpaceServiceServer()
}

// UnimplementedSpaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSpaceServiceServer struct {
}

func (UnimplementedSpaceServiceServer) Create(context.Context, *SpaceCreateRequest) (*SpaceCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSpaceServiceServer) Remove(context.Context, *SpaceRemoveRequest) (*SpaceRemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedSpaceServiceServer) Info(context.Context, *SpaceInfoRequest) (*SpaceInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedSpaceServiceServer) List(context.Context, *SpaceListRequest) (*SpaceListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSpaceServiceServer) mustEmbedUnimplementedSpaceServiceServer() {}

// UnsafeSpaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpaceServiceServer will
// result in compilation errors.
type UnsafeSpaceServiceServer interface {
	mustEmbedUnimplementedSpaceServiceServer()
}

func RegisterSpaceServiceServer(s grpc.ServiceRegistrar, srv SpaceServiceServer) {
	s.RegisterService(&SpaceService_ServiceDesc, srv)
}

func _SpaceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpaceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.space.SpaceService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).Create(ctx, req.(*SpaceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpaceRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.space.SpaceService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).Remove(ctx, req.(*SpaceRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpaceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.space.SpaceService/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).Info(ctx, req.(*SpaceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpaceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.space.SpaceService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).List(ctx, req.(*SpaceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpaceService_ServiceDesc is the grpc.ServiceDesc for SpaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.space.SpaceService",
	HandlerType: (*SpaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SpaceService_Create_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _SpaceService_Remove_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _SpaceService_Info_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SpaceService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/space/idl/space.proto",
}
