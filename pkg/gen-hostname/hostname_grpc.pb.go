// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: grpcGen/hostname.proto

package gen_hostname

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
	HostnameService_GetHostname_FullMethodName    = "/genHostname.HostnameService/GetHostname"
	HostnameService_UpdateHostname_FullMethodName = "/genHostname.HostnameService/UpdateHostname"
)

// HostnameServiceClient is the client API for HostnameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostnameServiceClient interface {
	GetHostname(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Hostname, error)
	UpdateHostname(ctx context.Context, in *Hostname, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type hostnameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostnameServiceClient(cc grpc.ClientConnInterface) HostnameServiceClient {
	return &hostnameServiceClient{cc}
}

func (c *hostnameServiceClient) GetHostname(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Hostname, error) {
	out := new(Hostname)
	err := c.cc.Invoke(ctx, HostnameService_GetHostname_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostnameServiceClient) UpdateHostname(ctx context.Context, in *Hostname, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, HostnameService_UpdateHostname_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostnameServiceServer is the server API for HostnameService service.
// All implementations must embed UnimplementedHostnameServiceServer
// for forward compatibility
type HostnameServiceServer interface {
	GetHostname(context.Context, *emptypb.Empty) (*Hostname, error)
	UpdateHostname(context.Context, *Hostname) (*emptypb.Empty, error)
	mustEmbedUnimplementedHostnameServiceServer()
}

// UnimplementedHostnameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHostnameServiceServer struct {
}

func (UnimplementedHostnameServiceServer) GetHostname(context.Context, *emptypb.Empty) (*Hostname, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostname not implemented")
}
func (UnimplementedHostnameServiceServer) UpdateHostname(context.Context, *Hostname) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHostname not implemented")
}
func (UnimplementedHostnameServiceServer) mustEmbedUnimplementedHostnameServiceServer() {}

// UnsafeHostnameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostnameServiceServer will
// result in compilation errors.
type UnsafeHostnameServiceServer interface {
	mustEmbedUnimplementedHostnameServiceServer()
}

func RegisterHostnameServiceServer(s grpc.ServiceRegistrar, srv HostnameServiceServer) {
	s.RegisterService(&HostnameService_ServiceDesc, srv)
}

func _HostnameService_GetHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostnameServiceServer).GetHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostnameService_GetHostname_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostnameServiceServer).GetHostname(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostnameService_UpdateHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hostname)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostnameServiceServer).UpdateHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostnameService_UpdateHostname_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostnameServiceServer).UpdateHostname(ctx, req.(*Hostname))
	}
	return interceptor(ctx, in, info, handler)
}

// HostnameService_ServiceDesc is the grpc.ServiceDesc for HostnameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HostnameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genHostname.HostnameService",
	HandlerType: (*HostnameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHostname",
			Handler:    _HostnameService_GetHostname_Handler,
		},
		{
			MethodName: "UpdateHostname",
			Handler:    _HostnameService_UpdateHostname_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcGen/hostname.proto",
}
