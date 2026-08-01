package cmd

import "github.com/spf13/cobra"

// NewRootCmd makes the root command of the CLI application.
// Make a new root command for each execution. A cobra command keeps the
// flags and the arguments that it got from Execute(). If two tests use
// the same command, the data of the first test changes the result of
// the second test.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "calculator",
		Short: "A Command-Line Interface (CLI) Calculator Written in Go.",

		// Only the run function in run.go shows an error. It writes
		// "calculator: " before the error and sets the exit code.
		// Without these two flags, cobra also shows the error and the
		// usage data. Then the CLI shows the same failure two times,
		// and the long usage data hides the error. All subcommands obey
		// these flags, but the --help flag still shows the usage data.
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	// Add child commands to the root command
	rootCmd.AddCommand(
		NewEvalCmd(),
		NewVersionCmd(),
	)

	return rootCmd
}
