package result

import (
	"github.com/testra-tech/testra-cli/api"
	"github.com/testra-tech/testra-cli/api/client/result"

	"strings"
	"github.com/testra-tech/testra-cli/api/client/test_group"
)

func GetResults(projectId string, execId string, status string) (*result.GetResultsOK, error) {
	statusInUpper := strings.ToUpper(status)
	getResultsParams := result.NewGetResultsParams().
		WithProjectID(projectId).
		WithExecutionID(execId).
		WithStatus(&statusInUpper)

	return api.NewTestraClient().Result.GetResults(getResultsParams)
}

func GetTestGroupsInExecution(projectId string, execId string) (*test_group.GetTestGroupsInExecutionOK, error) {

	testGroupsInExecutionParams := test_group.NewGetTestGroupsInExecutionParams().
		WithProjectID(projectId).
		WithExecutionID(execId)

	return api.NewTestraClient().TestGroup.
		GetTestGroupsInExecution(testGroupsInExecutionParams)
}
