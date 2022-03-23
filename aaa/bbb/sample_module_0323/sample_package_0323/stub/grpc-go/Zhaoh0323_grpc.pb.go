// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stubs

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

// Zhaoh0323Client is the client API for Zhaoh0323 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Zhaoh0323Client interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type zhaoh0323Client struct {
	cc grpc.ClientConnInterface
}

func NewZhaoh0323Client(cc grpc.ClientConnInterface) Zhaoh0323Client {
	return &zhaoh0323Client{cc}
}

func (c *zhaoh0323Client) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/sample_module_0323.sample_package_0323.Zhaoh0323/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Zhaoh0323Server is the server API for Zhaoh0323 service.
// All implementations must embed UnimplementedZhaoh0323Server
// for forward compatibility
type Zhaoh0323Server interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedZhaoh0323Server()
}

// UnimplementedZhaoh0323Server must be embedded to have forward compatible implementations.
type UnimplementedZhaoh0323Server struct {
}

func (UnimplementedZhaoh0323Server) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedZhaoh0323Server) mustEmbedUnimplementedZhaoh0323Server() {}

// UnsafeZhaoh0323Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Zhaoh0323Server will
// result in compilation errors.
type UnsafeZhaoh0323Server interface {
	mustEmbedUnimplementedZhaoh0323Server()
}

func RegisterZhaoh0323Server(s grpc.ServiceRegistrar, srv Zhaoh0323Server) {
	s.RegisterService(&Zhaoh0323_ServiceDesc, srv)
}

func _Zhaoh0323_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Zhaoh0323Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample_module_0323.sample_package_0323.Zhaoh0323/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Zhaoh0323Server).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Zhaoh0323_ServiceDesc is the grpc.ServiceDesc for Zhaoh0323 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Zhaoh0323_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sample_module_0323.sample_package_0323.Zhaoh0323",
	HandlerType: (*Zhaoh0323Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Zhaoh0323_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Zhaoh0323.proto",
}
