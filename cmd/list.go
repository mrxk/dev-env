package cmd

import (
	"fmt"

	"github.com/mrxk/dev-env/pkg/config"
	"github.com/mrxk/dev-env/pkg/constants"
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func List(cmd *cobra.Command, args []string) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	names := cfg.EnvNames()
	for _, name := range names {
		env := cfg.Envs[name]
		runEnv := env.WithName(env.Name() + constants.RunSuffix)
		spawnEnv := env.WithName(env.Name() + constants.SpawnSuffix)
		fmt.Printf("%s:\n", name)
		listForEnv("connect", env)
		listForEnv("run", runEnv)
		listForEnv("spawn", spawnEnv)
	}
	return nil
}

func listForEnv(name string, env *config.Env) {
	if docker.ImageExists(env) {
		fmt.Printf("    %s:\n", name)
		fmt.Printf("        image    : %s\n", env.ImageNameAndTag())
	}
	if docker.ContainerExists(env) {
		fmt.Printf("        container: %s\n", env.ContainerName())
	}
}
