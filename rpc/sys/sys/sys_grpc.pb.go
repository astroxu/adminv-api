// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: rpc/sys/sys.proto

package sys

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
	Sys_Login_FullMethodName          = "/sys.Sys/Login"
	Sys_LoginLogAdd_FullMethodName    = "/sys.Sys/LoginLogAdd"
	Sys_LoginLogList_FullMethodName   = "/sys.Sys/LoginLogList"
	Sys_LoginLogDelete_FullMethodName = "/sys.Sys/LoginLogDelete"
)

// SysClient is the client API for Sys service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysClient interface {
	// 登录
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	// 登录日志
	LoginLogAdd(ctx context.Context, in *LoginLogAddReq, opts ...grpc.CallOption) (*LoginLogAddResp, error)
	LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error)
	LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error)
}

type sysClient struct {
	cc grpc.ClientConnInterface
}

func NewSysClient(cc grpc.ClientConnInterface) SysClient {
	return &sysClient{cc}
}

func (c *sysClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, Sys_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) LoginLogAdd(ctx context.Context, in *LoginLogAddReq, opts ...grpc.CallOption) (*LoginLogAddResp, error) {
	out := new(LoginLogAddResp)
	err := c.cc.Invoke(ctx, Sys_LoginLogAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error) {
	out := new(LoginLogListResp)
	err := c.cc.Invoke(ctx, Sys_LoginLogList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error) {
	out := new(LoginLogDeleteResp)
	err := c.cc.Invoke(ctx, Sys_LoginLogDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysServer is the server API for Sys service.
// All implementations must embed UnimplementedSysServer
// for forward compatibility
type SysServer interface {
	// 登录
	Login(context.Context, *LoginReq) (*LoginResp, error)
	// 登录日志
	LoginLogAdd(context.Context, *LoginLogAddReq) (*LoginLogAddResp, error)
	LoginLogList(context.Context, *LoginLogListReq) (*LoginLogListResp, error)
	LoginLogDelete(context.Context, *LoginLogDeleteReq) (*LoginLogDeleteResp, error)
	mustEmbedUnimplementedSysServer()
}

// UnimplementedSysServer must be embedded to have forward compatible implementations.
type UnimplementedSysServer struct {
}

func (UnimplementedSysServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSysServer) LoginLogAdd(context.Context, *LoginLogAddReq) (*LoginLogAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginLogAdd not implemented")
}
func (UnimplementedSysServer) LoginLogList(context.Context, *LoginLogListReq) (*LoginLogListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginLogList not implemented")
}
func (UnimplementedSysServer) LoginLogDelete(context.Context, *LoginLogDeleteReq) (*LoginLogDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginLogDelete not implemented")
}
func (UnimplementedSysServer) mustEmbedUnimplementedSysServer() {}

// UnsafeSysServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysServer will
// result in compilation errors.
type UnsafeSysServer interface {
	mustEmbedUnimplementedSysServer()
}

func RegisterSysServer(s grpc.ServiceRegistrar, srv SysServer) {
	s.RegisterService(&Sys_ServiceDesc, srv)
}

func _Sys_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_LoginLogAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginLogAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).LoginLogAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_LoginLogAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).LoginLogAdd(ctx, req.(*LoginLogAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_LoginLogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginLogListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).LoginLogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_LoginLogList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).LoginLogList(ctx, req.(*LoginLogListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_LoginLogDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginLogDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).LoginLogDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_LoginLogDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).LoginLogDelete(ctx, req.(*LoginLogDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Sys_ServiceDesc is the grpc.ServiceDesc for Sys service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sys_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sys.Sys",
	HandlerType: (*SysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Sys_Login_Handler,
		},
		{
			MethodName: "LoginLogAdd",
			Handler:    _Sys_LoginLogAdd_Handler,
		},
		{
			MethodName: "LoginLogList",
			Handler:    _Sys_LoginLogList_Handler,
		},
		{
			MethodName: "LoginLogDelete",
			Handler:    _Sys_LoginLogDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/sys/sys.proto",
}
