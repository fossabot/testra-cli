package cmd

import (
	"fmt"

		"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/testra-tech/testra-cli/internal/config"
	"github.com/testra-tech/testra-cli/internal/execution"
			)

// executionsCmd represents the executions command
var executionsCmd = &cobra.Command{
	Use:   "executions",
	Short: "Helps to manage test executions in testra",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return fmt.Errorf("invalid command specified: %s", args[0])
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		//if(args.size != 0)
		fmt.Println("Invalid command for executions")
		cmd.Help()
	},
}

// createExecutionCmd represents the create sub command
var createExecutionCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new execution within a project",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		execution.HandleCreateExecution(cmd)
	},
}

var removeExecutionCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes existing test execution",
	Long:  `This will remove all test executions, test results, test scenarios and testcases within the project`,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		execution.HandleDeleteExecution(cmd)
	},
}

var showExecutionCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays a single execution info for the give execution id",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		execution.HandleGetExecution(cmd)
	},
}

var resultStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Displays a single execution info for the give execution id",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		execution.HandleExecutionStats(cmd)
	},
}

var listExecutionsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all the executions",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		execution.HandleListExecutions(cmd)
	},
}

func init() {
	rootCmd.AddCommand(executionsCmd)

	executionsCmd.AddCommand(createExecutionCmd)
	executionsCmd.AddCommand(removeExecutionCmd)
	executionsCmd.AddCommand(showExecutionCmd)
	executionsCmd.AddCommand(listExecutionsCmd)
	executionsCmd.AddCommand(resultStatsCmd)

	createExecutionCmd.Flags().StringP(PROJECT_ID_FLAG_NAME, "p", viper.GetString("defaultProjectId"), "Project Id")
	createExecutionCmd.Flags().StringP("description", "d", EMPTY_STR, "Short description about the test execution")
	createExecutionCmd.Flags().Bool("parallel", false, "Is tests run in parallel")
	createExecutionCmd.Flags().StringP("host", "h", EMPTY_STR, "hostname where the tests are going to be executed")
	createExecutionCmd.Flags().StringP("env", "e", EMPTY_STR, "Target test environment")
	createExecutionCmd.Flags().StringP("branch", "b", EMPTY_STR, "VCS (Git, SVN..) branch")
	createExecutionCmd.Flags().String("buildRef", EMPTY_STR, "CI build reference")
	createExecutionCmd.Flags().StringP("tags", "t", EMPTY_STR, "Tags")
	createExecutionCmd.Flags().BoolP(INTERACTIVE_FLAG_NAME, "i", false, "Interactive mode")

	removeExecutionCmd.Flags().StringP(PROJECT_ID_FLAG_NAME, "p", viper.GetString("defaultProjectId"), "Project Id")
	removeExecutionCmd.Flags().StringP(EXECUTION_ID_FLAG_NAME, "e", EMPTY_STR, "Test Execution Id")
	removeExecutionCmd.MarkFlagRequired(EXECUTION_ID_FLAG_NAME)

	showExecutionCmd.Flags().StringP(PROJECT_ID_FLAG_NAME, "p", viper.GetString("defaultProjectId"), "Project Id")
	showExecutionCmd.Flags().StringP(EXECUTION_ID_FLAG_NAME, "e", EMPTY_STR, "Test Execution Id")
	removeExecutionCmd.MarkFlagRequired(EXECUTION_ID_FLAG_NAME)

	listExecutionsCmd.Flags().StringP(PROJECT_ID_FLAG_NAME, "p", EMPTY_STR, "Project Id")
	listExecutionsCmd.Flags().StringP("env", "e", EMPTY_STR, "Filter text executions by Target environment")
	listExecutionsCmd.Flags().StringP("branch", "b", EMPTY_STR, "Filter text executions by VCS (git or svn) branch")
	listExecutionsCmd.Flags().StringP("tags", "t", EMPTY_STR, "Filter text executions by tags")
	listExecutionsCmd.Flags().BoolP(INTERACTIVE_FLAG_NAME, "i", false, "Interactive mode")

	resultStatsCmd.Flags().StringP(PROJECT_ID_FLAG_NAME, "p", viper.GetString("defaultProjectId"), "Project Id")
	resultStatsCmd.Flags().StringP(EXECUTION_ID_FLAG_NAME, "e", EMPTY_STR, "Execution Id")
	resultStatsCmd.Flags().BoolP("watch", "w", false, "Watching mode")
	resultStatsCmd.MarkFlagRequired(EXECUTION_ID_FLAG_NAME)
}
