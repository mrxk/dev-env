package cmd

import (
	"github.com/mrxk/dev-env/pkg/config"
	"github.com/mrxk/dev-env/pkg/constants"
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func Rebuild(cmd *cobra.Command, _ []string) error {
	types, err := cmd.PersistentFlags().GetStringSlice(constants.TypeOption)
	if err != nil {
		return err
	}
	for _, t := range types {
		switch t {
		case constants.AllType:
			err = rebuildConnectImage(cmd)
			if err != nil {
				return err
			}
			err = rebuildRunImage(cmd)
			if err != nil {
				return err
			}
			return rebuildSpawnImage(cmd)
		case constants.ConnectType:
			return rebuildConnectImage(cmd)
		case constants.RunType:
			return rebuildRunImage(cmd)
		case constants.SpawnType:
			return rebuildSpawnImage(cmd)
		}
	}
	return nil
}

func rebuildConnectImage(cmd *cobra.Command) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	return rebuildImage(env)
}

func rebuildRunImage(cmd *cobra.Command) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	runEnv := env.WithName(env.Name + constants.RunSuffix)
	return rebuildImage(runEnv)
}

func rebuildSpawnImage(cmd *cobra.Command) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	spawnEnv := env.WithName(env.Name + constants.SpawnSuffix)
	return rebuildImage(spawnEnv)
}

func rebuildImage(env *config.Env) error {
	err := docker.RemoveContainer(env)
	if err != nil {
		return err
	}
	err = docker.RemoveImage(env)
	if err != nil {
		return err
	}
	return docker.BuildImage(env)
}
