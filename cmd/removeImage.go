package cmd

import (
	"github.com/mrxk/dev-env/pkg/config"
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func RemoveImage(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	return removeImage(env)
}

func RemoveRImage(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	return removeImage(env.WithName(env.Name + runSuffix))
}

func RemoveSImage(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	spawnEnv := env.WithName(env.Name + spawnSuffix)
	err = docker.StopContainer(spawnEnv)
	if err != nil {
		return err
	}
	return removeImage(spawnEnv)
}

func removeImage(env *config.Env) error {
	err := docker.RemoveContainer(env)
	if err != nil {
		return err
	}
	return docker.RemoveImage(env)
}
