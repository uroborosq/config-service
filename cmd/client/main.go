package main

import (
	clientCli "github.com/uroborosq/config-service/internal/client-cli"
	"github.com/uroborosq/config-service/pkg/client"
)

func main() {
	clientCli.Execute(client.NewDnsClient(), client.NewHostnameClient())
}
