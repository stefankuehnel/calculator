package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// version is the version of the application.
// It is a variable and not a constant, because the Go linker can change
// only a variable. Set the version at build time with ldflags:
//
//	go build -ldflags "-X github.com/stefankuehnel/calculator/cmd.version=v1.2.3"
//
// If you do not set the variable at build time, it stays "dev".
var version = "dev"

// NewVersionCmd makes the version command.
// It makes a new version command for each execution. A Cobra command
// keeps the flags and the arguments that it got from Execute(). Thus a
// new command gives a clean state to each test.
//
// The command writes the version to cmd.OutOrStdout(). A test replaces
// this stream with a buffer and then reads the output.
func NewVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Write to cmd.OutOrStdout() and not to os.Stdout, because a
			// test replaces this stream with a buffer and reads the output.
			_, err := fmt.Fprintln(cmd.OutOrStdout(), version)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return versionCmd
}
