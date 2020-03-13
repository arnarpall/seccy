package cmd

import (
	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/pkg/client"
	"github.com/spf13/cobra"
)

type setCommand struct {
	setter client.Setter
	logger *log.Logger
	key    string
	value  string
}

func CreateSetCommand(logger *log.Logger, setter client.Setter) *cobra.Command {
	set := setCommand{
		setter: setter,
		logger: logger,
	}

	c := &cobra.Command{
		Use:   "set",
		Short: "set a secret value",
		Run: func(cmd *cobra.Command, args []string) {
			set.run()
		},
	}

	c.Flags().StringVar(&set.key, "key", "", "key to store the supplied value")
	_ = c.MarkFlagRequired("key")
	c.Flags().StringVar(&set.value, "value", "", "value to set")
	_ = c.MarkFlagRequired("value")

	return c
}

func (s *setCommand) run() {
	err := s.setter.Set(s.key, s.value)
	if err != nil {
		s.logger.Warnf("Could not set key %s with value %s, %v", s.key, s.value, err)
		return
	}
}
