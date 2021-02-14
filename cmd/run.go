package cmd

import (
	"fmt"
	"strings"

	"github.com/mrxk/dev-env/pkg/docker"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	env, err := envFromFlags(cmd.Flags())
	if err != nil {
		return err
	}
	runEnv := env.WithName(env.Name + "_r")
	err = docker.BuildImageIfNotExist(runEnv)
	if err != nil {
		return err
	}
	containerArgs := append(runEnv.ContainerArgs, "--entrypoint", "bash")
	runEnv = runEnv.WithContainerArgs(containerArgs)
	cmdArgs := []string{"-c", fmt.Sprintf("%s", strings.Join(args, " "))}
	err = docker.RecreateContainer(runEnv, cmdArgs)
	if err != nil {
		return err
	}
	detached, err := cmd.PersistentFlags().GetBool("detached")
	if err != nil {
		return err
	}
	return docker.StartContainer(runEnv, !detached)
}
