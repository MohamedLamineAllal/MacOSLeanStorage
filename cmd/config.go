package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command which provides subcommands for
// managing the application's configuration.
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage MacosLeanStorage configuration",
}

// openConfigCmd represents the "config open" subcommand which opens the
// configuration file in the system's default text editor or application.
var openConfigCmd = &cobra.Command{
	Use:   "open",
	Short: "Open the configuration file in the default application",
	RunE: func(cmd *cobra.Command, args []string) error {
		configFile := viper.ConfigFileUsed()
		if configFile == "" {
			return fmt.Errorf("no configuration file found")
		}

		colorInfo.Print("Opening configuration file in default application: ")
		colorPath.Println(configFile)
		return exec.Command("open", configFile).Run()
	},
}

// revealConfigCmd represents the "config reveal" subcommand which opens the
// directory containing the configuration file in Finder.
var revealConfigCmd = &cobra.Command{
	Use:   "reveal",
	Short: "Reveal the configuration file in Finder",
	RunE: func(cmd *cobra.Command, args []string) error {
		configFile := viper.ConfigFileUsed()
		if configFile == "" {
			return fmt.Errorf("no configuration file found")
		}

		colorInfo.Print("Revealing configuration file in Finder: ")
		colorPath.Println(configFile)
		return exec.Command("open", "-R", configFile).Run()
	},
}

// init adds the config command and its subcommands to the root command.
func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(openConfigCmd)
	configCmd.AddCommand(revealConfigCmd)
}
