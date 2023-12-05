// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: pheromones.proto

package pheromones

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

// PheromonesClient is the client API for Pheromones service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PheromonesClient interface {
	AddRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Response, error)
	DispatchAll(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Repeatedresponse, error)
	Dispatch(ctx context.Context, in *DispatchRequest, opts ...grpc.CallOption) (*Message, error)
}

type pheromonesClient struct {
	cc grpc.ClientConnInterface
}

func NewPheromonesClient(cc grpc.ClientConnInterface) PheromonesClient {
	return &pheromonesClient{cc}
}

func (c *pheromonesClient) AddRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pheromones.Pheromones/AddRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pheromonesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pheromones.Pheromones/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pheromonesClient) DispatchAll(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Repeatedresponse, error) {
	out := new(Repeatedresponse)
	err := c.cc.Invoke(ctx, "/pheromones.Pheromones/DispatchAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pheromonesClient) Dispatch(ctx context.Context, in *DispatchRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/pheromones.Pheromones/Dispatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PheromonesServer is the server API for Pheromones service.
// All implementations must embed UnimplementedPheromonesServer
// for forward compatibility
type PheromonesServer interface {
	AddRoute(context.Context, *RouteRequest) (*Response, error)
	Delete(context.Context, *DeleteRequest) (*Response, error)
	DispatchAll(context.Context, *Message) (*Repeatedresponse, error)
	Dispatch(context.Context, *DispatchRequest) (*Message, error)
	mustEmbedUnimplementedPheromonesServer()
}

// UnimplementedPheromonesServer must be embedded to have forward compatible implementations.
type UnimplementedPheromonesServer struct {
}

func (UnimplementedPheromonesServer) AddRoute(context.Context, *RouteRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoute not implemented")
}
func (UnimplementedPheromonesServer) Delete(context.Context, *DeleteRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPheromonesServer) DispatchAll(context.Context, *Message) (*Repeatedresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DispatchAll not implemented")
}
func (UnimplementedPheromonesServer) Dispatch(context.Context, *DispatchRequest) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dispatch not implemented")
}
func (UnimplementedPheromonesServer) mustEmbedUnimplementedPheromonesServer() {}

// UnsafePheromonesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PheromonesServer will
// result in compilation errors.
type UnsafePheromonesServer interface {
	mustEmbedUnimplementedPheromonesServer()
}

func RegisterPheromonesServer(s grpc.ServiceRegistrar, srv PheromonesServer) {
	s.RegisterService(&Pheromones_ServiceDesc, srv)
}

func _Pheromones_AddRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PheromonesServer).AddRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pheromones.Pheromones/AddRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PheromonesServer).AddRoute(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pheromones_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PheromonesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pheromones.Pheromones/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PheromonesServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pheromones_DispatchAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PheromonesServer).DispatchAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pheromones.Pheromones/DispatchAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PheromonesServer).DispatchAll(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pheromones_Dispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PheromonesServer).Dispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pheromones.Pheromones/Dispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PheromonesServer).Dispatch(ctx, req.(*DispatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Pheromones_ServiceDesc is the grpc.ServiceDesc for Pheromones service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pheromones_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pheromones.Pheromones",
	HandlerType: (*PheromonesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRoute",
			Handler:    _Pheromones_AddRoute_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Pheromones_Delete_Handler,
		},
		{
			MethodName: "DispatchAll",
			Handler:    _Pheromones_DispatchAll_Handler,
		},
		{
			MethodName: "Dispatch",
			Handler:    _Pheromones_Dispatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pheromones.proto",
}
