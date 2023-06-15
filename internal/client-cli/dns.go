package client_cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

type dnsClientInterface interface {
	GetServerList(string) ([]string, error)
	AddServer(string, string) error
	RemoveServer(string, string) error
}

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "manage dns servers on remote server-cli",
	Args:  cobra.NoArgs,
}

var getDnsCmd = &cobra.Command{
	Use:   "list",
	Short: "show list of dns servers on remote server-cli",
	Args:  cobra.NoArgs,
}

var addDnsCmd = &cobra.Command{
	Use:   "add",
	Short: "add server-cli of dns servers on remote server-cli",
	Args:  cobra.ExactArgs(1),
}

var removeDnsCommand = &cobra.Command{
	Use:   "remove",
	Short: "remove dns server-cli from remote server-cli",
	Args:  cobra.ExactArgs(1),
}

func initDnsFuncs(dnsClient dnsClientInterface) {
	dnsCmd.Run = func(cmd *cobra.Command, args []string) {
		servers, err := dnsClient.GetServerList(destination)
		if err != nil {
			log.Printf("ERR: %s\n", err.Error())
		} else if len(servers) == 0 {
			fmt.Println("No servers set")
		} else {
			fmt.Printf("Configured servers on %s\n", destination)
			for _, server := range servers {
				fmt.Println(server)
			}
		}
	}
	getDnsCmd.Run = dnsCmd.Run
	addDnsCmd.Run = func(cmd *cobra.Command, args []string) {
		err := dnsClient.AddServer(destination, args[0])
		if err != nil {
			log.Printf("ERR: %s\n", err.Error())
		} else {
			fmt.Printf("Server %s added successfully\n", args[0])
		}
	}
	removeDnsCommand.Run = func(cmd *cobra.Command, args []string) {
		err := dnsClient.RemoveServer(destination, args[0])
		if err != nil {
			log.Printf("ERR: %s\n", err.Error())
		} else {
			fmt.Printf("Server %s removed successfully\n", args[0])
		}
	}
}
