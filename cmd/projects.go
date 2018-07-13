package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/testra-tech/testra-cli/internal/config"
	"errors"
	"github.com/testra-tech/testra-cli/internal/handlers"
	"bufio"
	"os"
	"strings"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:                    "projects",
	Short:                  "Manage projects in Testra",
	Long:                   ``,
	BashCompletionFunction: "projects",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires at least one arg")
		}
		return fmt.Errorf("Invalid command specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// createProjectCmd represents the create sub command
var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new project",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		isInteractive, _ := cmd.Flags().GetBool(INTERACTIVE_FLAG_NAME)

		var name string
		var description string

		if isInteractive {
			var readLine string
			fmt.Println()
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Project Name : ")
			readLine, _ = reader.ReadString('\n')
			name = strings.TrimSuffix(readLine, "\n")

			fmt.Print("Project Description : ")
			readLine, _ = reader.ReadString('\n')
			description = strings.TrimSuffix(readLine, "\n")
		} else {
			name, _ = cmd.Flags().GetString("name")
			description, _ = cmd.Flags().GetString("description")
		}

		handlers.CreateProject(name, description)
	},
}

// createProjectCmd represents the create sub command
var updateProjectCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates existing project",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		isInteractive, _ := cmd.Flags().GetBool(INTERACTIVE_FLAG_NAME)

		var projectId string
		var name string
		var description string

		if isInteractive {
			var readLine string
			fmt.Println()
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Project Id : ")
			readLine, _ = reader.ReadString('\n')
			projectId = strings.TrimSuffix(readLine, "\n")

			fmt.Print("Project Name : ")
			readLine, _ = reader.ReadString('\n')
			name = strings.TrimSuffix(readLine, "\n")

			fmt.Print("Project Description : ")
			readLine, _ = reader.ReadString('\n')
			description = strings.TrimSuffix(readLine, "\n")
		} else {
			projectId, _ = cmd.Flags().GetString("id")
			name, _ = cmd.Flags().GetString("name")
			description, _ = cmd.Flags().GetString("description")
		}

		handlers.UpdateProject(projectId, name, description)
	},
}

// removeProjectCmd represents the remove sub command
var removeProjectCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes existing project from testra",
	Long:  `This will remove all test executions, test results, test scenarios and testcases within the project`,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectId, _ := cmd.Flags().GetString("id")
		handlers.DeleteProject(projectId)
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
		projectId, _ := cmd.Flags().GetString("id")
		handlers.GetProject(projectId)
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
		handlers.ListProjects()
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
