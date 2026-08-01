package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/stefankuehnel/calculator/internal/calculator"
)

// NewEvalCmd makes the eval command.
// It makes a new eval command for each execution. A Cobra command keeps
// the flags and the arguments that it got from Execute(). Thus a new
// command gives a clean state to each test.
//
// The command takes the expression as one positional argument. The
// expression obeys the grammar in grammar/Calculator.g4. If the
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
			result, err := calculator.Evaluate(args[0])
			if err != nil {
				// Give the error to Cobra. The run function in run.go
				// writes "calculator: " before the error and sets the
				// exit code.
				return err
			}

			// The format 'g' with the precision -1 gives the shortest
			// text that reads back as the same float64. Thus a result
			// without a fraction shows as "7" and not as "7.000000".
			text := strconv.FormatFloat(result, 'g', -1, 64)

			// Write to cmd.OutOrStdout() and not to os.Stdout, because a
			// test replaces this stream with a buffer and reads the output.
			_, err = fmt.Fprintln(cmd.OutOrStdout(), text)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return evalCmd
}
