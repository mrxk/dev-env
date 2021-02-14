package cmd

import (
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func RemoveContainer(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	return docker.RemoveContainer(env)
}

func RemoveRContainer(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	runEnv := env.WithName(env.Name + "_r")
	return docker.RemoveContainer(runEnv)
}

func RemoveSContainer(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	spawnEnv := env.WithName(env.Name + "_spawn")
	err = docker.StopContainer(spawnEnv)
	if err != nil {
		return err
	}
	return docker.RemoveContainer(spawnEnv)
}
