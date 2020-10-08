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
  help        Help about any command
  init        Initialize a dev-env in the current directory
  rm          Remove a dev-env container in the current directory
  rmi         Destroy a dev-env container and image in the current directory
  rmr         Remove a run dev-env container in the current directory
  rmri        Destroy a run dev-env container and image in the current directory
  run         Run a command via bash in a dev-env container in the current directory

Flags:
  -e, --env string   environment to use (default "main")
  -h, --help         help for dev-env

Use "dev-env [command] --help" for more information about a command.
```

## Configuration

Dev-env expects to find a `.dev-env` directory in the current working directory
containing a `config.json` file and a directory with a `Dockerfile` for each
environment in the config file.  The directory, a default config file, and a
`main` environment will be created if they do not exist. The config file
consists of a single top level objects with the following fields.

| field | type | description |
|-------|------|-------------|
| version | string | The version of this config file. Must be "1" |
| envs | map string to env object | A map of environment name to environment configuration |

Each environment in the `envs` array is an object with the following fields.

| field | type | description |
|-------|------|-------------|
| containerArgs | string array | A list of arguments to pass to docker create. Each argument in this list is passed to `os.ExpandEnv` before being added to the docker command. Optional. |
| name | string| The name of the container. Required. Init will create a name for the default `main` environment. |

Dev-env expects to find a directory in `.dev-env` named the same as each key in
the `envs` field in the config file. This directory is expected to contain a
`Dockerfile` and is the directory within which `docker build` will be run to
create the container.

### Example config file

```json
{
    "envs": {
        "main": {
            "containerArgs": [
                "--cpuset-cpus", "1,3",
		"-v", "$HOME/.ssh:/home/user/.ssh"
            ],
            "name": "nervous_golick"
        }
    }
}
```


## Details

Dev-env supports two styles of development environment management.  The
`connect` command creates an image (if one does not already exist) from the
given environment with the current working directory mounted in `/src`. It sets
the working directory of the image to be `/src`. It then creates a container
from this image (if one does not already exist) and starts the container. That
container will be re-used on subsequent `connect`s until it is removed with
`rm`. The `build` command can be used to force the image to built. Any existing
containers and images will be removed. The container can be destroyed with the
`rm` command. The image can be destroyed with the `rmi` command.

The second style of development environment management is exposed by the `run`
command.  This command creates an image from the given environment (if one does
not already exist) and then destroys and re-creates a single use container for
the execution of a single command. The image is assumed to contain `bash` and
the arguments passed to the `run` command will be passed to `bash -c` in the
container. This container and image can be destroyed with the `rmr` and `rmri`
commands. If `--interactive` is passed to `run` then the container will stay in
the foreground. Otherwise it will execute in the background and the container
will exit when the command is complete.

## License
[MIT](https://choosealicense.com/licenses/mit/)
