package execution

import (
	"github.com/testra/testra-cli/internal/project"
	"github.com/testra/testra-cli/internal/utils/datetime"
	"strings"
	"github.com/testra/testra-cli/api/models"
)

func MapToTestExecution(projectId string, execution *models.Execution) TestExecution {
	return TestExecution{
		ProjectName: project.GetProjectName(projectId),
		Description: execution.Description,
		IsParallel:  execution.Parallel,
		StartTime:   datetime.ConvertToDateTime(execution.StartTime),
		EndTime:     datetime.ConvertToDateTime(execution.EndTime),
		Environment: *execution.Environment,
		Branch:      *execution.Branch,
		tags:        strings.Join(execution.Tags, ", "),
		Host:        *execution.Host,
		BuildRef:    execution.BuildRef,
	}
}

func MapTestExecutionsToTableRows(executions []*models.Execution) [][]string {
	var tableRows [][]string
	for _, execution := range executions {
		tableRows = append(tableRows, mapTestExecutionToTableRow(execution))
	}
	return tableRows
}

func mapTestExecutionToTableRow(execution *models.Execution) []string {
	return []string{
		*execution.ID,
		datetime.ConvertToDateTime(execution.StartTime),
		datetime.ConvertToDateTime(execution.EndTime),
		*execution.Environment,
		*execution.Branch,
		strings.Join(execution.Tags, ",")}
}

func MapToEnrichedResultStats(executionStats *models.TestExecutionStats) (EnrichedResultStats) {
	noOfPassedTests := executionStats.Passed
	noOfFailedTests := executionStats.Failed
	noOfExpectedFailureTests := executionStats.ExpectedFailures
	otherStatusTests := executionStats.Others

	total := sum(noOfPassedTests, noOfFailedTests, otherStatusTests)

	return EnrichedResultStats{
		total:            total,
		passed:           noOfPassedTests,
		passPercent:      percentage(noOfPassedTests, total),
		failed:           noOfFailedTests,
		failPercent:      percentage(noOfFailedTests, total),
		expectedFailures: noOfExpectedFailureTests,
		others:           otherStatusTests,
		otherPercent:     percentage(otherStatusTests, total),
	}
}

func sum(values ... int64) int64 {
	var sum int64 = 0
	for _, val := range values {
		sum = sum + val
	}
	return sum
}

func percentage(value int64, total int64) float32 {
	if total <= 0 {
		return 0
	} else {
		return float32(value) / float32(total) * 100
	}
}
