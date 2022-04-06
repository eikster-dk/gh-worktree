package worktree

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/cli/safeexec"
)

func Add(branch string) error {
	cmdArgs := []string{"worktree", "add", branch}

	path, err := getCommonGitDirectory()
	if err != nil {
		return fmt.Errorf("could not get working directory: %w", err)
	}

	_, err = git(cmdArgs, path)
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
	c.Dir = directory

	return c.Output()
}
