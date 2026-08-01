package main

import (
	"log"

	"github.com/spf13/cobra"
)

// run starts rootCmd and returns an exit code for the process.
// It writes each error to rootCmd.ErrOrStderr() and not directly to
// os.Stderr. A test cannot examine os.Stderr. The function that calls
// run must also call os.Exit. os.Exit stops the process immediately. If
// run called os.Exit, a test could not examine the exit code or the
// error text.
func run(rootCmd *cobra.Command) int {
	if err := rootCmd.Execute(); err != nil {
		logger := log.Default()
		logger.SetPrefix("calculator: ")
		logger.SetOutput(rootCmd.ErrOrStderr())
		logger.SetFlags(0)
		logger.Println(err)

		return 1
	}

	return 0
}
