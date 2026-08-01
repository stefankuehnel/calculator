package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

// errWriter is an io.Writer that always fails when it writes.
// A bytes.Buffer never fails. Thus the tests use errWriter to reach the
// error branch of fmt.Fprintln and to hold the test coverage at 100%.
type errWriter struct{}

func (errWriter) Write(_ []byte) (int, error) {
	return 0, errors.New("write error")
}

// TestVersionCmd tests the version command.
// Obey these rules when you test a Cobra command:
//   - Always set the arguments and the output on the root command.
//     Execute() on the root command reads the arguments and sends them
//     to the correct subcommand. If you set the arguments on a child
//     command, Cobra does not do this step.
//   - Call NewRootCmd() to make a new command for each test. A cobra
//     command keeps the data of the last execution, and that data
//     changes the result of the next test.
//   - To read the output, call cmd.SetOut() and cmd.SetErr() on the root
//     command. A test cannot read os.Stdout.
func TestVersionCmd(t *testing.T) {
	t.Run("with no args", func(t *testing.T) {
		// Arrange
		rootCmd := NewRootCmd()

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetErr(out)

		// Set the arguments on the root command, because Execute() on the
		// root command sends them to the subcommand.
		rootCmd.SetArgs([]string{"version"})

		// Act
		err := rootCmd.Execute()

		// Assert
		if err != nil {
			t.Errorf("got error %q", err)
		}

		outStr := out.String()
		expectedText := fmt.Sprintf("%s\n", version)
		if outStr != expectedText {
			t.Errorf("expected: %q; got: %q", expectedText, outStr)
		}
	})

	t.Run("with write error", func(t *testing.T) {
		// Arrange
		rootCmd := NewRootCmd()

		// This stream always fails when it writes. Thus the test starts
		// the error branch of fmt.Fprintln in the RunE function.
		rootCmd.SetOut(errWriter{})
		rootCmd.SetErr(new(bytes.Buffer))

		// Set the arguments on the root command, because Execute() on the
		// root command sends them to the subcommand.
		rootCmd.SetArgs([]string{"version"})

		// Act
		err := rootCmd.Execute()

		// Assert
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})
}
