package cmd

import (
	"github.com/mrxk/dev-env/pkg/config"
	"github.com/mrxk/dev-env/pkg/constants"
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func Build(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	return rebuild(env)
}

func BuildR(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	runEnv := env.WithName(env.Name + constants.RunSuffix)
	return rebuild(runEnv)
}

func BuildS(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	spawnEnv := env.WithName(env.Name + constants.SpawnSuffix)
	return rebuild(spawnEnv)
}

func rebuild(env *config.Env) error {
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
