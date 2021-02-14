package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Manage a developement environment via docker in the current working directory.",
	Use:   "dev-env",
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		return Initialize(nil, nil)
	},
}

var buildCmd = &cobra.Command{
	RunE:  Build,
	Short: "Build a dev-env image in the current directory",
	Use:   "build",
}

var connectCmd = &cobra.Command{
	RunE:  Connect,
	Short: "Start or connect to a dev-env container in the current directory",
	Use:   "connect",
}

var execCmd = &cobra.Command{
	RunE:  Exec,
	Short: "Exec a command via bash in a spawned dev-env container in the current directory",
	Use:   "exec",
}

var initCmd = &cobra.Command{
	RunE:  Initialize,
	Short: "Initialize a dev-env in the current directory",
	Use:   "init",
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

var removeRContainerCmd = &cobra.Command{
	RunE:  RemoveRContainer,
	Short: "Remove a run dev-env container in the current directory",
	Use:   "rmr",
}

var removeRImageCmd = &cobra.Command{
	RunE:  RemoveRImage,
	Short: "Remove a run dev-env container and its associated image in the current directory",
	Use:   "rmri",
}

var removeSContainerCmd = &cobra.Command{
	RunE:  RemoveSContainer,
	Short: "Remove a spawn dev-env container in the current directory",
	Use:   "rms",
}

var removeSImageCmd = &cobra.Command{
	RunE:  RemoveSImage,
	Short: "Remove a spawn dev-env container and its associatedimage in the current directory",
	Use:   "rmsi",
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

func Execute() error {
	execCmd.PersistentFlags().BoolP("detached", "d", false, "run command in the background")
	runCmd.PersistentFlags().BoolP("detached", "d", false, "run command in the background")
	rootCmd.PersistentFlags().StringP("env", "e", "main", "environment to use")
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(removeContainerCmd)
	rootCmd.AddCommand(removeImageCmd)
	rootCmd.AddCommand(removeRContainerCmd)
	rootCmd.AddCommand(removeRImageCmd)
	rootCmd.AddCommand(removeSContainerCmd)
	rootCmd.AddCommand(removeSImageCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(spawnCmd)
	return rootCmd.Execute()
}
