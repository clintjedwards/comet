package utils

import (
	"context"
	"os/exec"
	"time"
)

// ExecuteCmd wraps a context around a given command and executes it.
// dir refers to the working directory of command to be run
func ExecuteCmd(path string, args []string, env []string, dir string) ([]byte, error) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Create command
	cmd := exec.CommandContext(ctx, path, args...)
	cmd.Env = env
	cmd.Dir = dir

	// Execute command
	return cmd.CombinedOutput()
}
