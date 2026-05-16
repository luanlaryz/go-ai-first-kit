package main

import (
	"fmt"
	"os"

	"github.com/inventa-co/go-ai-first-kit/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		if coded, ok := err.(interface{ ExitCode() int }); ok {
			os.Exit(coded.ExitCode())
		}
		os.Exit(2)
	}
}
