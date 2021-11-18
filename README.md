# Dev-env

Dev-env is a cli for managing a contanerized development environment with
docker. It is inspired by the [Visual Studio Code Remote -
Containers](https://code.visualstudio.com/docs/remote/containers) extension.

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

### dev-env

```text
Manage a developement environment via docker in the current working directory.

Usage:
  dev-env [command]

Available Commands:
  connect     Connect to a connect dev-env container in the current directory
  exec        Exec a command in a spawned dev-env container in the current directory
  help        Help about any command
  init        Initialize a dev-env in the current directory
  list        List containers and images in the current directory
  rebuild     Rebuild dev-env images in the current directory
  rm          Remove dev-env containers in the current directory
  rmi         Remove dev-env containers and their associated images in the current directory
  run         Run a command via sh in a run dev-env container in the current directory
  spawn       Start a spawn dev-env container in the current directory
  stop        Stop a spawned dev-env container in the current directory

Flags:
  -e, --environment string   environment to use (default "main")
  -h, --help                 help for dev-env

Use "dev-env [command] --help" for more information about a command.
```

### dev-env connect

```text
Connect to a connect dev-env container in the current directory

Usage:
  dev-env connect [flags]

Flags:
  -h, --help   help for connect

Global Flags:
  -e, --environment string   environment to use (default "main")
```

### dev-env exec

```text
Exec a command in a spawned dev-env container in the current directory

Usage:
  dev-env exec [flags]

Flags:
  -d, --detached   run command in the background
  -h, --help       help for exec

Global Flags:
  -e, --environment string   environment to use (default "main")
```

### dev-env help

```text
Help provides help for any command in the application.
Simply type dev-env help [path to command] for full details.

Usage:
  dev-env help [command] [flags]

Flags:
  -h, --help   help for help

Global Flags:
  -e, --environment string   environment to use (default "main")
```

### dev-env init

```text
Initialize a dev-env in the current directory

Usage:
  dev-env init [flags]

Flags:
  -h, --help   help for init

Global Flags:
  -e, --environment string   environment to use (default "main")
```

### dev-env list

```text
List containers and images in the current directory

Usage:
  dev-env list [flags]

Aliases:
  list, ls

Flags:
  -h, --help   help for list

Global Flags:
  -e, --environment string   environment to use (default "main")
```

### dev-env rebuild

```text
Rebuild dev-env images in the current directory

Usage:
  dev-env rebuild [flags]

Flags:
  -a, --all       rebuild all images
  -c, --connect   rebuild connect image
  -h, --help      help for rebuild
  -r, --run       rebuild run image
  -s, --spawn     rebuild spawn image

Global Flags:
  -e, --environment string   environment to use (default "main")
```

### dev-env rm

```text
Remove dev-env containers in the current directory

Usage:
  dev-env rm [flags]

Flags:
  -a, --all       remove all containers
  -c, --connect   remove connect container
  -h, --help      help for rm
  -r, --run       remove run container
  -s, --spawn     remove spawn container

Global Flags:
  -e, --environment string   environment to use (default "main")
```

### dev-env rmi

```text
Remove dev-env containers and their associated images in the current directory

Usage:
  dev-env rmi [flags]

Flags:
  -a, --all       remove all images
  -c, --connect   remove connect image
  -h, --help      help for rmi
  -r, --run       remove run image
  -s, --spawn     remove spawn image

Global Flags:
  -e, --environment string   environment to use (default "main")
```

### dev-env run

```text
Run a command via sh in a run dev-env container in the current directory

Usage:
  dev-env run [flags]

Flags:
  -d, --detached   run command in the background
  -h, --help       help for run

Global Flags:
  -e, --environment string   environment to use (default "main")
```

### dev-env spawn

```text
Start a spawn dev-env container in the current directory

Usage:
  dev-env spawn [flags]

Flags:
  -h, --help   help for spawn

Global Flags:
  -e, --environment string   environment to use (default "main")
```

### dev-env stop

```text
Stop a spawned dev-env container in the current directory

Usage:
  dev-env stop [flags]

Flags:
  -h, --help   help for stop

Global Flags:
  -e, --environment string   environment to use (default "main")
```

## Configuration

Dev-env expects to find a `.dev-env` directory in the current working directory
or a parent directory containing at least one sub-directory.  The name of that
sub-directory will be the name of the environment. Inside that directory must
be a `dev-env.json` file and a `Dockerfile`. This directory is where `docker
build` will be run to create the container. The `main` environment will be
created with a default `Dockerfile` if they do not exist. The resulting image
must contain functional `sh` in `$PATH` for the `run` command and a functional
`tail` command in `$PATH` for the `spawn` command.

The `dev-env.json` config file consists of a single top level object with the
following fields.

| field | type | description |
|-------|------|-------------|
| `containerArgs` | []string | A list of arguments to pass to docker create. Optional. |
| `name` | string | The name of the container. Optional. If omitted, a sanitized version of the project directory will be used. |
| `options` | map[string]string | A map of dev-env options. |

### Replacement tokens in containerArgs 

Dev-env replaces `${var}` or `$var` in each string in `containerArgs` according
to the values of the current environment variables. References to undefined
variables are replaced by the empty string.  In addition, the following
variables are replaced by dev-env itself.

| name | replacement |
|------|-------------|
| `PROJECTROOT` | The parent directory of the `.dev-env` directory |
| `DOCKERSOCK`  | The platform specific path to the local docker socket file |

### Example config file

```json
{
    "containerArgs": [
        "--cpuset-cpus", "1,3",
        "-v", "${HOME}/.ssh:/home/user/.ssh",
        "-v", "${PROJECTROOT}:/src",
        "-v", "${DOCKERSOCK}:/var/run/docker.sock"
    ],
    "name": "nervous_golick",
    "options": {
        "out-of-date-warnings": "true"
    }
}
```

### Options

* out-of-date-warnings: Boolean. Default "false". When "true", a warning
  will be issued if the dev-env container is older than the Dockerfile.

## Details

Dev-env supports three types of development environment management.

### Connect type

The following commands work together to manage a persistent interactive shell
environment.

 * `rebuild`: Build a dev-env image in the current directory
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

### Run type

The following commands work together to create an isolated environment for
running a single command that is isolated from other command invocations.

 * `rebuild --run`: Build a run dev-env image in the current directory
 * `run`: Run a command via sh in a dev-env container in the current directory
 * `rm --run`: Remove a run dev-env container in the current directory
 * `rmi --run`: Remove a run dev-env container and its associated image in the current directory

The `run` command creates an image from the given environment (if one does not
already exist) and then destroys and re-creates a single use container for the
execution of a single command. The image is assumed to contain `sh` and the
arguments passed to the `run` command will be passed to `sh -c` in the
container. If the arguments themselves contain flags then separate the entire
command with double dashes `--`. For example, `dev-env run -- ls -la /`. This
container and image can be destroyed with the `rmr` and `rmri` commands. If
`--detached` is passed to `run` then the container will be detached and control
will return to the shell. The container will exit when the command is complete.
Because the container is recreated each time, changes made in the container are
not persistent. The image and container maintained by these commands is
distinct from the one managed by the `connect` and `spawn`family of commands.

### Spawn type

The following commands work together to create a persistant environment for
running single commands.

 * `rebuild --spawn`: Build a spawn dev-env image in the current directory
 * `spawn`: Spawn a detached dev-env container in the current directory
 * `exec`: Exec a command in a spawned dev-env container in the current directory
 * `stop`: Stop a spawned dev-env container in the current directory
 * `rm --spawn`: Remove a spawn dev-env container in the current directory
 * `rmi --spawn`: Remove a spawn dev-env container and its associated image in the current directory

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
