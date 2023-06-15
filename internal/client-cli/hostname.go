package client_cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

type hostnameClientInterface interface {
	GetHostname(string) (string, error)
	UpdateHostname(string, string) error
}

var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "manage hostname on remote server-cli",
	Args:  cobra.NoArgs,
}

var setHostnameCmd = &cobra.Command{
	Use:   "update",
	Short: "update hostname on remote server-cli",
	Args:  cobra.ExactArgs(1),
}

var getHostnameCmd = &cobra.Command{
	Use:   "get",
	Short: "get hostname of remote server-cli",
	Args:  cobra.NoArgs,
}

func initHostnameFuncs(hostnameClient hostnameClientInterface) {
	hostnameCmd.Run = func(cmd *cobra.Command, args []string) {
		hostname, err := hostnameClient.GetHostname(destination)
		if err != nil {
			log.Printf("ERR: %s\n", err.Error())
		} else {
			fmt.Println("Hostname: ", hostname)
		}
	}
	getHostnameCmd.Run = hostnameCmd.Run
	setHostnameCmd.Run = func(cmd *cobra.Command, args []string) {
		err := hostnameClient.UpdateHostname(destination, args[0])
		if err != nil {
			log.Printf("ERR: %s\n", err.Error())
		} else {
			fmt.Printf("Hostname %s successfully set", args[0])
		}
	}
}
