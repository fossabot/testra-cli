package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"errors"
	"github.com/spf13/viper"
	"github.com/testra-tech/testra-cli/internal/handlers"
	"github.com/testra-tech/testra-cli/internal/config"
	"bufio"
	"os"
	"strings"
	"github.com/testra-tech/testra-cli/api/models"
	"strconv"
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
		isInteractive, _ := cmd.Flags().GetBool(INTERACTIVE_FLAG_NAME)

		var projectId string
		var description string
		var isParallel bool
		var host string
		var environment string
		var branch string
		var tags string
		var buildRef string

		if isInteractive {
			var readLine string
			fmt.Println()
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Project Id : ")
			readLine, _ = reader.ReadString('\n')
			projectId = strings.TrimSuffix(readLine, "\n")

			fmt.Print("Execution Description : ")
			readLine, _ = reader.ReadString('\n')
			description = strings.TrimSuffix(readLine, "\n")

			fmt.Print("Is Parallel execution? (true/false) : ")
			readLine, _ = reader.ReadString('\n')
			isParallel, _ = strconv.ParseBool(strings.TrimSuffix(readLine, "\n"))

			fmt.Print("Environment : ")
			readLine, _ = reader.ReadString('\n')
			environment = strings.TrimSuffix(readLine, "\n")

			fmt.Print("Branch : ")
			readLine, _ = reader.ReadString('\n')
			branch = strings.TrimSuffix(readLine, "\n")

			fmt.Print("Tags : ")
			readLine, _ = reader.ReadString('\n')
			tags = strings.TrimSuffix(readLine, "\n")

			fmt.Print("Host : ")
			readLine, _ = reader.ReadString('\n')
			host = strings.TrimSuffix(readLine, "\n")

			fmt.Print("Build Reference : ")
			readLine, _ = reader.ReadString('\n')
			buildRef = strings.TrimSuffix(readLine, "\n")

		} else {
			projectId = GetResolvedProjectId(cmd)
			description, _ = cmd.Flags().GetString("description")
			isParallel, _ = cmd.Flags().GetBool("parallel")
			host, _ = cmd.Flags().GetString("host")
			environment, _ = cmd.Flags().GetString("env")
			branch, _ = cmd.Flags().GetString("branch")
			tags, _ = cmd.Flags().GetString("tags")
			buildRef, _ = cmd.Flags().GetString("buildRef")
		}

		executionRequest := models.ExecutionRequest{
			branch,
			buildRef,
			description,
			0,
			environment,
			&host,
			&isParallel,
			strings.Split(tags, ","),
		}

		handlers.CreateExecution(projectId, &executionRequest, isInteractive)
	},
}

// removeExecutionCmd represents the remove sub command
var removeExecutionCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes existing test execution",
	Long:  `This will remove all test executions, test results, test scenarios and testcases within the project`,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := GetResolvedProjectId(cmd)
		executionId, _ := cmd.Flags().GetString(EXECUTION_ID_FLAG_NAME)
		handlers.DeleteExecution(projectId, executionId)
	},
}

// showExecutionCmd represents the show sub command
var showExecutionCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays a single execution info for the give execution id",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := GetResolvedProjectId(cmd)
		executionId, _ := cmd.Flags().GetString(EXECUTION_ID_FLAG_NAME)
		handlers.GetExecution(projectId, executionId)
	},
}

// showExecutionCmd represents the show sub command
var resultStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Displays a single execution info for the give execution id",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := GetResolvedProjectId(cmd)
		executionId, _ := cmd.Flags().GetString(EXECUTION_ID_FLAG_NAME)

		watchMode, _ := cmd.Flags().GetBool("watch")
		handlers.ExecutionResultStats(projectId, executionId, watchMode)
	},
}

// listExecutionsCmd represents the ls sub command
var listExecutionsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all the executions",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId := GetResolvedProjectId(cmd)

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
