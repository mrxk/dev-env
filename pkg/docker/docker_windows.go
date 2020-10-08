package docker

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mrxk/dev-env/pkg/config"
)

func StartContainer(env *config.Env, interactive bool) error {
	dockerArgs := []string{
		"start",
	}
	if interactive {
		dockerArgs = append(dockerArgs, "-i")
	}
	containerName := env.ContainerName()
	dockerArgs = append(dockerArgs, containerName)
	fmt.Println(append([]string{"docker"}, dockerArgs...))
	dockerCommand := exec.Command("docker", dockerArgs...)
	dockerCommand.Stdout = os.Stdout
	dockerCommand.Stderr = os.Stderr
	dockerCommand.Stdin = os.Stdin
	err := dockerCommand.Start()
	if err != nil {
		return err
	}
	err = dockerCommand.Wait()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		return err
	}
	os.Exit(0)
	return nil
}
