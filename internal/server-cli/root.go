package server_cli

import (
	"github.com/uroborosq/config-service/pkg/config"
	genDns "github.com/uroborosq/config-service/pkg/gen-dns"
	genHostname "github.com/uroborosq/config-service/pkg/gen-hostname"
	grpcServer "github.com/uroborosq/config-service/pkg/grpc-server"

	"context"
	grpcRuntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

var (
	grpcPort, gatewayPort, connectionName string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&grpcPort,
		"grpc-port",
		"G",
		"8080",
		"Port for grpc server")
	rootCmd.PersistentFlags().StringVarP(&gatewayPort,
		"gateway-port",
		"H",
		"8090",
		"Port for http gateway server")
	rootCmd.PersistentFlags().StringVarP(&connectionName,
		"connection-name",
		"C",
		"",
		"Connection's name in NetworkManager")
	_ = rootCmd.MarkPersistentFlagRequired("grpc-port")
	_ = rootCmd.MarkPersistentFlagRequired("gateway-port")
	_ = rootCmd.MarkPersistentFlagRequired("connection-name")

}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Run service on given ports",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		lsn, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatal(err.Error())
		}
		s := grpc.NewServer()
		genDns.RegisterDnsServiceServer(s, grpcServer.NewDnsGrpcServer(config.NewDnsService(connectionName)))
		genHostname.RegisterHostnameServiceServer(s, grpcServer.NewHostnameGrpcServer(config.NewHostnameService()))

		go func() {
			log.Println("Starting gRPC server-cli on ", lsn.Addr().String())
			log.Fatalln(s.Serve(lsn))
		}()

		gwMux := grpcRuntime.NewServeMux()
		conn, err := grpc.DialContext(
			context.Background(),
			":"+grpcPort,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)

		if err != nil {
			log.Fatal(err.Error())
		}
		err = genDns.RegisterDnsServiceHandler(context.Background(), gwMux, conn)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = genHostname.RegisterHostnameServiceHandler(context.Background(), gwMux, conn)
		if err != nil {
			log.Fatal(err.Error())
		}

		gwServer := http.NewServeMux()
		gwServer.HandleFunc("/swagger/swagger.json", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "../../api/config.swagger.json")
		})
		gwServer.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("../../web/swagger-ui/"))))
		gwServer.Handle("/", gwMux)

		log.Println("Starting HTTP gateway on", gatewayPort)
		log.Fatal(http.ListenAndServe(":"+gatewayPort, gwServer))
	},
}

func Execute() {
	rootCmd.Execute()
}
