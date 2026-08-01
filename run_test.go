package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stefankuehnel/calculator/cmd"
)

// TestRun tests the exit code of run and how run writes errors.
func TestRun(t *testing.T) {
	t.Run("with a valid command", func(t *testing.T) {
		// Arrange
		rootCmd := cmd.NewRootCmd()

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetErr(out)
		rootCmd.SetArgs([]string{"version"})

		// Act
		exitCode := run(rootCmd)

		// Assert
		if exitCode != 0 {
			t.Errorf("expected exit code 0, got %d", exitCode)
		}

		if out.Len() == 0 {
			t.Error("expected output, got none")
		}
	})

	t.Run("with an invalid command", func(t *testing.T) {
		// Arrange
		rootCmd := cmd.NewRootCmd()

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetErr(out)
		rootCmd.SetArgs([]string{"invalid-command"})

		// Act
		exitCode := run(rootCmd)

		// Assert
		if exitCode != 1 {
			t.Errorf("expected exit code 1, got %d", exitCode)
		}

		errStr := out.String()
		if !strings.Contains(errStr, "calculator: ") {
			t.Errorf("expected output to contain %q, got %q", "calculator: ", errStr)
		}
	})
}
