package project

import (
	"github.com/testra-tech/testra-cli/api/models"
	"github.com/testra-tech/testra-cli/api"
	"github.com/testra-tech/testra-cli/api/client/project"
	"github.com/testra-tech/testra-cli/internal/check"
)

func CreateProject(name string, desc string) (*project.CreateProjectCreated, error) {
	body := models.ProjectRequest{Description: desc, Name: &name}
	createProjectParams := project.
		NewCreateProjectParams().
		WithBody(&body)
	return projectClient().CreateProject(createProjectParams)
}

func UpdateProject(id string, name string, desc string) (*project.UpdateProjectOK, error) {
	body := models.ProjectRequest{Description: desc, Name: &name}
	updateProjectParams := project.
		NewUpdateProjectParams().
		WithBody(&body).
		WithID(id)
	return projectClient().UpdateProject(updateProjectParams)
}

func DeleteProject(id string) (*project.DeleteProjectOK, error) {
	deleteProjectParams := project.NewDeleteProjectParams().WithID(id)
	return projectClient().DeleteProject(deleteProjectParams)
}

func GetProject(id string) (*project.GetProjectOK, error) {
	getProjectParams := project.
		NewGetProjectParams().
		WithID(id)
	return api.NewTestraClient().Project.GetProject(getProjectParams)
}

func GetProjectName(projectId string) string {
	projectResp, pErr := GetProject(projectId)
	check.Err(pErr)
	return *projectResp.Payload.Name
}

func ListProjects() (*project.GetProjectsOK, error) {
	return projectClient().GetProjects(nil)
}

func projectClient() *project.Client {
	return api.NewTestraClient().Project
}
