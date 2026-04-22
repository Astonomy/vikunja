//go:build mage && windows

package main

import (
	"os/exec"
	"strconv"
)

// On Windows we don't need extra process-group setup for the build to compile.
func setProcessGroup(_ *exec.Cmd) {}

func killProcessGroup(cmd *exec.Cmd) error {
	if cmd.Process == nil {
		return nil
	}

	// Best-effort: kill full process tree first.
	taskkill := exec.Command("taskkill", "/PID", strconv.Itoa(cmd.Process.Pid), "/T", "/F")
	if err := taskkill.Run(); err != nil {
		// Fall back to killing only the direct process.
		if killErr := cmd.Process.Kill(); killErr != nil {
			return killErr
		}
	}

	return cmd.Wait()
}
