package client

import (
	"context"

	genHostname "github.com/uroborosq/config-service/pkg/gen-hostname"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HostnameClient struct {
}

func NewHostnameClient() *HostnameClient {
	return &HostnameClient{}
}

func (c *HostnameClient) GetHostname(url string) (string, error) {
	client, err := c.getClient(url)
	if err != nil {
		return "", err
	}
	hostname, err := client.GetHostname(context.Background(), &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return hostname.Hostname, err
}

func (c *HostnameClient) UpdateHostname(url string, newHostname string) error {
	client, err := c.getClient(url)
	if err != nil {
		return err
	}
	_, err = client.UpdateHostname(context.Background(), &genHostname.Hostname{Hostname: newHostname})
	return err
}

func (c *HostnameClient) getClient(url string) (genHostname.HostnameServiceClient, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return genHostname.NewHostnameServiceClient(conn), err
}
