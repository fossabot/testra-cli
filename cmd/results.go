package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"errors"
	"github.com/testra-tech/testra-cli/internal/handlers"
	"github.com/testra-tech/testra-cli/internal/config"
)

// resultsCmd represents the results command
var resultsCmd = &cobra.Command{
	Use:   "results",
	Short: "Get test executions results for scenarios or test cases",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return fmt.Errorf("invalid command specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("results called")
	},
}

var resultsListCmd = &cobra.Command{
	Use: "ls",
	Short: "lists test results",
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId, _ := cmd.Flags().GetString("projectId")
		executionId, _ := cmd.Flags().GetString("executionId")
		handlers.ListResults(projectId, executionId)
	},
}

func init() {
	rootCmd.AddCommand(resultsCmd)

	resultsCmd.AddCommand(resultsListCmd)

	resultsListCmd.Flags().String("projectId", "", "Project Id")
	resultsListCmd.Flags().String("executionId", "", "Execution Id")
	resultsListCmd.MarkFlagRequired("projectId")
	resultsListCmd.MarkFlagRequired("executionId")
}
