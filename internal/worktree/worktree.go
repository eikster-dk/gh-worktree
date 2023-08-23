package worktree

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/cli/safeexec"
)

func Add(branch string, path string) error {
    var branchPath string
    if path != "" {
        branchPath = filepath.Join(path, branch)
    } else {
        gitPath, err := getCommonGitDirectory()
        if err != nil {
            return fmt.Errorf("could not get working directory: %w", err)
        }

        branchPath = filepath.Join(gitPath, branch)
    }

    cmdArgs := []string{"worktree", "add", branchPath}

    _, err := git(cmdArgs, "")
	return err
}

func getCommonGitDirectory() (string, error) {
	args := []string{"rev-parse", "--git-common-dir"}
	b, err := git(args, "")
	if err != nil {
		return "", fmt.Errorf("could not get git common dir: %w", err)
	}

	root := filepath.Join(string(b), "..")

	return root, nil
}

func git(args []string, directory string) ([]byte, error) {
	cmd, err := safeexec.LookPath("git")
	if err != nil {
		return nil, err
	}
	c := exec.Command(cmd, args...)

	return c.Output()
}
