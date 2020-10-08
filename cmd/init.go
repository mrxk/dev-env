package cmd

import (
	"github.com/mrxk/dev-env/pkg/config"
	"github.com/spf13/cobra"
)

const (
	DefaultDockerFile = `
FROM debian:stretch
ENTRYPOINT ["/bin/bash"]
`
)

func Initialize(_ *cobra.Command, _ []string) error {
	err := writeDefaultConfigFile()
	if err != nil {
		return err
	}
	err = writeDefaultDockerFile()
	if err != nil {
		return err
	}
	return nil
}

func writeDefaultConfigFile() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	err = cfg.WriteIfNotExist()
	if err != nil {
		return err
	}
	return nil
}

func writeDefaultDockerFile() error {
	err := config.WriteConfigFileIfNotExist("main", config.DockerFile, []byte(DefaultDockerFile))
	if err != nil {
		return err
	}
	return nil
}
