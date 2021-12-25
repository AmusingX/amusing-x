// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package charonservice

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

// CharonServClient is the client API for CharonServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CharonServClient interface {
	Pong(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*PongResponse, error)
	Create(ctx context.Context, in *CategoryCreateRequest, opts ...grpc.CallOption) (*CategoryCreateResponse, error)
	Categories(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*CategoryListResponse, error)
	Delete(ctx context.Context, in *CategoryDeleteRequest, opts ...grpc.CallOption) (*CategoryDeleteResponse, error)
}

type charonServClient struct {
	cc grpc.ClientConnInterface
}

func NewCharonServClient(cc grpc.ClientConnInterface) CharonServClient {
	return &charonServClient{cc}
}

func (c *charonServClient) Pong(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Pong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) Create(ctx context.Context, in *CategoryCreateRequest, opts ...grpc.CallOption) (*CategoryCreateResponse, error) {
	out := new(CategoryCreateResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) Categories(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*CategoryListResponse, error) {
	out := new(CategoryListResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Categories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *charonServClient) Delete(ctx context.Context, in *CategoryDeleteRequest, opts ...grpc.CallOption) (*CategoryDeleteResponse, error) {
	out := new(CategoryDeleteResponse)
	err := c.cc.Invoke(ctx, "/charonservice.CharonServ/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharonServServer is the server API for CharonServ service.
// All implementations must embed UnimplementedCharonServServer
// for forward compatibility
type CharonServServer interface {
	Pong(context.Context, *BlankParams) (*PongResponse, error)
	Create(context.Context, *CategoryCreateRequest) (*CategoryCreateResponse, error)
	Categories(context.Context, *CategoryListRequest) (*CategoryListResponse, error)
	Delete(context.Context, *CategoryDeleteRequest) (*CategoryDeleteResponse, error)
	mustEmbedUnimplementedCharonServServer()
}

// UnimplementedCharonServServer must be embedded to have forward compatible implementations.
type UnimplementedCharonServServer struct {
}

func (UnimplementedCharonServServer) Pong(context.Context, *BlankParams) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pong not implemented")
}
func (UnimplementedCharonServServer) Create(context.Context, *CategoryCreateRequest) (*CategoryCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCharonServServer) Categories(context.Context, *CategoryListRequest) (*CategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Categories not implemented")
}
func (UnimplementedCharonServServer) Delete(context.Context, *CategoryDeleteRequest) (*CategoryDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCharonServServer) mustEmbedUnimplementedCharonServServer() {}

// UnsafeCharonServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CharonServServer will
// result in compilation errors.
type UnsafeCharonServServer interface {
	mustEmbedUnimplementedCharonServServer()
}

func RegisterCharonServServer(s grpc.ServiceRegistrar, srv CharonServServer) {
	s.RegisterService(&CharonServ_ServiceDesc, srv)
}

func _CharonServ_Pong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlankParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Pong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Pong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Pong(ctx, req.(*BlankParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Create(ctx, req.(*CategoryCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_Categories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Categories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Categories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Categories(ctx, req.(*CategoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharonServ_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharonServServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonservice.CharonServ/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharonServServer).Delete(ctx, req.(*CategoryDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CharonServ_ServiceDesc is the grpc.ServiceDesc for CharonServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CharonServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "charonservice.CharonServ",
	HandlerType: (*CharonServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pong",
			Handler:    _CharonServ_Pong_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CharonServ_Create_Handler,
		},
		{
			MethodName: "Categories",
			Handler:    _CharonServ_Categories_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CharonServ_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
