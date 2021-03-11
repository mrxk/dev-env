package cmd

import (
	"github.com/mrxk/dev-env/pkg/constants"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Manage a developement environment via docker in the current working directory.",
	Use:   "dev-env",
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		return Initialize(cmd, nil)
	},
}

var connectCmd = &cobra.Command{
	RunE:  Connect,
	Short: "Connect to a connect dev-env container in the current directory",
	Use:   "connect",
}

var execCmd = &cobra.Command{
	RunE:  Exec,
	Short: "Exec a command in a spawned dev-env container in the current directory",
	Use:   "exec",
}

var initCmd = &cobra.Command{
	RunE:  Initialize,
	Short: "Initialize a dev-env in the current directory",
	Use:   "init",
}

var listCmd = &cobra.Command{
	RunE:    List,
	Short:   "List containers and images in the current directory",
	Use:     "list",
	Aliases: []string{"ls"},
}

var rebuildCmd = &cobra.Command{
	RunE:  Rebuild,
	Short: "Rebuild dev-env images in the current directory",
	Use:   "rebuild",
}

var removeContainerCmd = &cobra.Command{
	RunE:  RemoveContainer,
	Short: "Remove dev-env containers in the current directory",
	Use:   "rm",
}

var removeImageCmd = &cobra.Command{
	RunE:  RemoveImage,
	Short: "Remove dev-env containers and their associated images in the current directory",
	Use:   "rmi",
}

var runCmd = &cobra.Command{
	RunE:  Run,
	Short: "Run a command via bash in a run dev-env container in the current directory",
	Use:   "run",
}

var spawnCmd = &cobra.Command{
	RunE:  Spawn,
	Short: "Start a spawn dev-env container in the current directory",
	Use:   "spawn",
}

var stopCmd = &cobra.Command{
	RunE:  Stop,
	Short: "Stop a spawned dev-env container in the current directory",
	Use:   "stop",
}

func Execute() error {
	rebuildCmd.PersistentFlags().BoolP(constants.AllOption, constants.AllShortOption, false, "rebuild all images")
	rebuildCmd.PersistentFlags().BoolP(constants.ConnectOption, constants.ConnectShortOption, false, "rebuild connect image")
	rebuildCmd.PersistentFlags().BoolP(constants.RunOption, constants.RunShortOption, false, "rebuild run image")
	rebuildCmd.PersistentFlags().BoolP(constants.SpawnOption, constants.SpawnShortOption, false, "rebuild spawn image")

	removeContainerCmd.PersistentFlags().BoolP(constants.AllOption, constants.AllShortOption, false, "remove all containers")
	removeContainerCmd.PersistentFlags().BoolP(constants.ConnectOption, constants.ConnectShortOption, false, "remove connect container")
	removeContainerCmd.PersistentFlags().BoolP(constants.RunOption, constants.RunShortOption, false, "remove run container")
	removeContainerCmd.PersistentFlags().BoolP(constants.SpawnOption, constants.SpawnShortOption, false, "remove spawn container")

	removeImageCmd.PersistentFlags().BoolP(constants.AllOption, constants.AllShortOption, false, "remove all images")
	removeImageCmd.PersistentFlags().BoolP(constants.ConnectOption, constants.ConnectShortOption, false, "remove connect image")
	removeImageCmd.PersistentFlags().BoolP(constants.RunOption, constants.RunShortOption, false, "remove run image")
	removeImageCmd.PersistentFlags().BoolP(constants.SpawnOption, constants.SpawnShortOption, false, "remove spawn image")

	execCmd.PersistentFlags().BoolP(constants.DetachedOption, constants.DetachedShortOption, false, "run command in the background")

	runCmd.PersistentFlags().BoolP(constants.DetachedOption, constants.DetachedShortOption, false, "run command in the background")

	rootCmd.PersistentFlags().StringP(constants.EnvironmentOption, constants.EnvironmentShortOption, "main", "environment to use")

	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(rebuildCmd)
	rootCmd.AddCommand(removeContainerCmd)
	rootCmd.AddCommand(removeImageCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(spawnCmd)
	rootCmd.AddCommand(stopCmd)
	return rootCmd.Execute()
}
