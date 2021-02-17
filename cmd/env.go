package cmd

import (
	"github.com/mrxk/dev-env/pkg/config"
	"github.com/mrxk/dev-env/pkg/constants"
	"github.com/spf13/pflag"
)

func envFromFlags(flags *pflag.FlagSet) (*config.Env, error) {
	envName, err := flags.GetString(constants.EnvironmentOption)
	if err != nil {
		return nil, err
	}
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	return cfg.EnvFor(envName)
}
