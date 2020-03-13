package cmd

import (
	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/pkg/client"
	"github.com/spf13/cobra"
)

type getCommand struct {
	logger *log.Logger
	getter client.Getter
	key string
}

func CreateGetCommand(logger *log.Logger, getter client.Getter) *cobra.Command {
	get := &getCommand{
		getter: getter,
		logger: logger,
	}

	c := &cobra.Command{
		Use:   "get",
		Short: "get a secret",
		Run: func(cmd *cobra.Command, args []string) {
			get.run()
		},
	}

	c.Flags().StringVar(&get.key, "key", "", "key to the the value for")
	_ = c.MarkFlagRequired("key")
	return c
}

func (g *getCommand) run() {
	v, err := g.getter.Get(g.key)
	if err != nil {
		g.logger.Warnf("Could not get %s, %v", g.key, err)
		return
	}

	g.logger.Info(v)
}
