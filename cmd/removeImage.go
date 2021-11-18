package cmd

import (
	"github.com/mrxk/dev-env/pkg/config"
	"github.com/mrxk/dev-env/pkg/constants"
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func RemoveImage(cmd *cobra.Command, _ []string) error {
	types, err := typesFromFlags(cmd.PersistentFlags())
	if err != nil {
		return err
	}
	for _, t := range types {
		switch t {
		case constants.AllOption:
			err = removeConnectImage(cmd)
			if err != nil {
				return err
			}
			err = removeRunImage(cmd)
			if err != nil {
				return err
			}
			return removeSpawnImage(cmd)
		case constants.ConnectOption:
			return removeConnectImage(cmd)
		case constants.RunOption:
			return removeRunImage(cmd)
		case constants.SpawnOption:
			return removeSpawnImage(cmd)
		}
	}
	return nil
}

func removeConnectImage(cmd *cobra.Command) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	return removeImage(env)
}

func removeRunImage(cmd *cobra.Command) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	return removeImage(env.WithName(env.Name() + constants.RunSuffix))
}

func removeSpawnImage(cmd *cobra.Command) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	spawnEnv := env.WithName(env.Name() + constants.SpawnSuffix)
	return removeImage(spawnEnv)
}

func removeImage(env *config.Env) error {
	err := docker.StopContainer(env)
	if err != nil {
		return err
	}
	err = docker.RemoveContainer(env)
	if err != nil {
		return err
	}
	return docker.RemoveImage(env)
}
