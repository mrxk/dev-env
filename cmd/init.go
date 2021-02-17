package cmd

import (
	"github.com/mrxk/dev-env/pkg/config"
	"github.com/mrxk/dev-env/pkg/constants"
	"github.com/spf13/cobra"
)

func Initialize(_ *cobra.Command, _ []string) error {
	err := writeDefaultConfigFile()
	if err != nil {
		return err
	}
	return writeDefaultDockerFile()
}

func writeDefaultConfigFile() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	return cfg.WriteIfNotExist()
}

func writeDefaultDockerFile() error {
	return config.WriteConfigFileIfNotExist(
		constants.DefaultEnvironment,
		config.DockerFile,
		[]byte(constants.DefaultDockerFile))
}
