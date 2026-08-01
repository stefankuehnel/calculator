package cmd

import (
	"bytes"
	"testing"
)

// TestEvalCmd tests the eval command.
// Obey these rules when you test a Cobra command:
//   - Always set the arguments and the output on the root command.
//     Execute() on the root command reads the arguments and sends them
//     to the correct subcommand. If you set the arguments on a child
//     command, Cobra does not do this step.
//   - Call NewRootCmd() to make a new command for each test. A Cobra
//     command keeps the data of the last execution, and that data
//     changes the result of the next test.
//   - To read the output, call cmd.SetOut() and cmd.SetErr() on the root
//     command. A test cannot read os.Stdout.
func TestEvalCmd(t *testing.T) {
	t.Run("with an expression", func(t *testing.T) {
		// Arrange
		rootCmd := NewRootCmd()

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetErr(out)

		// Set the arguments on the root command, because Execute() on the
		// root command sends them to the subcommand.
		rootCmd.SetArgs([]string{"eval", "1 + 2"})

		// Act
		err := rootCmd.Execute()

		// Assert
		if err != nil {
			t.Errorf("got error %q", err)
		}

		outStr := out.String()
		expectedText := "3\n"
		if outStr != expectedText {
			t.Errorf("expected: %q; got: %q", expectedText, outStr)
		}
	})

	t.Run("with an expression starting with a minus sign", func(t *testing.T) {
		// Arrange
		rootCmd := NewRootCmd()

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetErr(out)

		// Set the arguments on the root command, because Execute() on the
		// root command sends them to the subcommand.
		rootCmd.SetArgs([]string{"eval", "--", "-3 * 4"})

		// Act
		err := rootCmd.Execute()

		// Assert
		if err != nil {
			t.Errorf("got error %q", err)
		}

		outStr := out.String()
		expectedText := "-12\n"
		if outStr != expectedText {
			t.Errorf("expected: %q; got: %q", expectedText, outStr)
		}
	})

	t.Run("with an invalid expression", func(t *testing.T) {
		// Arrange
		rootCmd := NewRootCmd()

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetErr(out)

		rootCmd.SetArgs([]string{"eval", "1 +"})

		// Act
		err := rootCmd.Execute()

		// Assert
		if err == nil {
			t.Error("expected an error, got nil")
		}

		// A mistake in the expression is not a mistake in the command.
		// Thus the CLI must show no usage data.
		assertNoUsage(t, out.String())
	})

	t.Run("with a division by zero", func(t *testing.T) {
		// Arrange
		rootCmd := NewRootCmd()

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetErr(out)

		rootCmd.SetArgs([]string{"eval", "1 / 0"})

		// Act
		err := rootCmd.Execute()

		// Assert
		if err == nil {
			t.Error("expected an error, got nil")
		}

		assertNoUsage(t, out.String())
	})

	t.Run("with no args", func(t *testing.T) {
		// Arrange
		rootCmd := NewRootCmd()

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetErr(out)

		// Give no expression. ExactArgs(1) must reject this and give an error.
		rootCmd.SetArgs([]string{"eval"})

		// Act
		err := rootCmd.Execute()

		// Assert
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("with too many args", func(t *testing.T) {
		// Arrange
		rootCmd := NewRootCmd()

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetErr(out)

		// Give two expressions. ExactArgs(1) must reject this and give an error.
		rootCmd.SetArgs([]string{"eval", "1 + 2", "3 + 4"})

		// Act
		err := rootCmd.Execute()

		// Assert
		if err == nil {
			t.Error("expected an error, got nil")
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
		rootCmd.SetArgs([]string{"eval", "1 + 2"})

		// Act
		err := rootCmd.Execute()

		// Assert
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})
}
