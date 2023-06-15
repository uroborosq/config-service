package client_cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var destination string

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "Client to manage hostname and dns servers on remote server-cli",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use --help or -h to get usage information")
	},
}

func Execute(newDnsClient dnsClientInterface, newHostnameClient hostnameClientInterface) {
	rootCmd.PersistentFlags().StringVarP(&destination,
		"destination",
		"d",
		"",
		"Address of remote server-cli (required)")
	_ = rootCmd.MarkPersistentFlagRequired("destination")

	initDnsFuncs(newDnsClient)
	initHostnameFuncs(newHostnameClient)

	rootCmd.AddCommand(dnsCmd)
	rootCmd.AddCommand(hostnameCmd)

	dnsCmd.AddCommand(getDnsCmd, addDnsCmd, removeDnsCommand)
	hostnameCmd.AddCommand(getHostnameCmd, setHostnameCmd)

	rootCmd.Execute()
}
