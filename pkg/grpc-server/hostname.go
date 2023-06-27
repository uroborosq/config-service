package grpc_server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"unicode"

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
	if !isValidHostname(hostname.Hostname) {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "hostname doesn't follow the rules")
	}
	err := s.hostnameService.UpdateHostname(hostname.Hostname)
	return &emptypb.Empty{}, err
}

func (s *HostnameGrpcServer) GetHostname(_ context.Context, _ *emptypb.Empty) (*genHostname.Hostname, error) {
	hostname, err := s.hostnameService.GetHostname()
	return &genHostname.Hostname{Hostname: hostname}, err
}

// Quote from docs:
// Each element of the hostname must be from 1 to 63 characters long
// and the entire hostname, including the dots, can be at most 253
// characters long.  Valid characters for hostnames are ASCII(7)
// letters from a to z, the digits from 0 to 9, and the hyphen (-).
// A hostname may not start with a hyphen.

func isValidHostname(s string) bool {
	if len(s) > 253 {
		return false
	}

	subdomens := strings.Split(s, ".")
	if len(subdomens) > 4 {
		return false
	}
	for i := range subdomens {
		if len(subdomens[i]) < 64 && len(subdomens) > 0 {
			for _, r := range subdomens[i] {
				if !(unicode.IsDigit(r)) && !(r == '-') && ((r < 'a' || r > 'z') && (r < 'A' || r > 'Z')) {
					return false
				}
			}
		} else {
			return false
		}
	}
	return !strings.HasPrefix(s, "-")
}
