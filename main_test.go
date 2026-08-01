package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("with a valid command", func(t *testing.T) {
		// main() calls os.Exit. os.Exit stops the process immediately.
		// A direct call to main() also stops the test process. Thus
		// this code starts main() in a new process. Then the test can
		// examine the exit code.
		//
		// See: https://go.dev/talks/2014/testing.slide#23
		// See: https://stackoverflow.com/a/33404435
		if os.Getenv("BE_MAIN") == "1" {
			os.Args = []string{os.Args[0], "version"}
			main()
			return
		}

		// Arrange
		cmd := exec.Command(os.Args[0], "-test.run="+t.Name())
		cmd.Env = append(os.Environ(), "BE_MAIN=1")

		// Act
		out, err := cmd.CombinedOutput()

		// Assert
		if err != nil {
			t.Errorf("expected exit code 0, got error: %v", err)
		}

		if len(out) == 0 {
			t.Error("expected output, got none")
		}
	})

	t.Run("with an invalid command", func(t *testing.T) {
		// main() calls os.Exit. os.Exit stops the process immediately.
		// A direct call to main() also stops the test process. Thus
		// this code starts main() in a new process. Then the test can
		// examine the exit code.
		//
		// See: https://go.dev/talks/2014/testing.slide#23
		// See: https://stackoverflow.com/a/33404435
		if os.Getenv("BE_MAIN") == "1" {
			os.Args = []string{os.Args[0], "invalid-command"}
			main()
			return
		}

		// Arrange
		cmd := exec.Command(os.Args[0], "-test.run="+t.Name())
		cmd.Env = append(os.Environ(), "BE_MAIN=1")

		// Act
		out, err := cmd.CombinedOutput()

		// Assert
		e, ok := err.(*exec.ExitError)
		if !ok || e.Success() {
			t.Fatalf("process ran with err %v, want exit status 1", err)
		}

		if !strings.Contains(string(out), "calculator: ") {
			t.Errorf("expected output to contain %q, got %q", "calculator: ", out)
		}
	})
}
