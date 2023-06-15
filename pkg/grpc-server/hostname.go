package grpc_server

import (
	"context"

	"github.com/uroborosq/config-service/pkg/grpcGen"
	"google.golang.org/protobuf/types/known/emptypb"
)

type hostnameService interface {
	UpdateHostname(string) error
	GetHostname() (string, error)
}

type HostnameGrpcServer struct {
	grpcGen.UnimplementedHostnameServiceServer
	hostnameService hostnameService
}

func NewHostnameGrpcServer(service hostnameService) *HostnameGrpcServer {
	return &HostnameGrpcServer{
		hostnameService: service,
	}
}

func (s *HostnameGrpcServer) UpdateHostname(_ context.Context, hostname *grpcGen.Hostname) (*emptypb.Empty, error) {
	err := s.hostnameService.UpdateHostname(hostname.Hostname)
	return &emptypb.Empty{}, err
}

func (s *HostnameGrpcServer) GetHostname(_ context.Context, _ *emptypb.Empty) (*grpcGen.Hostname, error) {
	hostname, err := s.hostnameService.GetHostname()
	return &grpcGen.Hostname{Hostname: hostname}, err
}
