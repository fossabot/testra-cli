package cmd

import (
	"fmt"

		"errors"
	"github.com/spf13/cobra"
	"github.com/testra-tech/testra-cli/internal/config"
	"github.com/testra-tech/testra-cli/internal/project"
		)

var projectsCmd = &cobra.Command{
	Use:                    "projects",
	Short:                  "Manage projects in Testra",
	Long:                   ``,
	BashCompletionFunction: "projects",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return fmt.Errorf("invalid command specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new project",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		isInteractive, _ := cmd.Flags().GetBool(INTERACTIVE_FLAG_NAME)

		if isInteractive {
			project.HandleInteractiveCreate(cmd)
		} else {
			project.HandleCreate(cmd)
		}
	},
}

var updateProjectCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates existing project",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		isInteractive, _ := cmd.Flags().GetBool(INTERACTIVE_FLAG_NAME)

		if isInteractive {
			project.HandleInteractiveUpdate(cmd)
		} else {
			project.HandleUpdate(cmd)
		}
	},
}

var removeProjectCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes existing project from testra",
	Long:  `This will remove all test executions, test results, test scenarios and testcases within the project`,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		project.HandleDelete(cmd)
	},
}

// showProjectCmd represents the show sub command
var showProjectCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays a single project info for the give project id",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		project.HandleGet(cmd)
	},
}

// listProjectsCmd represents the ls sub command
var listProjectsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all the projects",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		project.HandleList()
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)

	// Sub commands in projects
	projectsCmd.AddCommand(createProjectCmd)
	projectsCmd.AddCommand(updateProjectCmd)
	projectsCmd.AddCommand(removeProjectCmd)
	projectsCmd.AddCommand(showProjectCmd)
	projectsCmd.AddCommand(listProjectsCmd)

	// Flags for createProjectCmd
	createProjectCmd.Flags().StringP("name", "n", EMPTY_STR, "Project name")
	createProjectCmd.Flags().StringP("description", "d", EMPTY_STR, "Project description")
	createProjectCmd.Flags().BoolP(INTERACTIVE_FLAG_NAME, "i", false, "Interactive mode")

	// Flags for createProjectCmd
	updateProjectCmd.Flags().StringP(PROJECT_ID_FLAG_NAME, "p", EMPTY_STR, "Project id")
	updateProjectCmd.Flags().StringP("name", "n", EMPTY_STR, "Project name")
	updateProjectCmd.Flags().StringP("description", "d", EMPTY_STR, "Project description")
	updateProjectCmd.Flags().BoolP(INTERACTIVE_FLAG_NAME, "i", false, "Interactive mode")

	// Flags for removeProjectCmd
	removeProjectCmd.Flags().StringP(PROJECT_ID_FLAG_NAME, "p", EMPTY_STR, "Project id")
	removeProjectCmd.MarkFlagRequired(PROJECT_ID_FLAG_NAME)

	// Flags for showProjectCmd
	showProjectCmd.Flags().StringP(PROJECT_ID_FLAG_NAME, "p", EMPTY_STR, "Project id")
	showProjectCmd.MarkFlagRequired(PROJECT_ID_FLAG_NAME)
}
