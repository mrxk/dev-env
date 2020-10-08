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
	interactive, err := cmd.PersistentFlags().GetBool("interactive")
	if err != nil {
		return err
	}
	err = docker.StartContainer(runEnv, interactive)
	if err != nil {
		return nil
	}
	return nil
}
