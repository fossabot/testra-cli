package project

import "github.com/testra/testra-cli/api/models"

func MapProjectToTableRow(project *models.Project) []string {
	return []string{*project.ID, *project.Name}
}
