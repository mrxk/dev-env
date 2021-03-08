package cmd

import (
	"github.com/mrxk/dev-env/pkg/config"
	"github.com/mrxk/dev-env/pkg/constants"
	"github.com/spf13/cobra"
)

func Initialize(cmd *cobra.Command, _ []string) error {
	flags := cmd.Flags()
	envName, err := envNameFromFlags(flags)
	if err != nil {
		return err
	}
	if envName == "" {
		envName = constants.DefaultEnvironment
	}
	cfg, err := config.NewConfigFor(envName)
	if err != nil {
		return err
	}
	return cfg.WriteIfNotExist()
}
