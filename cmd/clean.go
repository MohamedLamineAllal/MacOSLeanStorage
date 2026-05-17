package cmd

import (
	"fmt"

	"github.com/mohamedlamineallal/MacosLeanStorage/internal/config"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Scan and clean old files",
	Long:  `Scans the configured targets and deletes files that exceed the age threshold.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		processor := NewTargetProcessor(logger, cfg.IgnorePatterns, cfg.DryRun)
		allPaths, allCommands, commandNames, _, err := processor.ProcessTargets(cfg.Targets)
		if err != nil {
			return err
		}

		if len(allPaths) == 0 && len(allCommands) == 0 {
			fmt.Println("No files or commands found to clean.")
			return nil
		}

		if len(allPaths) > 0 {
			fmt.Printf("Cleaning %d files...\n", len(allPaths))
			count, size, err := processor.cleaner.Clean(allPaths)
			if err != nil {
				return err
			}
			fmt.Printf("Clean Summary: Deleted %d files, freed %.2f MB\n", count, float64(size)/(1024*1024))
		}

		for i, cmd := range allCommands {
			err := processor.cleaner.ExecuteCommand(cmd)
			if err == nil {
				processor.scheduler.UpdateCommandRunTime(commandNames[i])
			}
		}

		fmt.Printf("Mode: %s\n", map[bool]string{true: "DRY RUN", false: "LIVE"}[cfg.DryRun])
		return nil
	},

}

func init() {
rootCmd.AddCommand(cleanCmd)
}

