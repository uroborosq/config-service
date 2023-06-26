package grpc_server

import (
	"context"

	genHostname "github.com/uroborosq/config-service/pkg/gen-hostname"
	"google.golang.org/protobuf/types/known/emptypb"
)

type hostnameService interface {
	UpdateHostname(string) error
	GetHostname() (string, error)
}

type HostnameGrpcServer struct {
	genHostname.UnimplementedHostnameServiceServer
	hostnameService hostnameService
}

func NewHostnameGrpcServer(service hostnameService) *HostnameGrpcServer {
	return &HostnameGrpcServer{
		hostnameService: service,
	}
}

func (s *HostnameGrpcServer) UpdateHostname(_ context.Context, hostname *genHostname.Hostname) (*emptypb.Empty, error) {
	err := s.hostnameService.UpdateHostname(hostname.Hostname)
	return &emptypb.Empty{}, err
}

func (s *HostnameGrpcServer) GetHostname(_ context.Context, _ *emptypb.Empty) (*genHostname.Hostname, error) {
	hostname, err := s.hostnameService.GetHostname()
	return &genHostname.Hostname{Hostname: hostname}, err
}
