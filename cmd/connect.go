package cmd

import (
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func Connect(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	err = docker.BuildImageIfNotExist(env)
	if err != nil {
		return err
	}
	err = docker.CreateContainerIfNotExist(env, nil)
	if err != nil {
		return err
	}
	return docker.StartContainer(env, true)
}
