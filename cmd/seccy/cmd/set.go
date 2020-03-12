package cmd

import (
	"fmt"

	"github.com/arnarpall/seccy/pkg/client"
	"github.com/spf13/cobra"
)

func CreateSetCommand(setter client.Setter) *cobra.Command {
	return &cobra.Command{
		Use:   "set",
		Short: "set a secret value",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]
			val := args[1]
			err := setter.Set(key, val)
			if err != nil {
				fmt.Printf("Could not set key %s with value %s, %v\n", key, val, err)
				return
			}
		},
	}
}
