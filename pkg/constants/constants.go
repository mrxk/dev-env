package constants

const (
	AllOption          = "all"
	AllShortOption     = "a"
	ConfigDir          = ".dev-env"
	ConfigFile         = "dev-env.json"
	ConnectOption      = "connect"
	ConnectShortOption = "c"
	DefaultDockerFile  = `
FROM alpine:latest
LABEL creator=dev-env
ENTRYPOINT ["/bin/ash"]
`
	DefaultEnvironment      = "main"
	DetachedOption          = "detached"
	DetachedShortOption     = "d"
	DevEnvTag               = "dev_env"
	DockerFile              = "Dockerfile"
	EntrypointOption        = "--entrypoint"
	EnvironmentOption       = "environment"
	EnvironmentShortOption  = "e"
	NoCacheOption           = "no-cache"
	NoCacheShortOption      = "n"
	OutOfDateWarningsOption = "out-of-date-warnings"
	RunOption               = "run"
	RunShellCommand         = "sh"
	RunShortOption          = "r"
	RunSuffix               = "_run"
	ShellCommandOption      = "-c"
	SpawnOption             = "spawn"
	SpawnShortOption        = "s"
	SpawnSuffix             = "_spawn"
	SpawnTailArgs           = "/dev/null"
	SpawnTailCommand        = "tail"
	SpawnTailOptions        = "-f"
)

var (
	DefaultContainerArgs = []string{"-v", "${PROJECTROOT}:/src", "-w", "/src"}
)
