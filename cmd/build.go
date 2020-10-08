package cmd

import (
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func Build(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	err = docker.RemoveContainer(env)
	if err != nil {
		return err
	}
	err = docker.RemoveImage(env)
	if err != nil {
		return err
	}
	return docker.BuildImage(env)
}
