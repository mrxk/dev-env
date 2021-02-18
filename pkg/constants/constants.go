package constants

const (
	AllType           = "all"
	ConnectType       = "connect"
	DefaultDockerFile = `
FROM debian:stretch
LABEL creator=dev-env
ENTRYPOINT ["/bin/bash"]
`
	DefaultEnvironment          = "main"
	DetachedOption              = "detached"
	EntrypointOption            = "--entrypoint"
	EnvironmentOption           = "environment"
	RunShellCommand             = "bash"
	RunSuffix                   = "_run"
	RunType                     = "run"
	ShellCommandOption          = "-c"
	SkipOutOfDateWarningsOption = "skip-out-of-date-warnings"
	SpawnSuffix                 = "_spawn"
	SpawnTailArgs               = "/dev/null"
	SpawnTailCommand            = "tail"
	SpawnTailOptions            = "-f"
	SpawnType                   = "spawn"
	TypeOption                  = "type"
)

var (
	DefaultContainerArgs = []string{"-v", "${PROJECTROOT}:/src", "-w", "/src"}
)
