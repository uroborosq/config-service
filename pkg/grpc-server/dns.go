package grpc_server

import (
	"context"

	genDns "github.com/uroborosq/config-service/pkg/gen-dns"
	"google.golang.org/protobuf/types/known/emptypb"
)

type dnsConfigService interface {
	GetServerList() ([]string, error)
	AddServer(string) error
	RemoveServer(string) error
}

type DnsGrpcServer struct {
	genDns.UnimplementedDnsServiceServer
	dnsService dnsConfigService
}

func NewDnsGrpcServer(service dnsConfigService) *DnsGrpcServer {
	return &DnsGrpcServer{dnsService: service}
}

func (s *DnsGrpcServer) GetServerList(_ context.Context, _ *emptypb.Empty) (*genDns.DnsServerList, error) {
	servers, err := s.dnsService.GetServerList()
	return &genDns.DnsServerList{Servers: servers}, err
}

func (s *DnsGrpcServer) AddServer(_ context.Context, server *genDns.DnsServer) (*emptypb.Empty, error) {
	err := s.dnsService.AddServer(server.Address)
	return &emptypb.Empty{}, err
}

func (s *DnsGrpcServer) RemoveServer(_ context.Context, server *genDns.DnsServer) (*emptypb.Empty, error) {
	err := s.dnsService.RemoveServer(server.Address)
	return &emptypb.Empty{}, err
}
