package cmd

import (
	"github.com/spf13/cobra"
	"github.com/testra-tech/testra-cli/internal/config"
	"github.com/testra-tech/testra-cli/internal/utils"
	"os"
	"github.com/spf13/viper"
	"fmt"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Initialise or update testra configurations",
	Long: ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.SetConfigFileAbsPath()
	},
	Run: func(cmd *cobra.Command, args []string) {
		config.HandleConfig()
	},
}

var setSubCmd = &cobra.Command{
	Use:   "set",
	Short: "Set config using flags",
	Long: ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.SetConfigFileAbsPath()
	},
	Run: func(cmd *cobra.Command, args []string) {
		baseUrl, _ := cmd.Flags().GetString("baseUrl")

		isValid, _ := config.IsValidBaseUrl(baseUrl)

		if isValid {
			config.WriteConfigToFile(baseUrl)
			os.Exit(0)
		} else {
			utils.DangerF("Invalid base url. Given: %s", baseUrl)
			os.Exit(1)
		}
	},
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all config values",
	Long: ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		for _, key := range viper.AllKeys() {
			fmt.Printf("%s : %s\n", key, viper.Get(key))
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(setSubCmd)
	configCmd.AddCommand(lsCmd)

	setSubCmd.Flags().StringP("baseUrl", "u", "", "Testra API base url")
	setSubCmd.MarkFlagRequired("baseUrl")
}
