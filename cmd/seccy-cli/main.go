package main

import (
	"github.com/arnarpall/seccy/cmd/seccy-cli/cmd"
	"github.com/arnarpall/seccy/internal/env"
	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/internal/server"
	"github.com/arnarpall/seccy/pkg/client"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "seccy",
	Short: "Secrets keeper",
}

func main() {
	logger := log.Console()

	serverAddress := env.GetEnvOrDefault("SECCY_SERVER", server.DefaultServerAddress)

	c, err := client.New(serverAddress, logger)
	if err != nil {
		logger.Fatalf("Unable to connect to seccy server, make sure that the server has been started", "error", err)
	}

	rootCmd.AddCommand(cmd.CreateGetCommand(logger, c))
	rootCmd.AddCommand(cmd.CreateSetCommand(logger, c))
	rootCmd.AddCommand(cmd.CreateListKeysCommand(logger, c))

	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}
