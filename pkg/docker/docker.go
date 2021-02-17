package docker

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

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

func ContainerCreationTime(env *config.Env) (time.Time, error) {
	containerName := env.ContainerName()
	dockerCommand := exec.Command("docker", "inspect", containerName, "--format", "{{.Created}}")
	output, err := dockerCommand.CombinedOutput()
	if err != nil {
		return time.Time{}, err
	}
	trimmedOutput := strings.TrimSpace(string(output))
	return time.Parse("2006-01-02T15:04:05.999999999Z", trimmedOutput)
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

func ContainerRunning(env *config.Env) bool {
	containerName := env.ContainerName()
	filter := fmt.Sprintf("name=^%s$", containerName)
	dockerCommand := exec.Command("docker", "ps", "-q", "-f", filter)
	output, err := dockerCommand.CombinedOutput()
	if err != nil {
		fmt.Printf("%v\n", err)
		return false
	}
	return len(output) != 0
}

func CreateContainer(env *config.Env, cmdArgs []string) error {
	containerName := env.ContainerName()
	dockerArgs := []string{
		"create",
		"-i",
		"-t",
		"--name",
		containerName,
	}
	for _, containerArg := range env.ContainerArgs {
		expandedArg := os.Expand(containerArg, expand)
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

func CreateContainerIfNotExist(env *config.Env, cmdArgs []string) error {
	if ContainerExists(env) {
		return nil
	}
	return CreateContainer(env, cmdArgs)
}

func DockerFileModTime(env *config.Env) (time.Time, error) {
	dockerConfigDir, err := env.DockerBuildDir()
	if err != nil {
		return time.Time{}, err
	}
	dockerFilePath := filepath.Join(dockerConfigDir, "Dockerfile")
	stat, err := os.Stat(dockerFilePath)
	if err != nil {
		return time.Time{}, err
	}
	return stat.ModTime(), nil
}

func expand(name string) string {
	normalizedName := strings.ToUpper(name)
	switch normalizedName {
	case "PROJECTROOT":
		root, err := config.GetProjectRoot()
		if err == nil {
			return root
		}
		fallthrough
	default:
		return os.Getenv(name)
	}
}

func ImageCreationTime(env *config.Env) (time.Time, error) {
	imageNameAndTag := env.ImageNameAndTag()
	dockerCommand := exec.Command("docker", "images", imageNameAndTag, "--format", "{{.CreatedAt}}")
	output, err := dockerCommand.CombinedOutput()
	if err != nil {
		return time.Time{}, err
	}
	trimmedOutput := strings.TrimSpace(string(output))
	return time.Parse("2006-01-02 15:04:05 -0700 MST", trimmedOutput)
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

func SpawnContainer(env *config.Env) error {
	containerName := env.ContainerName()
	dockerArgs := []string{
		"start",
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
	return dockerCommand.Wait()
}

func SpawnContainerIfNotRunning(env *config.Env) error {
	if ContainerRunning(env) {
		return nil
	}
	return SpawnContainer(env)
}

func StopContainer(env *config.Env) error {
	if !ContainerExists(env) {
		return nil
	}
	containerName := env.ContainerName()
	dockerArgs := []string{
		"stop",
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
	return dockerCommand.Wait()
}

func WarnIfContainerOutOfDate(env *config.Env) {
	if !ContainerExists(env) {
		return
	}
	containerCreationTime, err := ContainerCreationTime(env)
	if err != nil {
		fmt.Println("WARNING: ", err)
		return
	}
	dockerFileModTime, err := DockerFileModTime(env)
	if err != nil {
		fmt.Println("WARNING: ", err)
		return
	}
	if dockerFileModTime.After(containerCreationTime) {
		fmt.Println("WARNING: Dockerfile modified after container created.")
	}
}

func WarnIfImageOutOfDate(env *config.Env) {
	if !ImageExists(env) {
		return
	}
	imageCreationTime, err := ImageCreationTime(env)
	if err != nil {
		fmt.Println("WARNING: ", err)
		return
	}
	dockerFileModTime, err := DockerFileModTime(env)
	if err != nil {
		fmt.Println("WARNING: ", err)
		return
	}
	if dockerFileModTime.After(imageCreationTime) {
		fmt.Println("WARNING: Dockerfile modified after image created.")
	}
}

func WarnIfOutOfDate(env *config.Env) {
	WarnIfContainerOutOfDate(env)
	WarnIfImageOutOfDate(env)
}
