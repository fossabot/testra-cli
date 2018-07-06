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
	Short:                  "This command is useful to manage projects in testra",
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

// createProjectCmd represents the create sub command
var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new project",
	Long:  ``,
	//Args: func(cmd *cobra.Command, args []string) error {
	//	return fmt.Errorf("invalid command specified: %s", args[0])
	//	//if len(args) == 1 && args[0] == "interactive" {
	//	//	return errors.New("requires at least one arg")
	//	//}
	//	//return fmt.Errorf("invalid command specified: %s", args[0])
	//},
	PreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		isInteractive, _ := cmd.Flags().GetBool("interactive")

		var name string
		var description string

		if isInteractive {
			var readLine string
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
	projectsCmd.AddCommand(removeProjectCmd)
	projectsCmd.AddCommand(showProjectCmd)
	projectsCmd.AddCommand(listProjectsCmd)

	// Persistent flags for createProjectCmd
	createProjectCmd.Flags().StringP("name", "n", "", "Project name")
	createProjectCmd.Flags().StringP("description", "d", "", "Project description")

	// Persistent flags for removeProjectCmd
	removeProjectCmd.Flags().String("id", "", "Project id")

	// Persistent flags for showProjectCmd
	showProjectCmd.Flags().String("id", "", "Project id")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//projectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
