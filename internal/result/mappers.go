package result

import (
	"github.com/testra-tech/testra-cli/api/models"
	"github.com/testra-tech/testra-cli/internal/constants/strs"
	"bytes"
	"github.com/testra-tech/testra-cli/internal/constants/colors"
	"github.com/pkg/errors"
	strUtils "github.com/testra-tech/testra-cli/internal/utils/strs"
	"strings"
)

func MapScenarioResultsToFormattedScenarioResults(scenarioResults []*models.EnrichedTestResult) *bytes.Buffer {
	var buffer bytes.Buffer
	for _, result := range scenarioResults {
		tags := strUtils.ColoriseText(colors.Cyan, mapScenarioTagsToSpaceDelimitedTags(result.Scenario.Tags))
		buffer.WriteString(strs.DoubleSpaces + tags)

		scenario := strUtils.ColoriseText(colors.Magenta, prefixWithScenarioKeyword(*result.Scenario.Name))
		buffer.WriteString(strs.NewLine + strs.DoubleSpaces + scenario)

		allSteps := append(result.Scenario.BackgroundSteps, result.Scenario.Steps...)
		stepsToStepResults := mapStepsToStepResults(allSteps, result.StepResults)
		buffer.WriteString((&stepsToStepResults).String())

		buffer.WriteString(strs.NewLine + strs.NewLine)
	}
	return &buffer
}

func mapStepsToStepResults(steps []*models.TestStep, stepResults []*models.StepResult) bytes.Buffer {
	var buffer bytes.Buffer
	for _, testStep := range steps {
		stepResult, _ := getStepResult(testStep.Index, stepResults)

		stepResultsWithStatusColors := mapStepResultsWithStatusColors(*stepResult, appendDuration(*testStep.Text, *stepResult.DurationInMs))
		buffer.WriteString((&stepResultsWithStatusColors).String())
	}
	return buffer
}

func appendDuration(stepTest string, duration int64) string {
	return stepTest + strs.DoubleSpaces + strs.OpenCircleBrak + "" + strs.CloseCircleBrak
}

func mapStepResultsWithStatusColors(stepResult models.StepResult, step string) bytes.Buffer {
	var buffer bytes.Buffer
	buffer.WriteString(strs.NewLine + strs.FourSpaces)

	switch *stepResult.Status {
	case "PASSED":
		buffer.WriteString(strUtils.ColoriseText(colors.Green, step))
	case "FAILED":
		buffer.WriteString(strUtils.ColoriseText(colors.Red, step))
		buffer.WriteString(strs.NewLine + strs.Tab + strings.TrimSpace(stepResult.Error))
	default:
		buffer.WriteString(strUtils.ColoriseText(colors.Yellow, step))
	}
	return buffer
}

func getStepResult(index *int64, result []*models.StepResult) (*models.StepResult, error) {
	for _, stepResult := range result {
		if *stepResult.Index == *index {
			return stepResult, nil
		}
	}
	return nil, errors.New("Step result not found")
}

func mapScenarioTagsToSpaceDelimitedTags(tags []string) string {
	spaceDelimitedTags := strs.EmptyStr
	for _, tag := range tags {
		spaceDelimitedTags = spaceDelimitedTags + tag + " "
	}
	return spaceDelimitedTags
}

func prefixWithScenarioKeyword(name string) string {
	return "Scenario: " + name
}

func MapTestCaseResultsToFormattedTestCaseResults(testCaseResults []*models.EnrichedTestResult,
	testGroups []*models.TestGroup) *bytes.Buffer {
	var buffer bytes.Buffer

	for _, testGroup := range testGroups {
		testGroupWithTestCaseResults := enrichTestGroupWithTestCaseResults(testGroup, testCaseResults)
		buffer.Write((&testGroupWithTestCaseResults).Bytes())
	}

	return &buffer
}

func enrichTestGroupWithTestCaseResults(testGroup *models.TestGroup, testCaseResults []*models.EnrichedTestResult) bytes.Buffer {
	var buffer bytes.Buffer
	testCaseResultsByGroupId := getTestCaseResultsByGroupId(*testGroup.ID, testCaseResults)
	if len(testCaseResultsByGroupId) != 0 {
		buffer.WriteString(strUtils.ColoriseText(colors.Magenta, *testGroup.Name + "." + *testGroup.SubGroup))
		for _, testCaseResult := range testCaseResultsByGroupId {
			buffer.WriteString(strs.NewLine + strs.DoubleSpaces + strs.TreeNode)
			buffer.WriteString(applyStatusColorToTestCaseName(testCaseResult))
		}
		buffer.WriteString(strs.NewLine + strs.NewLine)
	}
	return buffer
}

func applyStatusColorToTestCaseName(testCaseResult *models.EnrichedTestResult) string {
	var testCaseNameWithStatusColor string
	switch *testCaseResult.Status {
	case "PASSED":
		testCaseNameWithStatusColor = strUtils.ColoriseText(colors.Green, *testCaseResult.Testcase.Name)
	case "FAILED":
		testCaseNameWithStatusColor = strUtils.ColoriseText(colors.Red, *testCaseResult.Testcase.Name)
		testCaseNameWithStatusColor = testCaseNameWithStatusColor + strs.NewLine + strs.FourSpaces + testCaseResult.Error
	default:
		testCaseNameWithStatusColor = strUtils.ColoriseText(colors.Yellow, *testCaseResult.Testcase.Name)
	}
	return testCaseNameWithStatusColor
}

func getTestCaseResultsByGroupId(groupId string,
	testCaseResults []*models.EnrichedTestResult) []*models.EnrichedTestResult {
	var matchedCaseResults []*models.EnrichedTestResult

	for _, testCaseResult := range testCaseResults {
		if groupId == *testCaseResult.GroupID {
			matchedCaseResults = append(matchedCaseResults, testCaseResult)
		}
	}

	return matchedCaseResults
}
