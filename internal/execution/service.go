package execution

import (
	"github.com/testra/testra-cli/api"
	"github.com/testra/testra-cli/api/client/execution"
	"github.com/testra/testra-cli/api/models"
)

func CreateExecution(projectId string,
	execRequest *models.ExecutionRequest) (*execution.CreateExecutionCreated, error) {

	createExecutionParams := execution.NewCreateExecutionParams().
		WithProjectID(projectId).WithBody(execRequest)

	return api.NewTestraClient().Execution.CreateExecution(createExecutionParams)
}

func UpdateExecution(projectId string,
	executionId string,
	execRequest *models.ExecutionRequest) (*execution.UpdateExecutionOK, error) {

	updateExecutionParams := execution.NewUpdateExecutionParams().
		WithProjectID(projectId).
		WithID(executionId).
		WithBody(execRequest)

	return api.NewTestraClient().Execution.UpdateExecution(updateExecutionParams)
}

func DeleteExecution(projectId string, id string) (*execution.DeleteExecutionOK, error) {
	deleteExecutionParams := execution.NewDeleteExecutionParams().
		WithProjectID(projectId).
		WithID(id)

	return api.NewTestraClient().Execution.
		DeleteExecution(deleteExecutionParams)
}

func GetExecution(projectId string, id string) (*execution.GetExecutionOK, error) {
	getExecutionParams := execution.NewGetExecutionParams().
		WithProjectID(projectId).
		WithID(id)

	return api.NewTestraClient().Execution.
		GetExecution(getExecutionParams)
}

func ListExecutions(projectId string) (*execution.GetExecutionsOK, error) {
	getExecutionsParams := execution.NewGetExecutionsParams().WithProjectID(projectId)

	return api.NewTestraClient().Execution.GetExecutions(getExecutionsParams)
}

func GetResultStats(projectId string, execId string) (*execution.GetExecutionResultStatsOK, error) {
	getExecutionResultStatsParams := execution.NewGetExecutionResultStatsParams().
		WithProjectID(projectId).
		WithID(execId)

	return api.NewTestraClient().Execution.
		GetExecutionResultStats(getExecutionResultStatsParams)
}
