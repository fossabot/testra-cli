package project

import "github.com/testra-tech/testra-cli/api/models"

func MapProjectToTableRow(project *models.Project) []string {
	return []string{*project.ID, *project.Name}
}
