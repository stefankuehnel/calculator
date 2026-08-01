package cmd

import (
	"bytes"
	"strings"
	"testing"
)

// assertNoUsage fails the test if out contains the error data of cobra:
// the "Error: " line, the usage data, or the note about the --help flag.
// Only the run function in run.go shows an error. If cobra also shows
// the error, the CLI shows the same failure two times.
func assertNoUsage(t *testing.T, out string) {
	t.Helper()

	for _, unexpectedText := range []string{
		"Usage:",
		"Error:",
		"Run 'calculator --help' for usage.",
	} {
		if strings.Contains(out, unexpectedText) {
			t.Errorf("expected output not to contain %q; got: %q", unexpectedText, out)
		}
	}
}

// TestRootCmd tests how the root command shows an error.
// A user must see one short error line and no usage data. The usage
// data is long, and it hides the error message that the user needs.
func TestRootCmd(t *testing.T) {
	t.Run("with an unknown command", func(t *testing.T) {
		// Arrange
		rootCmd := NewRootCmd()

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetErr(out)

		rootCmd.SetArgs([]string{"invalid-command"})

		// Act
		err := rootCmd.Execute()

		// Assert
		if err == nil {
			t.Error("expected an error, got nil")
		}

		assertNoUsage(t, out.String())
	})

	t.Run("with the help flag", func(t *testing.T) {
		// Arrange
		rootCmd := NewRootCmd()

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetErr(out)

		// Cobra shows no usage data for an error. But it must continue to
		// show the help that the --help flag requests.
		rootCmd.SetArgs([]string{"--help"})

		// Act
		err := rootCmd.Execute()

		// Assert
		if err != nil {
			t.Errorf("got error %q", err)
		}

		outStr := out.String()
		expectedText := "Usage:"
		if !strings.Contains(outStr, expectedText) {
			t.Errorf("expected output to contain %q; got: %q", expectedText, outStr)
		}
	})
}

// TestSubcommands tests the output rules that all subcommands obey.
// The loop reads the subcommands from NewRootCmd(). Thus a new
// subcommand gets these tests automatically, and you do not write the
// same test again for each new command.
func TestSubcommands(t *testing.T) {
	subCmds := NewRootCmd().Commands()

	// If the list is empty, the loop below does no test. Then the test
	// passes for the wrong reason.
	if len(subCmds) == 0 {
		t.Fatal("expected at least one subcommand, got none")
	}

	for _, subCmd := range subCmds {
		t.Run(subCmd.Name(), func(t *testing.T) {
			t.Run("with the help flag", func(t *testing.T) {
				// Arrange
				rootCmd := NewRootCmd()

				out := new(bytes.Buffer)
				rootCmd.SetOut(out)
				rootCmd.SetErr(out)

				// Cobra shows no usage data for an error. But it must
				// continue to show the help that the --help flag requests.
				rootCmd.SetArgs([]string{subCmd.Name(), "--help"})

				// Act
				err := rootCmd.Execute()

				// Assert
				if err != nil {
					t.Errorf("got error %q", err)
				}

				outStr := out.String()
				expectedText := "Usage:"
				if !strings.Contains(outStr, expectedText) {
					t.Errorf("expected output to contain %q; got: %q", expectedText, outStr)
				}
			})

			t.Run("with an unknown flag", func(t *testing.T) {
				// Arrange
				rootCmd := NewRootCmd()

				out := new(bytes.Buffer)
				rootCmd.SetOut(out)
				rootCmd.SetErr(out)

				// An unknown flag makes each command fail. A command that
				// takes no arguments also fails. Thus this test works for
				// every subcommand.
				rootCmd.SetArgs([]string{subCmd.Name(), "--unknown-flag"})

				// Act
				err := rootCmd.Execute()

				// Assert
				if err == nil {
					t.Error("expected an error, got nil")
				}

				assertNoUsage(t, out.String())
			})
		})
	}
}
