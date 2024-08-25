package cmd

import (
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run HTTP and gRPC service",
	Long:  `Run HTTP and gRPC service`,
	Run:   RunServer,
}

func RunServer(cmd *cobra.Command, args []string) {
	// TODO(nick):
}
