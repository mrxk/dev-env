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
	err = docker.RemoveContainer(env)
	if err != nil {
		return err
	}
	return nil
}

func RemoveRContainer(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	runEnv := env.WithName(env.Name + "_r")
	err = docker.RemoveContainer(runEnv)
	if err != nil {
		return err
	}
	return nil
}
