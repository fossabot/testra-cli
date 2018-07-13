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
	Use:   "ls",
	Short: "lists test results",
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := GetResolvedProjectId(cmd)
		executionId, _ := cmd.Flags().GetString(EXECUTION_ID_FLAG_NAME)
		status, _ := cmd.Flags().GetString("status")

		handlers.ListResults(projectId, executionId, status)
	},
}

func init() {
	rootCmd.AddCommand(resultsCmd)

	resultsCmd.AddCommand(resultsListCmd)

	resultsListCmd.Flags().
		StringP(PROJECT_ID_FLAG_NAME, "p", EMPTY_STR, "Project Id")
	resultsListCmd.Flags().
		StringP(EXECUTION_ID_FLAG_NAME, "e", EMPTY_STR, "Execution Id")
	resultsListCmd.Flags().
		StringP("status", "s", EMPTY_STR, "Filter results on status. [passed, failed, skipped, pending, ambiguous, undefined ]")
	resultsListCmd.MarkFlagRequired(EXECUTION_ID_FLAG_NAME)
}
