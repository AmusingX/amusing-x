// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ganymedeservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GanymedeServiceClient is the client API for GanymedeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GanymedeServiceClient interface {
	Pong(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*PongResponse, error)
	Regexps(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*RegexpResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	CountryCodes(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*CountryCodeList, error)
	GetVerificationCode(ctx context.Context, in *VerificationCodeRequest, opts ...grpc.CallOption) (*VerificationCodeResponse, error)
	VerificationCodeCheck(ctx context.Context, in *VerificationCodeCheckRequest, opts ...grpc.CallOption) (*VerificationCodeCheckResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	OAuthLogin(ctx context.Context, in *OAuthLoginRequest, opts ...grpc.CallOption) (*OAuthLoginResponse, error)
	OAuthInfo(ctx context.Context, in *OAuthInfoRequest, opts ...grpc.CallOption) (*OAuthInfoResponse, error)
	IsLogin(ctx context.Context, in *IsLoginRequest, opts ...grpc.CallOption) (*IsLoginResponse, error)
	LogOut(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
}

type ganymedeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGanymedeServiceClient(cc grpc.ClientConnInterface) GanymedeServiceClient {
	return &ganymedeServiceClient{cc}
}

func (c *ganymedeServiceClient) Pong(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/ganymde.GanymedeService/Pong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganymedeServiceClient) Regexps(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*RegexpResponse, error) {
	out := new(RegexpResponse)
	err := c.cc.Invoke(ctx, "/ganymde.GanymedeService/Regexps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganymedeServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/ganymde.GanymedeService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganymedeServiceClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/ganymde.GanymedeService/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganymedeServiceClient) CountryCodes(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*CountryCodeList, error) {
	out := new(CountryCodeList)
	err := c.cc.Invoke(ctx, "/ganymde.GanymedeService/CountryCodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganymedeServiceClient) GetVerificationCode(ctx context.Context, in *VerificationCodeRequest, opts ...grpc.CallOption) (*VerificationCodeResponse, error) {
	out := new(VerificationCodeResponse)
	err := c.cc.Invoke(ctx, "/ganymde.GanymedeService/GetVerificationCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganymedeServiceClient) VerificationCodeCheck(ctx context.Context, in *VerificationCodeCheckRequest, opts ...grpc.CallOption) (*VerificationCodeCheckResponse, error) {
	out := new(VerificationCodeCheckResponse)
	err := c.cc.Invoke(ctx, "/ganymde.GanymedeService/VerificationCodeCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganymedeServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/ganymde.GanymedeService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganymedeServiceClient) OAuthLogin(ctx context.Context, in *OAuthLoginRequest, opts ...grpc.CallOption) (*OAuthLoginResponse, error) {
	out := new(OAuthLoginResponse)
	err := c.cc.Invoke(ctx, "/ganymde.GanymedeService/OAuthLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganymedeServiceClient) OAuthInfo(ctx context.Context, in *OAuthInfoRequest, opts ...grpc.CallOption) (*OAuthInfoResponse, error) {
	out := new(OAuthInfoResponse)
	err := c.cc.Invoke(ctx, "/ganymde.GanymedeService/OAuthInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganymedeServiceClient) IsLogin(ctx context.Context, in *IsLoginRequest, opts ...grpc.CallOption) (*IsLoginResponse, error) {
	out := new(IsLoginResponse)
	err := c.cc.Invoke(ctx, "/ganymde.GanymedeService/IsLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganymedeServiceClient) LogOut(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/ganymde.GanymedeService/LogOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GanymedeServiceServer is the server API for GanymedeService service.
// All implementations must embed UnimplementedGanymedeServiceServer
// for forward compatibility
type GanymedeServiceServer interface {
	Pong(context.Context, *BlankParams) (*PongResponse, error)
	Regexps(context.Context, *BlankParams) (*RegexpResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	CountryCodes(context.Context, *BlankParams) (*CountryCodeList, error)
	GetVerificationCode(context.Context, *VerificationCodeRequest) (*VerificationCodeResponse, error)
	VerificationCodeCheck(context.Context, *VerificationCodeCheckRequest) (*VerificationCodeCheckResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	OAuthLogin(context.Context, *OAuthLoginRequest) (*OAuthLoginResponse, error)
	OAuthInfo(context.Context, *OAuthInfoRequest) (*OAuthInfoResponse, error)
	IsLogin(context.Context, *IsLoginRequest) (*IsLoginResponse, error)
	LogOut(context.Context, *LogoutRequest) (*LogoutResponse, error)
	mustEmbedUnimplementedGanymedeServiceServer()
}

// UnimplementedGanymedeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGanymedeServiceServer struct {
}

func (*UnimplementedGanymedeServiceServer) Pong(context.Context, *BlankParams) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pong not implemented")
}
func (*UnimplementedGanymedeServiceServer) Regexps(context.Context, *BlankParams) (*RegexpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Regexps not implemented")
}
func (*UnimplementedGanymedeServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedGanymedeServiceServer) Join(context.Context, *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (*UnimplementedGanymedeServiceServer) CountryCodes(context.Context, *BlankParams) (*CountryCodeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountryCodes not implemented")
}
func (*UnimplementedGanymedeServiceServer) GetVerificationCode(context.Context, *VerificationCodeRequest) (*VerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVerificationCode not implemented")
}
func (*UnimplementedGanymedeServiceServer) VerificationCodeCheck(context.Context, *VerificationCodeCheckRequest) (*VerificationCodeCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerificationCodeCheck not implemented")
}
func (*UnimplementedGanymedeServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (*UnimplementedGanymedeServiceServer) OAuthLogin(context.Context, *OAuthLoginRequest) (*OAuthLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OAuthLogin not implemented")
}
func (*UnimplementedGanymedeServiceServer) OAuthInfo(context.Context, *OAuthInfoRequest) (*OAuthInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OAuthInfo not implemented")
}
func (*UnimplementedGanymedeServiceServer) IsLogin(context.Context, *IsLoginRequest) (*IsLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLogin not implemented")
}
func (*UnimplementedGanymedeServiceServer) LogOut(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOut not implemented")
}
func (*UnimplementedGanymedeServiceServer) mustEmbedUnimplementedGanymedeServiceServer() {}

func RegisterGanymedeServiceServer(s *grpc.Server, srv GanymedeServiceServer) {
	s.RegisterService(&_GanymedeService_serviceDesc, srv)
}

func _GanymedeService_Pong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlankParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GanymedeServiceServer).Pong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymde.GanymedeService/Pong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GanymedeServiceServer).Pong(ctx, req.(*BlankParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GanymedeService_Regexps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlankParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GanymedeServiceServer).Regexps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymde.GanymedeService/Regexps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GanymedeServiceServer).Regexps(ctx, req.(*BlankParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GanymedeService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GanymedeServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymde.GanymedeService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GanymedeServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GanymedeService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GanymedeServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymde.GanymedeService/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GanymedeServiceServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GanymedeService_CountryCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlankParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GanymedeServiceServer).CountryCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymde.GanymedeService/CountryCodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GanymedeServiceServer).CountryCodes(ctx, req.(*BlankParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _GanymedeService_GetVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GanymedeServiceServer).GetVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymde.GanymedeService/GetVerificationCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GanymedeServiceServer).GetVerificationCode(ctx, req.(*VerificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GanymedeService_VerificationCodeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationCodeCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GanymedeServiceServer).VerificationCodeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymde.GanymedeService/VerificationCodeCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GanymedeServiceServer).VerificationCodeCheck(ctx, req.(*VerificationCodeCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GanymedeService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GanymedeServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymde.GanymedeService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GanymedeServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GanymedeService_OAuthLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GanymedeServiceServer).OAuthLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymde.GanymedeService/OAuthLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GanymedeServiceServer).OAuthLogin(ctx, req.(*OAuthLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GanymedeService_OAuthInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GanymedeServiceServer).OAuthInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymde.GanymedeService/OAuthInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GanymedeServiceServer).OAuthInfo(ctx, req.(*OAuthInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GanymedeService_IsLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GanymedeServiceServer).IsLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymde.GanymedeService/IsLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GanymedeServiceServer).IsLogin(ctx, req.(*IsLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GanymedeService_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GanymedeServiceServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymde.GanymedeService/LogOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GanymedeServiceServer).LogOut(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GanymedeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ganymde.GanymedeService",
	HandlerType: (*GanymedeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pong",
			Handler:    _GanymedeService_Pong_Handler,
		},
		{
			MethodName: "Regexps",
			Handler:    _GanymedeService_Regexps_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _GanymedeService_Login_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _GanymedeService_Join_Handler,
		},
		{
			MethodName: "CountryCodes",
			Handler:    _GanymedeService_CountryCodes_Handler,
		},
		{
			MethodName: "GetVerificationCode",
			Handler:    _GanymedeService_GetVerificationCode_Handler,
		},
		{
			MethodName: "VerificationCodeCheck",
			Handler:    _GanymedeService_VerificationCodeCheck_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _GanymedeService_ResetPassword_Handler,
		},
		{
			MethodName: "OAuthLogin",
			Handler:    _GanymedeService_OAuthLogin_Handler,
		},
		{
			MethodName: "OAuthInfo",
			Handler:    _GanymedeService_OAuthInfo_Handler,
		},
		{
			MethodName: "IsLogin",
			Handler:    _GanymedeService_IsLogin_Handler,
		},
		{
			MethodName: "LogOut",
			Handler:    _GanymedeService_LogOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
