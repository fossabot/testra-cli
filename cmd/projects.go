package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/mgutz/ansi"
	"errors"
	"github.com/testra-tech/testra-cli/internal/config"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "This command is useful to manage projects in testra",
	Long: ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return fmt.Errorf("invalid command specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		if(len(args) != 0) {
			fmt.Println(ansi.Color("Invalid command for projects", "red"))
		}
		cmd.Help()
	},
}
// createProjectCmd represents the create sub command
var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new project",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
		fmt.Println("create called")
	},
}


// removeProjectCmd represents the remove sub command
var removeProjectCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes existing project from testra",
	Long: `This will remove all test executions, test results, test scenarios and testcases within the project`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
		fmt.Println("remove called")
	},
}

// showProjectCmd represents the show sub command
var showProjectCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays a single project info for the give project id",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
		fmt.Println("show called")
	},
}

// listProjectsCmd represents the ls sub command
var listProjectsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all the projects",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
		fmt.Println("list called")
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
	createProjectCmd.Flags().String("name", "", "Project name")
	createProjectCmd.Flags().String("description", "", "Project description")
	createProjectCmd.MarkFlagRequired("name")
	createProjectCmd.MarkFlagRequired("description")

	// Persistent flags for removeProjectCmd
	removeProjectCmd.Flags().String("id", "", "Project id")

	// Persistent flags for showProjectCmd
	showProjectCmd.Flags().String("id", "", "Project id")


	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//projectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


type geometry interface {
	area() float64
	perim() float64
}