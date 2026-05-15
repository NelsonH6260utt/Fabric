package main

import (
	"fmt"
	"os"

	"github.com/danielmiessler/fabric/cli"
)

// main is the entry point for the Fabric application.
// Fabric is an open-source framework for augmenting humans using AI.
// It provides a modular system for applying AI patterns to everyday tasks.
//
// Personal fork notes:
//   - Using this for experimenting with custom AI patterns and local LLM workflows.
//   - See the patterns/ directory for any custom patterns I've added.
func main() {
	if err := cli.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
