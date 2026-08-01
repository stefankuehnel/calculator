package main

import (
	"os"

	"github.com/stefankuehnel/calculator/cmd"
)

func main() {
	os.Exit(run(cmd.NewRootCmd()))
}
