// +build !windows

package docker

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/mrxk/dev-env/pkg/config"
)

func StartContainer(env *config.Env, interactive bool) error {
	dockerArgs := []string{
		"docker",
		"start",
	}
	if interactive {
		dockerArgs = append(dockerArgs, "-i")
	}
	containerName := env.ContainerName()
	dockerArgs = append(dockerArgs, containerName)
	dockerBinary, err := exec.LookPath("docker")
	if err != nil {
		return err
	}
	fmt.Println(dockerArgs)
	return syscall.Exec(dockerBinary, dockerArgs, os.Environ())
}

func ExecContainer(env *config.Env, args []string, interactive bool) error {
	dockerArgs := []string{
		"docker",
		"exec",
	}
	if !interactive {
		dockerArgs = append(dockerArgs, "-d")
	}
	containerName := env.ContainerName()
	dockerArgs = append(dockerArgs, containerName)
	dockerArgs = append(dockerArgs, args...)
	dockerBinary, err := exec.LookPath("docker")
	if err != nil {
		return err
	}
	fmt.Println(dockerArgs)
	return syscall.Exec(dockerBinary, dockerArgs, os.Environ())
}

func dockerSocket() string {
	return "/var/run/docker.sock"
}
