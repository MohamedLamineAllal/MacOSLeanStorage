package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"github.com/mohamedlamineallal/MrLeanStorage/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open the configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := config.GetDefaultConfigPath()
		if err != nil {
			return err
		}
		return exec.Command("open", path).Run()
	},
}

var revealCmd = &cobra.Command{
	Use:   "reveal",
	Short: "Reveal the configuration file in Finder",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := config.GetDefaultConfigPath()
		if err != nil {
			return err
		}
		return exec.Command("open", "-R", path).Run()
	},
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
	configCmd.AddCommand(openCmd)
	configCmd.AddCommand(revealCmd)
	rootCmd.AddCommand(configCmd)
}
