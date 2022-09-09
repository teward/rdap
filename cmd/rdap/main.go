package main

import (
	"os"

	"github.com/teward/rdap"
)

func main() {
	exitCode := rdap.RunCLI(os.Args[1:], os.Stdout, os.Stderr, rdap.CLIOptions{})

	os.Exit(exitCode)
}
