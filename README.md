# Dev-env

Dev-env is a cli for managing a contanerized development environment with docker. It is
inspired by the
[Visual Studio Code Remote - Containers](https://code.visualstudio.com/docs/remote/containers)
extension.

## Installation

```bash
go get github.com/mrxk/dev-env
```

## Quick start

The following command will create and connect to a simple `debian:stretch`
container. The default configuration will be created.

```bash
dev-env connect
```

## Usage

```text
Manage a developement environment via docker in the current working directory.

Usage:
  dev-env [command]

Available Commands:
  build       Build a dev-env image in the current directory
  connect     Start or connect to a dev-env container in the current directory
  exec        Exec a command in a spawned dev-env container in the current directory
  help        Help about any command
  init        Initialize a dev-env in the current directory
  rm          Remove a dev-env container in the current directory
  rmi         Remove a dev-env container and its associated image in the current directory
  rmr         Remove a run dev-env container in the current directory
  rmri        Remove a run dev-env container and its associated image in the current directory
  rms         Remove a spawn dev-env container in the current directory
  rmsi        Remove a spawn dev-env container and its associated image in the current directory
  run         Run a command via bash in a dev-env container in the current directory
  spawn       Spawn a detached dev-env container in the current directory

Flags:
  -e, --env string   environment to use (default "main")
  -h, --help         help for dev-env

Use "dev-env [command] --help" for more information about a command.
```

## Configuration

Dev-env expects to find a `.dev-env` directory in the current working directory
containing at least one directory. The name of that directory will be the name
of the environment. Inside that directory must be a `dev-env.json` file and a
`Dockerfile`. This directory is where `docker build` will be run to create the
container. The `main` environment will be created with a default `Dockerfile`
will be created if they do not exist. The resulting image must contain
functional `bash` for the `run` command and a functional `tail` for the `spawn`
command.

The `dev-env.json` config file consists of a single top level object with the
following fields.

| field | type | description |
|-------|------|-------------|
| containerArgs | string array | A list of arguments to pass to docker create. Each argument in this list is passed to `os.ExpandEnv` before being added to the docker command. Optional. |
| name | string| The name of the container. Required. Init will create a name for the default `main` environment. |

### Example config file

```json
{
    "containerArgs": [
	"--cpuset-cpus", "1,3",
	"-v", "$HOME/.ssh:/home/user/.ssh"
    ],
    "name": "nervous_golick"
}
```

## Details

Dev-env supports three styles of development environment management.

### Connect style

The following commands work together to manage a persistent interactive shell
environment.

 * `build`: Build a dev-env image in the current directory
 * `connect`: Start or connect to a dev-env container in the current directory
 * `rm`: Remove a dev-env container in the current directory
 * `rmi`: Remove a dev-env container and its associated image in the current directory

The `connect` command creates an image (if one does not already exist) from the
given environment with the current working directory mounted in `/src`. It sets
the working directory of the image to be `/src`. It then creates a container
from this image (if one does not already exist) and starts the container. That
container will be re-used on subsequent `connect`s until it is removed with
`rm`. The `build` command can be used to force this image to built. Any
existing containers and images will be removed. The container can be destroyed
with the `rm` command. The image can be destroyed with the `rmi` command. The
image and container maintained by these commands is distinct from the one
managed by the `run` and `spawn`family of commands.

### Run style

The following commands work together to create an isolated environment for
running a single command that is isolated from other command invocations.

 * `run`: Run a command via bash in a dev-env container in the current directory
 * `rmr`: Remove a run dev-env container in the current directory
 * `rmri`: Remove a run dev-env container and its associated image in the current directory

The `run` command creates an image from the given environment (if one does not
already exist) and then destroys and re-creates a single use container for the
execution of a single command. The image is assumed to contain `bash` and the
arguments passed to the `run` command will be passed to `bash -c` in the
container. If the arguments themselves contain flags then separate the entire
command with double dashes `--`. For example, `dev-env run -- ls -la /`. This
container and image can be destroyed with the `rmr` and `rmri` commands. If
`--detached` is passed to `run` then the container will be detached and control
will return to the shell. The container will exit when the command is complete.
Because the container is recreated each time, changes made in the container are
not persistent. The image and container maintained by these commands is
distinct from the one managed by the `connect` and `spawn`family of commands.

### Spawn style

The following commands work together to create a persistant environment for
running single commands.

 * `spawn`: Spawn a detached dev-env container in the current directory
 * `exec`: Exec a command via bash in a spawned dev-env container in the current directory
 * `rms`: Remove a spawn dev-env container in the current directory
 * `rmsi`: Remove a spawn dev-env container and its associated image in the current directory

The `spawn` command creates an image from the given environment (if one does
not already exist) and then creates and starts a persistent container in the
background. The image is assumed to contain `tail` and `tail -f /dev/null` is
used to keep the container from exiting. The arguments to the `exec` command
will be executed in this spawned container. If the arguments themselves contain
flags then separate the entire command with double dashes `--`. For example,
`dev-env exec -- ls -la /`. This container and image can be destroyed with the
`rms` and `rmsi` commands. If `--detached` is passed to `exec` then the
container will be detached and control will return to the shell. The container
will exit when the command is complete. Because the container is persistent ,
changes made in the container are persistent. The image and container
maintained by these commands is distinct from the one managed by the `connect`
and `run` family of commands.

## License
[MIT](https://choosealicense.com/licenses/mit/)
