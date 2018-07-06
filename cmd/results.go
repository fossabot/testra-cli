package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"errors"
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

var listCmd = &cobra.Command{
	Use: "ls",
	Short: "lists test results",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("results called")
	},
}

func init() {
	rootCmd.AddCommand(resultsCmd)


}
