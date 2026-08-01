package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewEvalCmd makes the eval command.
// It makes a new eval command for each execution. A cobra command keeps
// the flags and the arguments that it got from Execute(). Thus a new
// command gives a clean state to each test.
//
// The command takes the expression as one positional argument. If the
// expression starts with a minus sign, put "--" before the expression.
// Without "--", the flag parser reads the expression as a flag:
//
//	calculator eval -- "-3 * 4"
//
// The command writes the result to cmd.OutOrStdout(). A test replaces
// this stream with a buffer and then reads the output.
func NewEvalCmd() *cobra.Command {
	evalCmd := &cobra.Command{
		Use:   "eval <expression>",
		Short: "Evaluate an arithmetic expression",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Write to cmd.OutOrStdout() and not to os.Stdout, because a
			// test replaces this stream with a buffer and reads the output
			_, err := fmt.Fprintln(cmd.OutOrStdout(), args[0])
			if err != nil {
				return err
			}

			return nil
		},
	}

	return evalCmd
}
