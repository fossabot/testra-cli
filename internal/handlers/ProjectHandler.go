package handlers

import (
	"github.com/testra-tech/testra-cli/internal/utils"
	"github.com/testra-tech/testra-cli/api/client/project"
	"fmt"
	"github.com/testra-tech/testra-cli/api/models"
)

var headers = []string{"Id", "Name", "Description"}

func CreateProject(name string, desc string) {
	body := models.ProjectRequest{desc, &name}

	createProjectresp, err :=
		ApiClient().Project.CreateProject(project.NewCreateProjectParams().WithBody(&body))
	checkErr(err)

	fmt.Printf("%s created successfully", *createProjectresp.Payload.Name)
}

func UpdateProject(id string, name string, desc string) {
	body := models.ProjectRequest{desc, &name}

	createProjectresp, err :=
		ApiClient().Project.UpdateProject(project.NewUpdateProjectParams().WithBody(&body).WithID(id))
	checkErr(err)

	fmt.Printf("%s updated successfully", *createProjectresp.Payload.Name)
}

func DeleteProject(id string) {

	getProjectresp, err := ApiClient().Project.GetProject(project.NewGetProjectParams().WithID(id))
	checkErr(err)

	projectName := *getProjectresp.Payload.Name

	_, err = ApiClient().Project.DeleteProject(project.NewDeleteProjectParams().WithID(id))
	checkErr(err)
	fmt.Printf("%s deleted successfully", projectName)
}

func GetProject(id string) {

	resp, err := ApiClient().Project.GetProject(project.NewGetProjectParams().WithID(id))
	checkErr(err)

	var rows [][]string
	rows = append(rows, []string{*resp.Payload.ID, *resp.Payload.Name, resp.Payload.Description})

	utils.PrintTab(headers, rows)
}

func ListProjects() {

	resp, err := ApiClient().Project.GetProjects(nil)
	checkErr(err)

	if len(resp.Payload) == 0 {
		fmt.Println("No projects found")
		return
	}

	var rows [][]string

	for _, element := range resp.Payload {
		rows = append(rows, []string{*element.ID, *element.Name, element.Description})
	}

	utils.PrintTab(headers, rows)
}
