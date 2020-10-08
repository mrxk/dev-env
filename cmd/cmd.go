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

var initCmd = &cobra.Command{
	RunE:  Initialize,
	Short: "Initialize a dev-env in the current directory",
	Use:   "init",
}

var removeContainer = &cobra.Command{
	RunE:  RemoveContainer,
	Short: "Remove a dev-env container in the current directory",
	Use:   "rm",
}

var removeImage = &cobra.Command{
	RunE:  RemoveImage,
	Short: "Destroy a dev-env container and image in the current directory",
	Use:   "rmi",
}

var removeRContainer = &cobra.Command{
	RunE:  RemoveRContainer,
	Short: "Remove a run dev-env container in the current directory",
	Use:   "rmr",
}

var removeRImage = &cobra.Command{
	RunE:  RemoveRImage,
	Short: "Destroy a run dev-env container and image in the current directory",
	Use:   "rmri",
}

var runCmd = &cobra.Command{
	RunE:  Run,
	Short: "Run a command via bash in a dev-env container in the current directory",
	Use:   "run",
}

func Execute() error {
	runCmd.PersistentFlags().BoolP("interactive", "i", false, "run command interactively")
	rootCmd.PersistentFlags().StringP("env", "e", "main", "environment to use")
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(removeContainer)
	rootCmd.AddCommand(removeImage)
	rootCmd.AddCommand(removeRContainer)
	rootCmd.AddCommand(removeRImage)
	rootCmd.AddCommand(runCmd)
	return rootCmd.Execute()
}
