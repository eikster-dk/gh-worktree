package main

import (
	"fmt"
	"os"

	gh "github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	root := NewRoot()
	_, err := root.ExecuteC()

	return err
}

func NewClone() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clone [owner/repository]",
		Short:   "Will clone a github repository into a folder",
		Example: "gh worktree clone eikster-dk/gh-worktree",
		RunE: func(cmd *cobra.Command, args []string) error {
			repo := args[0]
			if repo == "" {
				return fmt.Errorf("[owner/repository] argument is required")
			}

			stdOut, _, err := gh.Exec("repo", "clone", repo, ".repo.git", "--", "--bare")
			if err != nil {
				return err
			}
			fmt.Println(stdOut.String())

			err = os.WriteFile(".git", []byte("gitdir: ./.repo.git"), 0644)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}

func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "worktree <command> <subcommand> [flags]",
		Short: "github extension to ease the use of working with worktree and gh cli",
		Long:  `Work seamlessly across git worktree and gh cli tooling`,

		SilenceErrors: true,
		SilenceUsage:  false,
		Example:       `gh worktree`,
	}

	cmd.AddCommand(NewClone())

	return cmd
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
