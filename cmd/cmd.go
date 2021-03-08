package cmd

import (
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
	Short: "Start or connect to a dev-env container in the current directory",
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
	Short: "Rebuild a dev-env image in the current directory",
	Use:   "rebuild",
}

var removeContainerCmd = &cobra.Command{
	RunE:  RemoveContainer,
	Short: "Remove a dev-env container in the current directory",
	Use:   "rm",
}

var removeImageCmd = &cobra.Command{
	RunE:  RemoveImage,
	Short: "Remove a dev-env container and its associated image in the current directory",
	Use:   "rmi",
}

var runCmd = &cobra.Command{
	RunE:  Run,
	Short: "Run a command via bash in a dev-env container in the current directory",
	Use:   "run",
}

var spawnCmd = &cobra.Command{
	RunE:  Spawn,
	Short: "Spawn a detached dev-env container in the current directory",
	Use:   "spawn",
}

var stopCmd = &cobra.Command{
	RunE:  Stop,
	Short: "Stop a spawned dev-env container in the current directory",
	Use:   "stop",
}

func Execute() error {
	rebuildCmd.PersistentFlags().StringSliceP("type", "t", []string{"connect"}, "type of image to rebuild [all, connect, run, spawn]")
	removeContainerCmd.PersistentFlags().StringSliceP("type", "t", []string{"connect"}, "type of image to rebuild [all, connect, run, spawn]")
	removeImageCmd.PersistentFlags().StringSliceP("type", "t", []string{"connect"}, "type of image to rebuild [all, connect, run, spawn]")
	execCmd.PersistentFlags().BoolP("detached", "d", false, "run command in the background")
	runCmd.PersistentFlags().BoolP("detached", "d", false, "run command in the background")
	rootCmd.PersistentFlags().StringP("environment", "e", "main", "environment to use")
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
