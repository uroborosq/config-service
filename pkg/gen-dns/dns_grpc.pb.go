// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: grpcGen/dns.proto

package gen_dns

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DnsService_GetServerList_FullMethodName = "/genDns.DnsService/GetServerList"
	DnsService_AddServer_FullMethodName     = "/genDns.DnsService/AddServer"
	DnsService_RemoveServer_FullMethodName  = "/genDns.DnsService/RemoveServer"
)

// DnsServiceClient is the client API for DnsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DnsServiceClient interface {
	GetServerList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DnsServerList, error)
	AddServer(ctx context.Context, in *DnsServer, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveServer(ctx context.Context, in *DnsServer, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dnsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsServiceClient(cc grpc.ClientConnInterface) DnsServiceClient {
	return &dnsServiceClient{cc}
}

func (c *dnsServiceClient) GetServerList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DnsServerList, error) {
	out := new(DnsServerList)
	err := c.cc.Invoke(ctx, DnsService_GetServerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsServiceClient) AddServer(ctx context.Context, in *DnsServer, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DnsService_AddServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsServiceClient) RemoveServer(ctx context.Context, in *DnsServer, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DnsService_RemoveServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnsServiceServer is the server API for DnsService service.
// All implementations must embed UnimplementedDnsServiceServer
// for forward compatibility
type DnsServiceServer interface {
	GetServerList(context.Context, *emptypb.Empty) (*DnsServerList, error)
	AddServer(context.Context, *DnsServer) (*emptypb.Empty, error)
	RemoveServer(context.Context, *DnsServer) (*emptypb.Empty, error)
	mustEmbedUnimplementedDnsServiceServer()
}

// UnimplementedDnsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDnsServiceServer struct {
}

func (UnimplementedDnsServiceServer) GetServerList(context.Context, *emptypb.Empty) (*DnsServerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerList not implemented")
}
func (UnimplementedDnsServiceServer) AddServer(context.Context, *DnsServer) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddServer not implemented")
}
func (UnimplementedDnsServiceServer) RemoveServer(context.Context, *DnsServer) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveServer not implemented")
}
func (UnimplementedDnsServiceServer) mustEmbedUnimplementedDnsServiceServer() {}

// UnsafeDnsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DnsServiceServer will
// result in compilation errors.
type UnsafeDnsServiceServer interface {
	mustEmbedUnimplementedDnsServiceServer()
}

func RegisterDnsServiceServer(s grpc.ServiceRegistrar, srv DnsServiceServer) {
	s.RegisterService(&DnsService_ServiceDesc, srv)
}

func _DnsService_GetServerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsServiceServer).GetServerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsService_GetServerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsServiceServer).GetServerList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsService_AddServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsServiceServer).AddServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsService_AddServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsServiceServer).AddServer(ctx, req.(*DnsServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsService_RemoveServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsServiceServer).RemoveServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsService_RemoveServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsServiceServer).RemoveServer(ctx, req.(*DnsServer))
	}
	return interceptor(ctx, in, info, handler)
}

// DnsService_ServiceDesc is the grpc.ServiceDesc for DnsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DnsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genDns.DnsService",
	HandlerType: (*DnsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerList",
			Handler:    _DnsService_GetServerList_Handler,
		},
		{
			MethodName: "AddServer",
			Handler:    _DnsService_AddServer_Handler,
		},
		{
			MethodName: "RemoveServer",
			Handler:    _DnsService_RemoveServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcGen/dns.proto",
}