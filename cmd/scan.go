package cmd

import (
	"fmt"

	"github.com/mohamedlamineallal/MacosLeanStorage/internal/config"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan targets for old files",
	Long:  `Scans the configured targets and lists files that exceed the age threshold.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		processor := NewTargetProcessor(logger, cfg.IgnorePatterns, cfg.DryRun)
		allPaths, allCommands, _, totalSize, err := processor.ProcessTargets(cfg.Targets)
		if err != nil {
			return err
		}

		totalFiles := len(allPaths)

		// Display scan results
		// ... existing display logic ...
		fmt.Printf("\nSummary: Found %d files, total size: %.2f MB, %d commands scheduled\n", totalFiles, float64(totalSize)/(1024*1024), len(allCommands))

		if cfg.DryRun {
			fmt.Println("Running in DRY RUN mode. No files were deleted.")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
