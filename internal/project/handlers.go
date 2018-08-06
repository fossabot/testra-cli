package project

import (
	"github.com/testra-tech/testra-cli/internal/check"
	consolewriter "github.com/testra-tech/testra-cli/internal/utils/console/writer"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/testra-tech/testra-cli/internal/utils/console/reader"
)

var projectHeaders = []string{"Id", "Name"}

func HandleCreate(cmd *cobra.Command) {
	name, description := getProjectNameAndDescriptionFromCmdFlags(cmd)

	createProject(name, description)
}

func HandleInteractiveCreate(cmd *cobra.Command) {
	name, description := getProjectNameAndDescFromUserInput()

	createProject(name, description)
}

func createProject(name string, description string) {
	resp, err := CreateProject(name, description)
	check.Err(err)
	fmt.Println(*resp.Payload.ID)
}

func HandleUpdate(cmd *cobra.Command) {
	name, description := getProjectNameAndDescriptionFromCmdFlags(cmd)

	updateProject(getProjectIdFromCmdFlag(cmd), name, description)
}

func HandleInteractiveUpdate(cmd *cobra.Command) {
	projectId := reader.ReadUserInput("Project Id : ")
	name, description := getProjectNameAndDescFromUserInput()

	updateProject(projectId, name, description)
}

func getProjectNameAndDescriptionFromCmdFlags(cmd *cobra.Command) (string, string) {
	name, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")
	return name, description
}

func getProjectNameAndDescFromUserInput() (string, string) {
	name := reader.ReadUserInput("Project Name : ")
	description := reader.ReadUserInput("Project Description : ")
	return name, description
}

func updateProject(projectId string, name string, description string) {
	resp, err := UpdateProject(projectId, name, description)
	check.Err(err)
	fmt.Println(*resp.Payload.ID)
}

func HandleDelete(cmd *cobra.Command) {
	_, err := DeleteProject(getProjectIdFromCmdFlag(cmd))
	check.Err(err)
	fmt.Println("OK")
}

func HandleGet(cmd *cobra.Command) {
	projectIdFromCmdFlag := getProjectIdFromCmdFlag(cmd)
	resp, err := GetProject(projectIdFromCmdFlag)
	check.Err(err)
	consolewriter.DumpStruct(*resp.Payload)
}

func getProjectIdFromCmdFlag(cmd *cobra.Command) string {
	projectId, _ := cmd.Flags().GetString("projectId")
	return projectId
}

func HandleList() {
	resp, err := ListProjects()
	check.Err(err)

	var tableRows [][]string
	for _, element := range resp.Payload {
		tableRows = append(tableRows, MapProjectToTableRow(element))
	}

	consolewriter.PrintTable(projectHeaders, tableRows)
}
