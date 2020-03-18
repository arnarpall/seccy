package cmd

import (
	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/pkg/client"
	"github.com/spf13/cobra"
)

type listKeysCommand struct {
	lister client.Lister
	logger *log.Logger
	key    string
	value  string
}

func CreateListKeysCommand(logger *log.Logger, lister client.Lister) *cobra.Command {
	set := listKeysCommand{
		lister: lister,
		logger: logger,
	}

	c := &cobra.Command{
		Use:   "list",
		Short: "list all keys",
		Run: func(cmd *cobra.Command, args []string) {
			set.run()
		},
	}

	return c
}

func (s *listKeysCommand) run() {
	keyChan, err := s.lister.ListKeys()
	if err != nil {
		s.logger.Warnf("Could not list all keys, %v", s.key, err)
		return
	}

	for key := range keyChan {
		s.logger.Info(key)
	}
}
