package cli

import "github.com/spf13/cobra"

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
	cmd.AddCommand(NewPr())

	return cmd
}
