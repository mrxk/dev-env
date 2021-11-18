package cmd

import (
	"github.com/mrxk/dev-env/pkg/config"
	"github.com/mrxk/dev-env/pkg/constants"
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func RemoveContainer(cmd *cobra.Command, _ []string) error {
	types, err := typesFromFlags(cmd.PersistentFlags())
	if err != nil {
		return err
	}
	for _, t := range types {
		switch t {
		case constants.AllOption:
			err = removeConnectContainer(cmd)
			if err != nil {
				return err
			}
			err = removeRunContainer(cmd)
			if err != nil {
				return err
			}
			return removeSpawnContainer(cmd)
		case constants.ConnectOption:
			return removeConnectContainer(cmd)
		case constants.RunOption:
			return removeRunContainer(cmd)
		case constants.SpawnOption:
			return removeSpawnContainer(cmd)
		}
	}
	return nil
}

func removeConnectContainer(cmd *cobra.Command) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	return removeContainer(env)
}

func removeRunContainer(cmd *cobra.Command) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	env = env.WithName(env.Name() + constants.RunSuffix)
	return removeContainer(env)
}

func removeSpawnContainer(cmd *cobra.Command) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	env = env.WithName(env.Name() + constants.SpawnSuffix)
	return removeContainer(env)
}

func removeContainer(env *config.Env) error {
	err := docker.StopContainer(env)
	if err != nil {
		return err
	}
	return docker.RemoveContainer(env)
}
