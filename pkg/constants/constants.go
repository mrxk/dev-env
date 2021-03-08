package constants

const (
	AllType           = "all"
	ConfigDir         = ".dev-env"
	ConfigFile        = "dev-env.json"
	ConnectType       = "connect"
	DefaultDockerFile = `
FROM debian:stretch
LABEL creator=dev-env
ENTRYPOINT ["/bin/bash"]
`
	DefaultEnvironment      = "main"
	DetachedOption          = "detached"
	DevEnvTag               = "dev_env"
	DockerFile              = "Dockerfile"
	EntrypointOption        = "--entrypoint"
	EnvironmentOption       = "environment"
	OutOfDateWarningsOption = "out-of-date-warnings"
	RunShellCommand         = "bash"
	RunSuffix               = "_run"
	RunType                 = "run"
	ShellCommandOption      = "-c"
	SpawnSuffix             = "_spawn"
	SpawnTailArgs           = "/dev/null"
	SpawnTailCommand        = "tail"
	SpawnTailOptions        = "-f"
	SpawnType               = "spawn"
	TypeOption              = "type"
)

var (
	DefaultContainerArgs = []string{"-v", "${PROJECTROOT}:/src", "-w", "/src"}
)
