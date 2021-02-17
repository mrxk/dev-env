package cmd

import (
	"github.com/mrxk/dev-env/pkg/constants"
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func RemoveContainer(cmd *cobra.Command, _ []string) error {
	types, err := cmd.PersistentFlags().GetStringSlice(constants.TypeOption)
	if err != nil {
		return err
	}
	for _, t := range types {
		switch t {
		case constants.AllType:
			err = removeConnectContainer(cmd)
			if err != nil {
				return err
			}
			err = removeRunContainer(cmd)
			if err != nil {
				return err
			}
			return removeSpawnContainer(cmd)
		case constants.ConnectType:
			return removeConnectContainer(cmd)
		case constants.RunType:
			return removeRunContainer(cmd)
		case constants.SpawnType:
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
	return docker.RemoveContainer(env)
}

func removeRunContainer(cmd *cobra.Command) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	runEnv := env.WithName(env.Name + constants.RunSuffix)
	err = docker.StopContainer(runEnv)
	if err != nil {
		return err
	}
	return docker.RemoveContainer(runEnv)
}

func removeSpawnContainer(cmd *cobra.Command) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	spawnEnv := env.WithName(env.Name + constants.SpawnSuffix)
	err = docker.StopContainer(spawnEnv)
	if err != nil {
		return err
	}
	return docker.RemoveContainer(spawnEnv)
}
