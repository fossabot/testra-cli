package result

import (
	"bytes"
	"github.com/testra-tech/testra-cli/api/models"
	"github.com/testra-tech/testra-cli/internal/check"
	"github.com/testra-tech/testra-cli/internal/utils/syscmd"
	"github.com/testra-tech/testra-cli/api/client/result"
	"github.com/spf13/cobra"
	"github.com/testra-tech/testra-cli/internal/commons"
	"github.com/testra-tech/testra-cli/internal/constants/flags"
	"github.com/testra-tech/testra-cli/internal/utils/spinner"
)

func HandleList(cmd *cobra.Command) {
	projectId := commons.GetResolvedProjectId(cmd)
	executionId, _ := cmd.Flags().GetString(flags.EXECUTION_ID)
	status, _ := cmd.Flags().GetString("status")

	var formattedResultsBuffer bytes.Buffer

	getResultsAndFormat := func() {
		resp, err := GetResults(projectId, executionId, status)
		check.Err(err)

		formattedResultsBuffer.Write(getFormattedScenarioResults(resp).Bytes())
		formattedResultsBuffer.Write(getFormattedTestCaseResults(projectId, executionId, resp).Bytes())
	}

	spinner.LoadWithSpinner(getResultsAndFormat)

	syscmd.LessIt(formattedResultsBuffer.String())
}

func getFormattedScenarioResults(resp *result.GetResultsOK) *bytes.Buffer {
	scenarioResults := getScenarioResults(resp)
	return MapScenarioResultsToFormattedScenarioResults(scenarioResults)
}

func getFormattedTestCaseResults(projectId string, executionId string, resp *result.GetResultsOK) *bytes.Buffer {
	testCaseResults := getTestCaseResults(resp)

	testGroups, err := GetTestGroupsInExecution(projectId, executionId)
	check.Err(err)

	return MapTestCaseResultsToFormattedTestCaseResults(testCaseResults, testGroups.Payload)
}

func getScenarioResults(resp *result.GetResultsOK) []*models.EnrichedTestResult {
	var scenarioResults []*models.EnrichedTestResult
	for _, result := range resp.Payload {
		if *result.ResultType == "SCENARIO" {
			scenarioResults = append(scenarioResults, result)
		}
	}
	return scenarioResults
}

func getTestCaseResults(resp *result.GetResultsOK) []*models.EnrichedTestResult {
	var scenarioResults []*models.EnrichedTestResult
	for _, result := range resp.Payload {
		if *result.ResultType == "TEST_CASE" {
			scenarioResults = append(scenarioResults, result)
		}
	}
	return scenarioResults
}
