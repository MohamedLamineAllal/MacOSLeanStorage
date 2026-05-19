package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
}

var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload the configuration for the running background agent",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Find the PID of the running mls serve process
		pid, err := findMLSServePID()
		if err != nil {
			return fmt.Errorf("could not find running mls serve process: %w", err)
		}

		// Send SIGHUP to the process
		process, err := os.FindProcess(pid)
		if err != nil {
			return err
		}
		err = process.Signal(os.Interrupt) // SIGHUP is not directly supported on all platforms via os.Signal, but we can use syscall
		// Using syscall for SIGHUP
		return process.Signal(syscall.SIGHUP)
	},
}

func findMLSServePID() (int, error) {
	out, err := exec.Command("pgrep", "-f", "mls serve").Output()
	if err != nil {
		return 0, err
	}
	pidStr := strings.TrimSpace(string(out))
	return strconv.Atoi(pidStr)
}

func init() {
	configCmd.AddCommand(reloadCmd)
	rootCmd.AddCommand(configCmd)
}
