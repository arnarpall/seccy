package cmd

import (
	"fmt"

	"github.com/arnarpall/seccy/pkg/vault"
	"github.com/spf13/cobra"
)

func CreateGetCommand(getter vault.Getter) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "get a secret",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]
			value, err := getter.Get(key)
			if err != nil {
				fmt.Printf("Could not get %s, %v\n", key, err)
				return
			}

			fmt.Printf("%s\n", value)
		},
	}
}
