// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: idl/machine.proto

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

const (
	MachineService_Create_FullMethodName = "/mesh.MachineService/Create"
	MachineService_Remove_FullMethodName = "/mesh.MachineService/Remove"
	MachineService_Change_FullMethodName = "/mesh.MachineService/Change"
	MachineService_Info_FullMethodName   = "/mesh.MachineService/Info"
	MachineService_List_FullMethodName   = "/mesh.MachineService/List"
)

// MachineServiceClient is the client API for MachineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MachineServiceClient interface {
	Create(ctx context.Context, in *MachineCreateRequest, opts ...grpc.CallOption) (*MachineCreateReply, error)
	Remove(ctx context.Context, in *MachineRemoveRequest, opts ...grpc.CallOption) (*MachineRemoveReply, error)
	Change(ctx context.Context, in *MachineChangeRequest, opts ...grpc.CallOption) (*MachineChangeReply, error)
	Info(ctx context.Context, in *MachineInfoRequest, opts ...grpc.CallOption) (*MachineInfoReply, error)
	List(ctx context.Context, in *MachineListRequest, opts ...grpc.CallOption) (*MachineListReply, error)
}

type machineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMachineServiceClient(cc grpc.ClientConnInterface) MachineServiceClient {
	return &machineServiceClient{cc}
}

func (c *machineServiceClient) Create(ctx context.Context, in *MachineCreateRequest, opts ...grpc.CallOption) (*MachineCreateReply, error) {
	out := new(MachineCreateReply)
	err := c.cc.Invoke(ctx, MachineService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) Remove(ctx context.Context, in *MachineRemoveRequest, opts ...grpc.CallOption) (*MachineRemoveReply, error) {
	out := new(MachineRemoveReply)
	err := c.cc.Invoke(ctx, MachineService_Remove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) Change(ctx context.Context, in *MachineChangeRequest, opts ...grpc.CallOption) (*MachineChangeReply, error) {
	out := new(MachineChangeReply)
	err := c.cc.Invoke(ctx, MachineService_Change_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) Info(ctx context.Context, in *MachineInfoRequest, opts ...grpc.CallOption) (*MachineInfoReply, error) {
	out := new(MachineInfoReply)
	err := c.cc.Invoke(ctx, MachineService_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) List(ctx context.Context, in *MachineListRequest, opts ...grpc.CallOption) (*MachineListReply, error) {
	out := new(MachineListReply)
	err := c.cc.Invoke(ctx, MachineService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MachineServiceServer is the server API for MachineService service.
// All implementations must embed UnimplementedMachineServiceServer
// for forward compatibility
type MachineServiceServer interface {
	Create(context.Context, *MachineCreateRequest) (*MachineCreateReply, error)
	Remove(context.Context, *MachineRemoveRequest) (*MachineRemoveReply, error)
	Change(context.Context, *MachineChangeRequest) (*MachineChangeReply, error)
	Info(context.Context, *MachineInfoRequest) (*MachineInfoReply, error)
	List(context.Context, *MachineListRequest) (*MachineListReply, error)
	mustEmbedUnimplementedMachineServiceServer()
}

// UnimplementedMachineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMachineServiceServer struct {
}

func (UnimplementedMachineServiceServer) Create(context.Context, *MachineCreateRequest) (*MachineCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMachineServiceServer) Remove(context.Context, *MachineRemoveRequest) (*MachineRemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedMachineServiceServer) Change(context.Context, *MachineChangeRequest) (*MachineChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Change not implemented")
}
func (UnimplementedMachineServiceServer) Info(context.Context, *MachineInfoRequest) (*MachineInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedMachineServiceServer) List(context.Context, *MachineListRequest) (*MachineListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMachineServiceServer) mustEmbedUnimplementedMachineServiceServer() {}

// UnsafeMachineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MachineServiceServer will
// result in compilation errors.
type UnsafeMachineServiceServer interface {
	mustEmbedUnimplementedMachineServiceServer()
}

func RegisterMachineServiceServer(s grpc.ServiceRegistrar, srv MachineServiceServer) {
	s.RegisterService(&MachineService_ServiceDesc, srv)
}

func _MachineService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).Create(ctx, req.(*MachineCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineService_Remove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).Remove(ctx, req.(*MachineRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_Change_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).Change(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineService_Change_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).Change(ctx, req.(*MachineChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineService_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).Info(ctx, req.(*MachineInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).List(ctx, req.(*MachineListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MachineService_ServiceDesc is the grpc.ServiceDesc for MachineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MachineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mesh.MachineService",
	HandlerType: (*MachineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MachineService_Create_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _MachineService_Remove_Handler,
		},
		{
			MethodName: "Change",
			Handler:    _MachineService_Change_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _MachineService_Info_Handler,
		},
		{
			MethodName: "List",
			Handler:    _MachineService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/machine.proto",
}
