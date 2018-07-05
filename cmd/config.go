package cmd

import (
	"github.com/spf13/cobra"
	"github.com/testra-tech/testra-cli/internal/config"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Initialise or Update testra configurations",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		config.HandleConfig()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
