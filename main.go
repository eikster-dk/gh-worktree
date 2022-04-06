package main

import (
	"fmt"
	"os"

	"github.com/eikster-dk/gh-worktree/internal/cli"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	root := cli.NewRoot()
	_, err := root.ExecuteC()

	return err
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
