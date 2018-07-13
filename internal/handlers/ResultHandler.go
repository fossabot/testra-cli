package handlers

import (
	"github.com/testra-tech/testra-cli/api"
	"github.com/testra-tech/testra-cli/api/client/result"

	"fmt"
	"log"
	"bytes"
	"strings"
	"os"
	"github.com/testra-tech/testra-cli/internal/utils"
	"github.com/testra-tech/testra-cli/api/models"
	"github.com/testra-tech/testra-cli/internal/colors"
	"os/exec"
)

const (
	DoubleSpaces = " "
)

func ListResults(projectId string, execId string, status string) {
	utils.StartSpin()
	defer utils.StopSpin()
	statusInUpper := strings.ToUpper(status)
	resp, err :=
		api.TestraClient().Result.
			GetResults(result.NewGetResultsParams().
			WithProjectID(projectId).
			WithExecutionID(execId).
			WithStatus(&statusInUpper))
	checkErr(err)

	var buffer bytes.Buffer

	for _, result := range resp.Payload {
		buffer.WriteString("\n")

		fmt.Fprintln(&buffer, colors.CYAN, "Scenario: ", *result.Scenario.Name, colors.RESET)

		allSteps := append(result.Scenario.BackgroundSteps, result.Scenario.Steps...)

		for _, step := range allSteps {
			resultStatus, err := getStepResult(step.Index, result)

			if err != nil {
				utils.Warn("Step result not found for [" + *step.Text + "] in scenario [" + *result.Scenario.Name + "]")
				continue
			}

			switch *resultStatus {
			case "PASSED":
				fmt.Fprintln(&buffer, DoubleSpaces, colors.GREEN, *step.Text, colors.RESET)
			case "FAILED":
				fmt.Fprintln(&buffer, DoubleSpaces, colors.RED, *step.Text, colors.RESET)
				fmt.Fprintln(&buffer, DoubleSpaces, "\t" + result.Error)
			default:
				fmt.Fprintln(&buffer, DoubleSpaces, colors.YELLOW, *step.Text, colors.YELLOW)
			}
		}
	}

	utils.StopSpin()
	lessIt(buffer.String())
}

func getStepResult(index *int64, result *models.EnrichedTestResult) (*string, error) {
	for _, stepResult := range result.StepResults {
		if *stepResult.Index == *index {
			return stepResult.Status, nil
		}
	}
	return nil, fmt.Errorf("Step result not found")
}

func lessIt(s string) {
	// Could read $PAGER rather than hardcoding the path.less
	cmd := exec.Command("/usr/bin/less", []string{"-R"}...)

	// Feed it with the string you want to display.
	cmd.Stdin = strings.NewReader(s)

	// This is crucial - otherwise it will write to a null device.
	cmd.Stdout = os.Stdout

	// Fork off a process and wait for it to terminate.
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
