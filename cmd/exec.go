package cmd

import (
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func Exec(cmd *cobra.Command, args []string) error {
	err := Spawn(cmd, nil)
	if err != nil {
		return err
	}
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	runEnv := env.WithName(env.Name + spawnSuffix)
	detached, err := cmd.PersistentFlags().GetBool(detachedOption)
	if err != nil {
		return err
	}
	return docker.ExecContainer(runEnv, args, !detached)
}
