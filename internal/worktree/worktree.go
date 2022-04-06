package worktree

import (
	"os/exec"

	"github.com/cli/safeexec"
)

func Add(branch string) error {
	cmdArgs := []string{"worktree", "add", branch}

	return git(cmdArgs)
}

func git(args []string) error {
	cmd, err := safeexec.LookPath("git")
	if err != nil {
		return err
	}
	c := exec.Command(cmd, args...)

	return c.Run()
}
