package main

import (
	"fmt"
	"os"

	"github.com/arnarpall/seccy/cmd/seccy/cmd"
	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/pkg/client"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "seccy",
	Short: "Client is a crazy secrets keeper",
}

func main() {
	logger := log.New()
	c, err := client.New(":4040", logger)
	if err != nil {
		logger.Fatalf("Unable to connect to seccy server, make sure that the server has been started", "error", err)
	}

	rootCmd.AddCommand(cmd.CreateGetCommand(c))
	rootCmd.AddCommand(cmd.CreateSetCommand(c))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
