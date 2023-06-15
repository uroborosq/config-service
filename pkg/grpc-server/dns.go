package grpc_server

import (
	"context"

	"github.com/uroborosq/config-service/pkg/grpcGen"
	"google.golang.org/protobuf/types/known/emptypb"
)

type dnsConfigService interface {
	GetServerList() ([]string, error)
	AddServer(string) error
	RemoveServer(string) error
}

type DnsGrpcServer struct {
	grpcGen.UnimplementedDnsServiceServer
	dnsService dnsConfigService
}

func NewDnsGrpcServer(service dnsConfigService) *DnsGrpcServer {
	return &DnsGrpcServer{dnsService: service}
}

func (s *DnsGrpcServer) GetServerList(_ context.Context, _ *emptypb.Empty) (*grpcGen.DnsServerList, error) {
	servers, err := s.dnsService.GetServerList()
	return &grpcGen.DnsServerList{Servers: servers}, err
}

func (s *DnsGrpcServer) AddServer(_ context.Context, server *grpcGen.DnsServer) (*emptypb.Empty, error) {
	err := s.dnsService.AddServer(server.Address)
	return &emptypb.Empty{}, err
}

func (s *DnsGrpcServer) RemoveServer(_ context.Context, server *grpcGen.DnsServer) (*emptypb.Empty, error) {
	err := s.dnsService.RemoveServer(server.Address)
	return &emptypb.Empty{}, err
}
