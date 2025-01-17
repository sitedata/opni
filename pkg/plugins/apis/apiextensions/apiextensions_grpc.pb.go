// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/pkg/plugins/apis/apiextensions/apiextensions.proto

package apiextensions

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ManagementAPIExtensionClient is the client API for ManagementAPIExtension service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagementAPIExtensionClient interface {
	Descriptors(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ServiceDescriptorProtoList, error)
}

type managementAPIExtensionClient struct {
	cc grpc.ClientConnInterface
}

func NewManagementAPIExtensionClient(cc grpc.ClientConnInterface) ManagementAPIExtensionClient {
	return &managementAPIExtensionClient{cc}
}

func (c *managementAPIExtensionClient) Descriptors(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ServiceDescriptorProtoList, error) {
	out := new(ServiceDescriptorProtoList)
	err := c.cc.Invoke(ctx, "/apiextensions.ManagementAPIExtension/Descriptors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagementAPIExtensionServer is the server API for ManagementAPIExtension service.
// All implementations must embed UnimplementedManagementAPIExtensionServer
// for forward compatibility
type ManagementAPIExtensionServer interface {
	Descriptors(context.Context, *emptypb.Empty) (*ServiceDescriptorProtoList, error)
	mustEmbedUnimplementedManagementAPIExtensionServer()
}

// UnimplementedManagementAPIExtensionServer must be embedded to have forward compatible implementations.
type UnimplementedManagementAPIExtensionServer struct {
}

func (UnimplementedManagementAPIExtensionServer) Descriptors(context.Context, *emptypb.Empty) (*ServiceDescriptorProtoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Descriptors not implemented")
}
func (UnimplementedManagementAPIExtensionServer) mustEmbedUnimplementedManagementAPIExtensionServer() {
}

// UnsafeManagementAPIExtensionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagementAPIExtensionServer will
// result in compilation errors.
type UnsafeManagementAPIExtensionServer interface {
	mustEmbedUnimplementedManagementAPIExtensionServer()
}

func RegisterManagementAPIExtensionServer(s grpc.ServiceRegistrar, srv ManagementAPIExtensionServer) {
	s.RegisterService(&ManagementAPIExtension_ServiceDesc, srv)
}

func _ManagementAPIExtension_Descriptors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementAPIExtensionServer).Descriptors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiextensions.ManagementAPIExtension/Descriptors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementAPIExtensionServer).Descriptors(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagementAPIExtension_ServiceDesc is the grpc.ServiceDesc for ManagementAPIExtension service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagementAPIExtension_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apiextensions.ManagementAPIExtension",
	HandlerType: (*ManagementAPIExtensionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Descriptors",
			Handler:    _ManagementAPIExtension_Descriptors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/plugins/apis/apiextensions/apiextensions.proto",
}

// HTTPAPIExtensionClient is the client API for HTTPAPIExtension service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HTTPAPIExtensionClient interface {
	Configure(ctx context.Context, in *CertConfig, opts ...grpc.CallOption) (*HTTPAPIExtensionConfig, error)
}

type hTTPAPIExtensionClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPAPIExtensionClient(cc grpc.ClientConnInterface) HTTPAPIExtensionClient {
	return &hTTPAPIExtensionClient{cc}
}

func (c *hTTPAPIExtensionClient) Configure(ctx context.Context, in *CertConfig, opts ...grpc.CallOption) (*HTTPAPIExtensionConfig, error) {
	out := new(HTTPAPIExtensionConfig)
	err := c.cc.Invoke(ctx, "/apiextensions.HTTPAPIExtension/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPAPIExtensionServer is the server API for HTTPAPIExtension service.
// All implementations must embed UnimplementedHTTPAPIExtensionServer
// for forward compatibility
type HTTPAPIExtensionServer interface {
	Configure(context.Context, *CertConfig) (*HTTPAPIExtensionConfig, error)
	mustEmbedUnimplementedHTTPAPIExtensionServer()
}

// UnimplementedHTTPAPIExtensionServer must be embedded to have forward compatible implementations.
type UnimplementedHTTPAPIExtensionServer struct {
}

func (UnimplementedHTTPAPIExtensionServer) Configure(context.Context, *CertConfig) (*HTTPAPIExtensionConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedHTTPAPIExtensionServer) mustEmbedUnimplementedHTTPAPIExtensionServer() {}

// UnsafeHTTPAPIExtensionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HTTPAPIExtensionServer will
// result in compilation errors.
type UnsafeHTTPAPIExtensionServer interface {
	mustEmbedUnimplementedHTTPAPIExtensionServer()
}

func RegisterHTTPAPIExtensionServer(s grpc.ServiceRegistrar, srv HTTPAPIExtensionServer) {
	s.RegisterService(&HTTPAPIExtension_ServiceDesc, srv)
}

func _HTTPAPIExtension_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPAPIExtensionServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiextensions.HTTPAPIExtension/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPAPIExtensionServer).Configure(ctx, req.(*CertConfig))
	}
	return interceptor(ctx, in, info, handler)
}

// HTTPAPIExtension_ServiceDesc is the grpc.ServiceDesc for HTTPAPIExtension service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HTTPAPIExtension_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apiextensions.HTTPAPIExtension",
	HandlerType: (*HTTPAPIExtensionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _HTTPAPIExtension_Configure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/plugins/apis/apiextensions/apiextensions.proto",
}

// StreamAPIExtensionClient is the client API for StreamAPIExtension service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamAPIExtensionClient interface {
	// rpc Services(google.protobuf.Empty) returns (ServiceDescriptorList);
	Todo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type streamAPIExtensionClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamAPIExtensionClient(cc grpc.ClientConnInterface) StreamAPIExtensionClient {
	return &streamAPIExtensionClient{cc}
}

func (c *streamAPIExtensionClient) Todo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/apiextensions.StreamAPIExtension/Todo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamAPIExtensionServer is the server API for StreamAPIExtension service.
// All implementations must embed UnimplementedStreamAPIExtensionServer
// for forward compatibility
type StreamAPIExtensionServer interface {
	// rpc Services(google.protobuf.Empty) returns (ServiceDescriptorList);
	Todo(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedStreamAPIExtensionServer()
}

// UnimplementedStreamAPIExtensionServer must be embedded to have forward compatible implementations.
type UnimplementedStreamAPIExtensionServer struct {
}

func (UnimplementedStreamAPIExtensionServer) Todo(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Todo not implemented")
}
func (UnimplementedStreamAPIExtensionServer) mustEmbedUnimplementedStreamAPIExtensionServer() {}

// UnsafeStreamAPIExtensionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamAPIExtensionServer will
// result in compilation errors.
type UnsafeStreamAPIExtensionServer interface {
	mustEmbedUnimplementedStreamAPIExtensionServer()
}

func RegisterStreamAPIExtensionServer(s grpc.ServiceRegistrar, srv StreamAPIExtensionServer) {
	s.RegisterService(&StreamAPIExtension_ServiceDesc, srv)
}

func _StreamAPIExtension_Todo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamAPIExtensionServer).Todo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiextensions.StreamAPIExtension/Todo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamAPIExtensionServer).Todo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamAPIExtension_ServiceDesc is the grpc.ServiceDesc for StreamAPIExtension service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamAPIExtension_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apiextensions.StreamAPIExtension",
	HandlerType: (*StreamAPIExtensionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Todo",
			Handler:    _StreamAPIExtension_Todo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/plugins/apis/apiextensions/apiextensions.proto",
}

// UnaryAPIExtensionClient is the client API for UnaryAPIExtension service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UnaryAPIExtensionClient interface {
	UnaryDescriptor(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*descriptorpb.ServiceDescriptorProto, error)
}

type unaryAPIExtensionClient struct {
	cc grpc.ClientConnInterface
}

func NewUnaryAPIExtensionClient(cc grpc.ClientConnInterface) UnaryAPIExtensionClient {
	return &unaryAPIExtensionClient{cc}
}

func (c *unaryAPIExtensionClient) UnaryDescriptor(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*descriptorpb.ServiceDescriptorProto, error) {
	out := new(descriptorpb.ServiceDescriptorProto)
	err := c.cc.Invoke(ctx, "/apiextensions.UnaryAPIExtension/UnaryDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnaryAPIExtensionServer is the server API for UnaryAPIExtension service.
// All implementations must embed UnimplementedUnaryAPIExtensionServer
// for forward compatibility
type UnaryAPIExtensionServer interface {
	UnaryDescriptor(context.Context, *emptypb.Empty) (*descriptorpb.ServiceDescriptorProto, error)
	mustEmbedUnimplementedUnaryAPIExtensionServer()
}

// UnimplementedUnaryAPIExtensionServer must be embedded to have forward compatible implementations.
type UnimplementedUnaryAPIExtensionServer struct {
}

func (UnimplementedUnaryAPIExtensionServer) UnaryDescriptor(context.Context, *emptypb.Empty) (*descriptorpb.ServiceDescriptorProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryDescriptor not implemented")
}
func (UnimplementedUnaryAPIExtensionServer) mustEmbedUnimplementedUnaryAPIExtensionServer() {}

// UnsafeUnaryAPIExtensionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UnaryAPIExtensionServer will
// result in compilation errors.
type UnsafeUnaryAPIExtensionServer interface {
	mustEmbedUnimplementedUnaryAPIExtensionServer()
}

func RegisterUnaryAPIExtensionServer(s grpc.ServiceRegistrar, srv UnaryAPIExtensionServer) {
	s.RegisterService(&UnaryAPIExtension_ServiceDesc, srv)
}

func _UnaryAPIExtension_UnaryDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnaryAPIExtensionServer).UnaryDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiextensions.UnaryAPIExtension/UnaryDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnaryAPIExtensionServer).UnaryDescriptor(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// UnaryAPIExtension_ServiceDesc is the grpc.ServiceDesc for UnaryAPIExtension service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UnaryAPIExtension_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apiextensions.UnaryAPIExtension",
	HandlerType: (*UnaryAPIExtensionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryDescriptor",
			Handler:    _UnaryAPIExtension_UnaryDescriptor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/plugins/apis/apiextensions/apiextensions.proto",
}
