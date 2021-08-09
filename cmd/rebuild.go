package cmd

import (
	"github.com/mrxk/dev-env/pkg/config"
	"github.com/mrxk/dev-env/pkg/constants"
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func Rebuild(cmd *cobra.Command, _ []string) error {
	flags := cmd.PersistentFlags()
	noCache, err := flags.GetBool(constants.NoCacheOption)
	if err != nil {
		return err
	}
	args := []string{}
	if noCache {
		args = append(args, "--no-cache")
	}
	types, err := typesFromFlags(flags)
	if err != nil {
		return err
	}
	for _, t := range types {
		switch t {
		case constants.AllOption:
			err = rebuildConnectImage(cmd, args...)
			if err != nil {
				return err
			}
			err = rebuildRunImage(cmd)
			if err != nil {
				return err
			}
			return rebuildSpawnImage(cmd, args...)
		case constants.ConnectOption:
			return rebuildConnectImage(cmd, args...)
		case constants.RunOption:
			return rebuildRunImage(cmd, args...)
		case constants.SpawnOption:
			return rebuildSpawnImage(cmd, args...)
		}
	}
	return nil
}

func rebuildConnectImage(cmd *cobra.Command, args ...string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	return rebuildImage(env, args...)
}

func rebuildRunImage(cmd *cobra.Command, args ...string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	runEnv := env.WithName(env.Name() + constants.RunSuffix)
	return rebuildImage(runEnv, args...)
}

func rebuildSpawnImage(cmd *cobra.Command, args ...string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	spawnEnv := env.WithName(env.Name() + constants.SpawnSuffix)
	err = docker.StopContainer(spawnEnv)
	if err != nil {
		return err
	}
	return rebuildImage(spawnEnv, args...)
}

func rebuildImage(env *config.Env, args ...string) error {
	err := docker.RemoveContainer(env)
	if err != nil {
		return err
	}
	err = docker.RemoveImage(env)
	if err != nil {
		return err
	}
	return docker.BuildImage(env, args...)
}
