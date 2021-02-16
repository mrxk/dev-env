package cmd

import (
	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func Spawn(cmd *cobra.Command, _ []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	spawnEnv := env.WithName(env.Name + spawnSuffix)
	docker.WarnIfOutOfDate(spawnEnv)
	err = docker.BuildImageIfNotExist(spawnEnv)
	if err != nil {
		return err
	}
	containerArgs := append(spawnEnv.ContainerArgs, entrypointOption, spawnTailCommand)
	spawnEnv = spawnEnv.WithContainerArgs(containerArgs)
	cmdArgs := []string{spawnTailOptions, spawnTailArgs}
	err = docker.CreateContainerIfNotExist(spawnEnv, cmdArgs)
	if err != nil {
		return err
	}
	return docker.SpawnContainerIfNotRunning(spawnEnv)
}
