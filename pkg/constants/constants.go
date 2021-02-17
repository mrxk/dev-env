package constants

const (
	DefaultDockerFile = `
FROM debian:stretch
LABEL creator=dev-env
ENTRYPOINT ["/bin/bash"]
`
	DefaultEnvironment = "main"
	DetachedOption     = "detached"
	EntrypointOption   = "--entrypoint"
	EnvironmentOption  = "env"
	RunShellCommand    = "bash"
	RunSuffix          = "_run"
	ShellCommandOption = "-c"
	SpawnSuffix        = "_spawn"
	SpawnTailArgs      = "/dev/null"
	SpawnTailCommand   = "tail"
	SpawnTailOptions   = "-f"
)

var (
	DefaultContainerArgs = []string{"-v", "${PROJECTROOT}:/src", "-w", "/src"}
)
