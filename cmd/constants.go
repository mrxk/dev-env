package cmd

const (
	defaultDockerFile = `
FROM debian:stretch
ENTRYPOINT ["/bin/bash"]
`
	defaultEnvironment = "main"
	detachedOption     = "detached"
	entrypointOption   = "--entrypoint"
	environmentOption  = "env"
	runShellCommand    = "bash"
	runSuffix          = "_run"
	shellCommandOption = "-c"
	spawnSuffix        = "_spawn"
	spawnTailArgs      = "/dev/null"
	spawnTailCommand   = "tail"
	spawnTailOptions   = "-f"
)
