package main

import (
	"github.com/arnarpall/seccy/cmd/seccy-cli/cmd"
	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/pkg/client"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "seccy",
	Short: "Secrets keeper",
}

func main() {
	logger := log.Console()

	c, err := client.New(":4040", logger)
	if err != nil {
		logger.Fatalf("Unable to connect to seccy server, make sure that the server has been started", "error", err)
	}

	rootCmd.AddCommand(cmd.CreateGetCommand(logger, c))
	rootCmd.AddCommand(cmd.CreateSetCommand(logger, c))

	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}

