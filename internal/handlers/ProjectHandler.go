package handlers

import (
	"github.com/testra-tech/testra-cli/internal/utils"
	"github.com/testra-tech/testra-cli/api/client/project"
	"github.com/testra-tech/testra-cli/api/models"
	"github.com/testra-tech/testra-cli/api"

)

var projectHeaders = []string{"Id", "Name"}

func CreateProject(name string, desc string) {
	body := models.ProjectRequest{desc, &name}

	resp, err :=
		api.TestraClient().Project.CreateProject(project.NewCreateProjectParams().WithBody(&body))
	checkErr(err)

	utils.SuccessF("%s created successfully", *resp.Payload.Name)
}

func UpdateProject(id string, name string, desc string) {
	body := models.ProjectRequest{desc, &name}

	resp, err :=
		api.TestraClient().Project.UpdateProject(project.NewUpdateProjectParams().WithBody(&body).WithID(id))
	checkErr(err)

	utils.SuccessF("%s updated successfully", *resp.Payload.Name)
}

func DeleteProject(id string) {

	resp, err := api.TestraClient().Project.GetProject(project.NewGetProjectParams().WithID(id))
	checkErr(err)

	projectName := *resp.Payload.Name

	_, err = api.TestraClient().Project.DeleteProject(project.NewDeleteProjectParams().WithID(id))
	checkErr(err)

	utils.SuccessF("%s deleted successfully", projectName)
}

func GetProject(id string) {

	resp, err := api.TestraClient().Project.GetProject(project.NewGetProjectParams().WithID(id))
	checkErr(err)

	utils.DumpStruct(*resp.Payload)
}

func ListProjects() {

	resp, err := api.TestraClient().Project.GetProjects(nil)
	checkErr(err)

	if len(resp.Payload) == ZERO {
		utils.Info("No projects found")
		return
	}

	var rows [][]string

	for _, element := range resp.Payload {
		rows = append(rows, []string{*element.ID, *element.Name})
	}

	utils.PrintTab(projectHeaders, rows)
}
