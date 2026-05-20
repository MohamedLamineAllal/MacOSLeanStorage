package cmd

import (
	"github.com/mohamedlamineallal/MrLeanStorage/internal/config"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command which identifies files that can be cleaned up
// without actually deleting them.
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
		verbose, _ := cmd.Flags().GetBool("verbose")
		return processor.Run(cfg.Targets, false, verbose)
	},
}

// init adds the scan command to the root command.
func init() {
	rootCmd.AddCommand(scanCmd)
}
