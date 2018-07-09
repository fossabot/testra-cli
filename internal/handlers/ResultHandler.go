package handlers

import (
	"github.com/testra-tech/testra-cli/api"
	"github.com/testra-tech/testra-cli/api/client/result"

	"github.com/ttacon/chalk"
	"fmt"
	"os/exec"
	"log"
	"bytes"
	"strings"
	"os"
	"github.com/testra-tech/testra-cli/internal/utils"
	"github.com/testra-tech/testra-cli/api/models"
)

const (
	DoubleSpaces = " "
)

var greenStyle = chalk.Green.NewStyle()
var redStyle = chalk.Red.NewStyle()
var yellowStyle = chalk.Yellow.NewStyle()

/*func UpdateResult(id string, result models.TestResult) {

}*/

func ListResults(projectId string, execId string) {
	utils.StartSpin()
	defer utils.StopSpin()
	resp, err :=
		api.TestraClient().Result.
			GetResults(result.NewGetResultsParams().
			WithProjectID(projectId).
			WithExecutionID(execId))
	checkErr(err)

	var buffer bytes.Buffer

	for _, result := range resp.Payload {
		buffer.WriteString("\n")

		fmt.Fprintln(&buffer, chalk.Cyan, "Scenario: ", *result.Scenario.Name, chalk.Reset)

		allSteps := append(result.Scenario.BackgroundSteps, result.Scenario.Steps...)

		for _, step := range allSteps {
			resultStatus, err := getStepResult(step.Index, result)

			if err != nil {
				utils.Warn("Step result not found for [" + *step.Text + "] in scenario [" + *result.Scenario.Name + "]")
				continue
			}

			switch *resultStatus {
			case "PASSED":
				fmt.Fprintln(&buffer, DoubleSpaces, chalk.Green, *step.Text, chalk.Reset)
			case "FAILED":
				fmt.Fprintln(&buffer, DoubleSpaces, chalk.Red, *step.Text, chalk.Reset)
			default:
				fmt.Fprintln(&buffer, DoubleSpaces, chalk.Yellow, *step.Text, chalk.Reset)
			}
		}
	}

	utils.StopSpin()
	lessIt(buffer.String())
}

func getStepResult(index *int64, result *models.EnrichedTestResult) (*string, error) {
	for _, stepResult := range  result.StepResults {
		if *stepResult.Index == *index {
			return stepResult.Result, nil
		}
	}
	return nil, fmt.Errorf("Step result not found")
}

func lessIt(s string) {
	// Could read $PAGER rather than hardcoding the path.less
	cmd := exec.Command("/usr/bin/less")

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
