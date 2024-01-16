// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: auth/v1/auth.proto

package v1

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
	Auth_Login_FullMethodName           = "/api.auth.v1.Auth/Login"
	Auth_LoginForApp_FullMethodName     = "/api.auth.v1.Auth/LoginForApp"
	Auth_LoginTest_FullMethodName       = "/api.auth.v1.Auth/LoginTest"
	Auth_Decrypt_FullMethodName         = "/api.auth.v1.Auth/Decrypt"
	Auth_GetInfo_FullMethodName         = "/api.auth.v1.Auth/GetInfo"
	Auth_SendCode_FullMethodName        = "/api.auth.v1.Auth/SendCode"
	Auth_VerifyLoginCode_FullMethodName = "/api.auth.v1.Auth/VerifyLoginCode"
	Auth_VerifyBindCode_FullMethodName  = "/api.auth.v1.Auth/VerifyBindCode"
	Auth_LoginByApple_FullMethodName    = "/api.auth.v1.Auth/LoginByApple"
	Auth_LoginByPhone_FullMethodName    = "/api.auth.v1.Auth/LoginByPhone"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	// 登录
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// app 登录
	LoginForApp(ctx context.Context, in *LoginForAppRequest, opts ...grpc.CallOption) (*LoginForAppReply, error)
	// 测试登录
	LoginTest(ctx context.Context, in *LoginTestRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// 解密
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptReply, error)
	// 获取登录信息
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// 注册短信
	SendCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*SendCodeReply, error)
	// 验证登录短信
	VerifyLoginCode(ctx context.Context, in *VerifyLoginCodeRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// 验证绑定短信
	VerifyBindCode(ctx context.Context, in *VerifyBindCodeRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// 苹果登录
	LoginByApple(ctx context.Context, in *LoginByAppleRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// 微信手机号登录
	LoginByPhone(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginByPhoneReply, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, Auth_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LoginForApp(ctx context.Context, in *LoginForAppRequest, opts ...grpc.CallOption) (*LoginForAppReply, error) {
	out := new(LoginForAppReply)
	err := c.cc.Invoke(ctx, Auth_LoginForApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LoginTest(ctx context.Context, in *LoginTestRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, Auth_LoginTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptReply, error) {
	out := new(DecryptReply)
	err := c.cc.Invoke(ctx, Auth_Decrypt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, Auth_GetInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SendCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*SendCodeReply, error) {
	out := new(SendCodeReply)
	err := c.cc.Invoke(ctx, Auth_SendCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) VerifyLoginCode(ctx context.Context, in *VerifyLoginCodeRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, Auth_VerifyLoginCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) VerifyBindCode(ctx context.Context, in *VerifyBindCodeRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, Auth_VerifyBindCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LoginByApple(ctx context.Context, in *LoginByAppleRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, Auth_LoginByApple_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LoginByPhone(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginByPhoneReply, error) {
	out := new(LoginByPhoneReply)
	err := c.cc.Invoke(ctx, Auth_LoginByPhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	// 登录
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// app 登录
	LoginForApp(context.Context, *LoginForAppRequest) (*LoginForAppReply, error)
	// 测试登录
	LoginTest(context.Context, *LoginTestRequest) (*LoginReply, error)
	// 解密
	Decrypt(context.Context, *DecryptRequest) (*DecryptReply, error)
	// 获取登录信息
	GetInfo(context.Context, *GetInfoRequest) (*LoginReply, error)
	// 注册短信
	SendCode(context.Context, *SendCodeRequest) (*SendCodeReply, error)
	// 验证登录短信
	VerifyLoginCode(context.Context, *VerifyLoginCodeRequest) (*LoginReply, error)
	// 验证绑定短信
	VerifyBindCode(context.Context, *VerifyBindCodeRequest) (*LoginReply, error)
	// 苹果登录
	LoginByApple(context.Context, *LoginByAppleRequest) (*LoginReply, error)
	// 微信手机号登录
	LoginByPhone(context.Context, *LoginRequest) (*LoginByPhoneReply, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) LoginForApp(context.Context, *LoginForAppRequest) (*LoginForAppReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginForApp not implemented")
}
func (UnimplementedAuthServer) LoginTest(context.Context, *LoginTestRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginTest not implemented")
}
func (UnimplementedAuthServer) Decrypt(context.Context, *DecryptRequest) (*DecryptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (UnimplementedAuthServer) GetInfo(context.Context, *GetInfoRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedAuthServer) SendCode(context.Context, *SendCodeRequest) (*SendCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCode not implemented")
}
func (UnimplementedAuthServer) VerifyLoginCode(context.Context, *VerifyLoginCodeRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyLoginCode not implemented")
}
func (UnimplementedAuthServer) VerifyBindCode(context.Context, *VerifyBindCodeRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyBindCode not implemented")
}
func (UnimplementedAuthServer) LoginByApple(context.Context, *LoginByAppleRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByApple not implemented")
}
func (UnimplementedAuthServer) LoginByPhone(context.Context, *LoginRequest) (*LoginByPhoneReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByPhone not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LoginForApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginForAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LoginForApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_LoginForApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LoginForApp(ctx, req.(*LoginForAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LoginTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LoginTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_LoginTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LoginTest(ctx, req.(*LoginTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Decrypt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SendCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SendCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_SendCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SendCode(ctx, req.(*SendCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_VerifyLoginCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyLoginCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).VerifyLoginCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_VerifyLoginCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).VerifyLoginCode(ctx, req.(*VerifyLoginCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_VerifyBindCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyBindCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).VerifyBindCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_VerifyBindCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).VerifyBindCode(ctx, req.(*VerifyBindCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LoginByApple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByAppleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LoginByApple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_LoginByApple_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LoginByApple(ctx, req.(*LoginByAppleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LoginByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LoginByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_LoginByPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LoginByPhone(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.auth.v1.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "LoginForApp",
			Handler:    _Auth_LoginForApp_Handler,
		},
		{
			MethodName: "LoginTest",
			Handler:    _Auth_LoginTest_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _Auth_Decrypt_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _Auth_GetInfo_Handler,
		},
		{
			MethodName: "SendCode",
			Handler:    _Auth_SendCode_Handler,
		},
		{
			MethodName: "VerifyLoginCode",
			Handler:    _Auth_VerifyLoginCode_Handler,
		},
		{
			MethodName: "VerifyBindCode",
			Handler:    _Auth_VerifyBindCode_Handler,
		},
		{
			MethodName: "LoginByApple",
			Handler:    _Auth_LoginByApple_Handler,
		},
		{
			MethodName: "LoginByPhone",
			Handler:    _Auth_LoginByPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/auth.proto",
}
