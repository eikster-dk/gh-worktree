package cli

import (
	"errors"
	"fmt"
	"strconv"

	gh "github.com/cli/go-gh"
	"github.com/eikster-dk/gh-worktree/internal/worktree"
	"github.com/spf13/cobra"
)

func NewPr() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "pr [number] [path]",
		Short:   "Will checkout the pr into a worktree branch",
		Example: "gh worktree pr 41",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("the pr number is required")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			number, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			branch, err := findByNumber(number)
			if err != nil {
				return err
			}

			return worktree.Add(branch)
		},
	}

	return cmd
}

func findByNumber(number int64) (string, error) {
	type response struct {
		Head struct {
			Ref string
		}
	}

	repo, err := gh.CurrentRepository()
	if err != nil {
		return "", fmt.Errorf("could not get current repository: %w", err)
	}

	restApi, err := gh.RESTClient(nil)
	if err != nil {
		return "", fmt.Errorf("could not get gh rest client: %w", err)
	}

	var resp response
	err = restApi.Get(fmt.Sprintf("repos/%s/%s/pulls/%d", repo.Owner(), repo.Name(), number), &resp)
	if err != nil {
		return "", fmt.Errorf("could not get pull request information: %w", err)
	}

	return resp.Head.Ref, nil
}
