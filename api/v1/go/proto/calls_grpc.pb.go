// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: calls.proto

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

const (
	MainApi_SignIn_FullMethodName = "/go_boiler.calls.MainApi/SignIn"
	MainApi_SignUp_FullMethodName = "/go_boiler.calls.MainApi/SignUp"
)

// MainApiClient is the client API for MainApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MainApiClient interface {
	SignIn(ctx context.Context, in *SignInCallRequest, opts ...grpc.CallOption) (*SignInCallResponse, error)
	SignUp(ctx context.Context, in *SignUpCallRequest, opts ...grpc.CallOption) (*SignUpCallResponse, error)
}

type mainApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMainApiClient(cc grpc.ClientConnInterface) MainApiClient {
	return &mainApiClient{cc}
}

func (c *mainApiClient) SignIn(ctx context.Context, in *SignInCallRequest, opts ...grpc.CallOption) (*SignInCallResponse, error) {
	out := new(SignInCallResponse)
	err := c.cc.Invoke(ctx, MainApi_SignIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainApiClient) SignUp(ctx context.Context, in *SignUpCallRequest, opts ...grpc.CallOption) (*SignUpCallResponse, error) {
	out := new(SignUpCallResponse)
	err := c.cc.Invoke(ctx, MainApi_SignUp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainApiServer is the server API for MainApi service.
// All implementations must embed UnimplementedMainApiServer
// for forward compatibility
type MainApiServer interface {
	SignIn(context.Context, *SignInCallRequest) (*SignInCallResponse, error)
	SignUp(context.Context, *SignUpCallRequest) (*SignUpCallResponse, error)
	mustEmbedUnimplementedMainApiServer()
}

// UnimplementedMainApiServer must be embedded to have forward compatible implementations.
type UnimplementedMainApiServer struct {
}

func (UnimplementedMainApiServer) SignIn(context.Context, *SignInCallRequest) (*SignInCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedMainApiServer) SignUp(context.Context, *SignUpCallRequest) (*SignUpCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedMainApiServer) mustEmbedUnimplementedMainApiServer() {}

// UnsafeMainApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MainApiServer will
// result in compilation errors.
type UnsafeMainApiServer interface {
	mustEmbedUnimplementedMainApiServer()
}

func RegisterMainApiServer(s grpc.ServiceRegistrar, srv MainApiServer) {
	s.RegisterService(&MainApi_ServiceDesc, srv)
}

func _MainApi_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainApiServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainApi_SignIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainApiServer).SignIn(ctx, req.(*SignInCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainApi_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainApiServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainApi_SignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainApiServer).SignUp(ctx, req.(*SignUpCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MainApi_ServiceDesc is the grpc.ServiceDesc for MainApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MainApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_boiler.calls.MainApi",
	HandlerType: (*MainApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _MainApi_SignIn_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _MainApi_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calls.proto",
}
