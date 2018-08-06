package cmd

import (
	"github.com/spf13/cobra"
	"github.com/testra-tech/testra-cli/internal/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Initialise and manage testra configuration",
	Long:  ``,
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitViperGlobalConfigs()
	},
	Run: func(cmd *cobra.Command, args []string) {
		config.HandleCreate()
	},
}

var setSubCmd = &cobra.Command{
	Use:   "set",
	Short: "Set config using flags",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		config.HandleSet(cmd)
	},
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all config values",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		config.HandleList()
	},
}

func init() {
	RootCmd.AddCommand(configCmd)

	configCmd.AddCommand(setSubCmd)
	configCmd.AddCommand(lsCmd)

	setSubCmd.Flags().StringP(BASE_URL_FLAG_NAME, "u", EMPTY_STR, "Testra API base url")
	setSubCmd.Flags().StringP(PROJECT_ID_FLAG_NAME, "p", EMPTY_STR, "Testra project Id")
}
