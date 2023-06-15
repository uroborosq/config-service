package client

import (
	"context"
	"github.com/uroborosq/config-service/pkg/grpcGen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DnsClient struct {
}

func NewDnsClient() *DnsClient {
	return &DnsClient{}
}

func (c *DnsClient) GetServerList(url string) ([]string, error) {
	client, err := c.getClient(url)
	if err != nil {
		return nil, err
	}
	servers, err := client.GetServerList(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return servers.Servers, nil
}

func (c *DnsClient) AddServer(url string, dnsAddress string) error {
	client, err := c.getClient(url)
	if err != nil {
		return err
	}
	_, err = client.AddServer(context.Background(), &grpcGen.DnsServer{Address: dnsAddress})
	return err
}

func (c *DnsClient) RemoveServer(url string, dnsAddress string) error {
	client, err := c.getClient(url)
	if err != nil {
		return err
	}
	_, err = client.RemoveServer(context.Background(), &grpcGen.DnsServer{Address: dnsAddress})
	return err
}

func (c *DnsClient) getClient(url string) (grpcGen.DnsServiceClient, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpcGen.NewDnsServiceClient(conn), err
}
