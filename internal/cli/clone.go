package cli

import (
	"fmt"
	"strings"

	gh "github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

func NewClone() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clone [owner/repository] [directory]",
		Short:   "Will clone a github repository into a folder",
		Example: "gh worktree clone eikster-dk/gh-worktree",
		RunE: func(cmd *cobra.Command, args []string) error {
			repo := args[0]
			if repo == "" {
				return fmt.Errorf("[owner/repository] argument is required")
			}

			var directory = ""
			ss := strings.Split(repo, "/")
			if len(ss) == 1 {
				directory = repo
			} else {
				directory = ss[1]
			}

			if len(args) > 1 {
				directory = args[1]
			}

			repoPath := fmt.Sprintf("%s/%s", directory, ".git")
			_, stdErr, err := gh.Exec("repo", "clone", repo, repoPath, "--", "--bare")
			if err != nil {
				return err
			}
			fmt.Println(stdErr.String())

			fmt.Println("repository has been cloned and ready for git worktree")
			return nil
		},
	}

	return cmd
}
