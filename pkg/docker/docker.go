package docker

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mrxk/dev-env/pkg/config"
)

func BuildImage(env *config.Env) error {
	dockerConfigDir, err := env.DockerBuildDir()
	if err != nil {
		return err
	}
	imageNameAndTag := env.ImageNameAndTag()
	dockerArgs := []string{
		"build",
		dockerConfigDir,
		"-t",
		imageNameAndTag,
	}
	fmt.Println(append([]string{"docker"}, dockerArgs...))
	dockerCommand := exec.Command("docker", dockerArgs...)
	dockerCommand.Stdout = os.Stdout
	dockerCommand.Stderr = os.Stderr
	err = dockerCommand.Start()
	if err != nil {
		return err
	}
	err = dockerCommand.Wait()
	if err != nil {
		return err
	}
	return nil
}

func BuildImageIfNotExist(env *config.Env) error {
	if ImageExists(env) {
		return nil
	}
	return BuildImage(env)
}

func ContainerExists(env *config.Env) bool {
	containerName := env.ContainerName()
	filter := fmt.Sprintf("name=^%s$", containerName)
	dockerCommand := exec.Command("docker", "ps", "--all", "-q", "-f", filter)
	output, err := dockerCommand.CombinedOutput()
	if err != nil {
		fmt.Printf("%v\n", err)
		return false
	}
	return len(output) != 0
}

func CreateContainer(env *config.Env, cmdArgs []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	containerName := env.ContainerName()
	dockerArgs := []string{
		"create",
		"-i",
		"-t",
		"-v",
		cwd + ":/src",
		"-w",
		"/src",
		"--name",
		containerName,
	}
	for _, containerArg := range env.ContainerArgs {
		expandedArg := os.ExpandEnv(containerArg)
		dockerArgs = append(dockerArgs, expandedArg)
	}
	imageNameAndTag := env.ImageNameAndTag()
	dockerArgs = append(dockerArgs, imageNameAndTag)
	for _, cmdArg := range cmdArgs {
		expandedArg := os.ExpandEnv(cmdArg)
		dockerArgs = append(dockerArgs, expandedArg)
	}
	fmt.Println(append([]string{"docker"}, dockerArgs...))
	dockerCommand := exec.Command("docker", dockerArgs...)
	dockerCommand.Stdout = os.Stdout
	dockerCommand.Stderr = os.Stderr
	err = dockerCommand.Start()
	if err != nil {
		return err
	}
	err = dockerCommand.Wait()
	if err != nil {
		return err
	}
	return nil
}

func CreateContainerIfNotExist(env *config.Env, cmdArgs []string) error {
	if ContainerExists(env) {
		return nil
	}
	return CreateContainer(env, cmdArgs)
}

func ImageExists(env *config.Env) bool {
	imageNameAndTag := env.ImageNameAndTag()
	dockerCommand := exec.Command("docker", "images", "-q", imageNameAndTag)
	output, err := dockerCommand.CombinedOutput()
	if err != nil {
		fmt.Printf("%v\n", err)
		return false
	}
	return len(output) != 0
}

func RecreateContainer(env *config.Env, cmdArgs []string) error {
	if ContainerExists(env) {
		RemoveContainer(env)
	}
	return CreateContainer(env, cmdArgs)
}

func RemoveContainer(env *config.Env) error {
	if !ContainerExists(env) {
		return nil
	}
	containerName := env.ContainerName()
	dockerArgs := []string{
		"rm",
		containerName,
	}
	fmt.Println(append([]string{"docker"}, dockerArgs...))
	dockerCommand := exec.Command("docker", dockerArgs...)
	dockerCommand.Stdout = os.Stdout
	dockerCommand.Stderr = os.Stderr
	err := dockerCommand.Start()
	if err != nil {
		return err
	}
	err = dockerCommand.Wait()
	if err != nil {
		return err
	}
	return nil
}

func RemoveImage(env *config.Env) error {
	if !ImageExists(env) {
		return nil
	}
	imageNameAndTag := env.ImageNameAndTag()
	dockerArgs := []string{
		"rmi",
		imageNameAndTag,
	}
	fmt.Println(append([]string{"docker"}, dockerArgs...))
	dockerCommand := exec.Command("docker", dockerArgs...)
	dockerCommand.Stdout = os.Stdout
	dockerCommand.Stderr = os.Stderr
	err := dockerCommand.Start()
	if err != nil {
		return err
	}
	err = dockerCommand.Wait()
	if err != nil {
		return err
	}
	return nil
}
