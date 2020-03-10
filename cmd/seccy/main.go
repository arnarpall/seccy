package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arnarpall/seccy/cmd/seccy/cmd"
	"github.com/arnarpall/seccy/internal/encrypt"
	"github.com/arnarpall/seccy/internal/store/file"
	"github.com/arnarpall/seccy/pkg/vault"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "seccy",
	Short: "Seccy is a crazy secrets keeper",
}

func main() {
 	enc, err := encrypt.NoOp("my-key")
	if err != nil {
		log.Fatal(err)
	}

	store, err := file.NewFileStore(enc, "/tmp/arnar.vault")
	if err != nil {
		log.Fatal(err)
	}

	client := vault.NewClient(store)
	rootCmd.AddCommand(cmd.CreateGetCommand(client))
	rootCmd.AddCommand(cmd.CreateSetCommand(client))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
