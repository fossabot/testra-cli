package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"errors"
	"github.com/spf13/viper"
	"github.com/testra-tech/testra-cli/internal/handlers"
	"github.com/testra-tech/testra-cli/internal/config"
)

// executionsCmd represents the executions command
var executionsCmd = &cobra.Command{
	Use:   "executions",
	Short: "Helps to manage test executions in testra",
	Long: ``,
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
	Long: ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

// removeExecutionCmd represents the remove sub command
var removeExecutionCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes existing test execution",
	Long: `This will remove all test executions, test results, test scenarios and testcases within the project`,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

// showExecutionCmd represents the show sub command
var showExecutionCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays a single execution info for the give execution id",
	Long: ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId, _ := cmd.Flags().GetString("projectId")
		executionId, _ := cmd.Flags().GetString("executionId")
		handlers.GetExecution(projectId, executionId)
	},
}

// showExecutionCmd represents the show sub command
var resultStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Displays a single execution info for the give execution id",
	Long: ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId, _ := cmd.Flags().GetString("projectId")
		executionId, _ := cmd.Flags().GetString("executionId")

		watchMode, _ := cmd.Flags().GetBool("watch")
		handlers.ExecutionResultStats(projectId, executionId, watchMode)
	},
}

// listExecutionsCmd represents the ls sub command
var listExecutionsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all the executions",
	Long: ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId, _ := cmd.Flags().GetString("projectId")
		handlers.ListExecutions(projectId)
	},
}

func init() {
	rootCmd.AddCommand(executionsCmd)

	executionsCmd.AddCommand(createExecutionCmd)
	executionsCmd.AddCommand(removeExecutionCmd)
	executionsCmd.AddCommand(showExecutionCmd)
	executionsCmd.AddCommand(listExecutionsCmd)
	executionsCmd.AddCommand(resultStatsCmd)

	createExecutionCmd.Flags().String("projectId", viper.GetString("projectId"), "Project Id")
	createExecutionCmd.Flags().String("description", "", "Short description about the test execution")
	createExecutionCmd.Flags().String("host", "", "hostname where the tests are going to be executed")
	createExecutionCmd.Flags().String("environment", "", "Target test environment")
	createExecutionCmd.Flags().String("branch", "", "VCS (Git, SVN..) branch")
	createExecutionCmd.Flags().String("buildRef", "", "CI build reference")
	createExecutionCmd.Flags().StringArray("tags", []string {}, "Tags")
	createExecutionCmd.Flags().BoolP("interactive", "i", false, "Interactive mode")

	removeExecutionCmd.Flags().String("projectId", viper.GetString("projectId"), "Project Id")
	removeExecutionCmd.Flags().String("executionId", "", "Test Execution Id")
	removeExecutionCmd.MarkFlagRequired("executionId")

	showExecutionCmd.Flags().String("projectId", viper.GetString("projectId"), "Project Id")
	showExecutionCmd.Flags().String("executionId", "", "Test Execution Id")
	removeExecutionCmd.MarkFlagRequired("executionId")

	listExecutionsCmd.Flags().String("projectId", viper.GetString("projectId"), "Project Id")
	listExecutionsCmd.Flags().String("environment", "", "Filter text executions by Target environment")
	listExecutionsCmd.Flags().String("branch", "", "Filter text executions by VCS (git or svn) branch")
	listExecutionsCmd.Flags().String("tags", "", "Filter text executions by tags")
	listExecutionsCmd.Flags().BoolP("interactive", "i", false, "Interactive mode")

	resultStatsCmd.Flags().String("projectId", "", "Project Id")
	resultStatsCmd.Flags().String("executionId", "", "Execution Id")
	resultStatsCmd.Flags().BoolP("watch", "w",false, "Watching mode")
}
