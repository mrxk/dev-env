package cmd

import (
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func Stop(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	spawnEnv := env.WithName(env.Name + spawnSuffix)
	return docker.StopContainer(spawnEnv)
}
