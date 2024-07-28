// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerClient interface {
	AdminCmd(ctx context.Context, in *AdminRequest, opts ...grpc.CallOption) (*AdminResult, error)
	// Implant cmds:
	ImplantCmd(ctx context.Context, in *ImplantRequest, opts ...grpc.CallOption) (*ImplantResult, error)
	ImplantAlive(ctx context.Context, in *AliveMsg, opts ...grpc.CallOption) (*Empty, error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) AdminCmd(ctx context.Context, in *AdminRequest, opts ...grpc.CallOption) (*AdminResult, error) {
	out := new(AdminResult)
	err := c.cc.Invoke(ctx, "/Server/AdminCmd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) ImplantCmd(ctx context.Context, in *ImplantRequest, opts ...grpc.CallOption) (*ImplantResult, error) {
	out := new(ImplantResult)
	err := c.cc.Invoke(ctx, "/Server/ImplantCmd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) ImplantAlive(ctx context.Context, in *AliveMsg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Server/ImplantAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
// All implementations must embed UnimplementedServerServer
// for forward compatibility
type ServerServer interface {
	AdminCmd(context.Context, *AdminRequest) (*AdminResult, error)
	// Implant cmds:
	ImplantCmd(context.Context, *ImplantRequest) (*ImplantResult, error)
	ImplantAlive(context.Context, *AliveMsg) (*Empty, error)
	mustEmbedUnimplementedServerServer()
}

// UnimplementedServerServer must be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (UnimplementedServerServer) AdminCmd(context.Context, *AdminRequest) (*AdminResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCmd not implemented")
}
func (UnimplementedServerServer) ImplantCmd(context.Context, *ImplantRequest) (*ImplantResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImplantCmd not implemented")
}
func (UnimplementedServerServer) ImplantAlive(context.Context, *AliveMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImplantAlive not implemented")
}
func (UnimplementedServerServer) mustEmbedUnimplementedServerServer() {}

// UnsafeServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerServer will
// result in compilation errors.
type UnsafeServerServer interface {
	mustEmbedUnimplementedServerServer()
}

func RegisterServerServer(s grpc.ServiceRegistrar, srv ServerServer) {
	s.RegisterService(&Server_ServiceDesc, srv)
}

func _Server_AdminCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).AdminCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server/AdminCmd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).AdminCmd(ctx, req.(*AdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_ImplantCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImplantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).ImplantCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server/ImplantCmd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).ImplantCmd(ctx, req.(*ImplantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_ImplantAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AliveMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).ImplantAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server/ImplantAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).ImplantAlive(ctx, req.(*AliveMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// Server_ServiceDesc is the grpc.ServiceDesc for Server service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Server_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminCmd",
			Handler:    _Server_AdminCmd_Handler,
		},
		{
			MethodName: "ImplantCmd",
			Handler:    _Server_ImplantCmd_Handler,
		},
		{
			MethodName: "ImplantAlive",
			Handler:    _Server_ImplantAlive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/proto/grpc.proto",
}

// PushCommandClient is the client API for PushCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushCommandClient interface {
	PushCmd(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResult, error)
}

type pushCommandClient struct {
	cc grpc.ClientConnInterface
}

func NewPushCommandClient(cc grpc.ClientConnInterface) PushCommandClient {
	return &pushCommandClient{cc}
}

func (c *pushCommandClient) PushCmd(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResult, error) {
	out := new(PushResult)
	err := c.cc.Invoke(ctx, "/PushCommand/PushCmd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushCommandServer is the server API for PushCommand service.
// All implementations must embed UnimplementedPushCommandServer
// for forward compatibility
type PushCommandServer interface {
	PushCmd(context.Context, *PushRequest) (*PushResult, error)
	mustEmbedUnimplementedPushCommandServer()
}

// UnimplementedPushCommandServer must be embedded to have forward compatible implementations.
type UnimplementedPushCommandServer struct {
}

func (UnimplementedPushCommandServer) PushCmd(context.Context, *PushRequest) (*PushResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushCmd not implemented")
}
func (UnimplementedPushCommandServer) mustEmbedUnimplementedPushCommandServer() {}

// UnsafePushCommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushCommandServer will
// result in compilation errors.
type UnsafePushCommandServer interface {
	mustEmbedUnimplementedPushCommandServer()
}

func RegisterPushCommandServer(s grpc.ServiceRegistrar, srv PushCommandServer) {
	s.RegisterService(&PushCommand_ServiceDesc, srv)
}

func _PushCommand_PushCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushCommandServer).PushCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PushCommand/PushCmd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushCommandServer).PushCmd(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PushCommand_ServiceDesc is the grpc.ServiceDesc for PushCommand service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PushCommand_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PushCommand",
	HandlerType: (*PushCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushCmd",
			Handler:    _PushCommand_PushCmd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/proto/grpc.proto",
}
