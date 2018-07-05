package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"errors"
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

// removeExecutionCmd represents the remove sub command
var removeExecutionCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes existing test execution",
	Long: `This will remove all test executions, test results, test scenarios and testcases within the project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

// showExecutionCmd represents the show sub command
var showExecutionCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays a single execution info for the give execution id",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
	},
}

// listExecutionCmd represents the ls sub command
var listExecutionCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all the executions",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}


func init() {
	rootCmd.AddCommand(executionsCmd)

}
