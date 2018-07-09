package handlers

import (
	"github.com/testra-tech/testra-cli/api/models"
	"github.com/testra-tech/testra-cli/internal/utils"
	"github.com/testra-tech/testra-cli/api/client/execution"
	"github.com/gosuri/uilive"
	"fmt"
	"time"
	"strings"
	"os"
	"os/signal"
	"syscall"
	"github.com/ttacon/chalk"
	"github.com/testra-tech/testra-cli/api"
	"github.com/testra-tech/testra-cli/api/client/project"
)

var ExecutionHeaders = []string{"Id", "Start Time", "End Time", "Environment", "Branch", "Tags"}

func CreateExecution(projectId string, execRequest *models.ExecutionRequest, isInteractive bool) {
	createExecutionResp, err :=
		api.TestraClient().Execution.
			CreateExecution(execution.NewCreateExecutionParams().
			WithProjectID(projectId).WithBody(execRequest))
	checkErr(err)

	if isInteractive {
		utils.SuccessF("Execution created successfully %s", *createExecutionResp.Payload.ID)
	} else {
		fmt.Println(*createExecutionResp.Payload.ID)
	}
}

func UpdateExecution(projectId string, executionId string, execRequest *models.ExecutionRequest) {
	_, err :=
		api.TestraClient().Execution.
			UpdateExecution(execution.NewUpdateExecutionParams().
			WithProjectID(projectId).
			WithID(executionId).
			WithBody(execRequest))
	checkErr(err)

	utils.SuccessF("Execution updated successfully %s", executionId)
}

func DeleteExecution(projectId string, id string) {
	_, err := api.TestraClient().Execution.
		DeleteExecution(execution.NewDeleteExecutionParams().
		WithProjectID(projectId).
		WithID(id))
	checkErr(err)

	utils.SuccessF("Execution deleted successfully %d", id)
}

func GetExecution(projectId string, id string) {
	resp, err := api.TestraClient().Execution.
		GetExecution(execution.NewGetExecutionParams().
		WithProjectID(projectId).
		WithID(id))
	checkErr(err)

	var endTime string = ""

	if resp.Payload.EndTime != nil {
		endTime = time.Unix(*resp.Payload.StartTime/1000, 0).Format(DateTimeFormat)
	}

	projResp, pErr := api.TestraClient().Project.GetProject(project.NewGetProjectParams().WithID(projectId))
	checkErr(pErr)
	utils.DumpStruct(TestExecution{
		*projResp.Payload.Name,
		resp.Payload.Description,
		resp.Payload.IsParallel,
		time.Unix(*resp.Payload.StartTime/1000, 0).Format(DateTimeFormat),
		endTime,
		*resp.Payload.Environment,
		*resp.Payload.Branch,
		strings.Join(resp.Payload.Tags, ", "),
		*resp.Payload.Host,
		resp.Payload.BuildRef,
	})
}

func ListExecutions(projectId string) {
	resp, err := api.TestraClient().Execution.GetExecutions(execution.NewGetExecutionsParams().WithProjectID(projectId))
	checkErr(err)

	if len(resp.Payload) == 0 {
		utils.InfoF("No Executions found for project %s", projectId)
		return
	}

	var rows [][]string

	for _, execution := range resp.Payload {
		var endTime string = ""
		if execution.EndTime != nil {
			endTime = time.Unix(*execution.EndTime/1000, 0).Format("_2 Jan 2006 15:04:05")
		}

		rows = append(rows, []string{
			*execution.ID,
			time.Unix(*execution.StartTime/1000, 0).Format("_2 Jan 2006 15:04:05"),
			endTime,
			*execution.Environment,
			*execution.Branch,
			strings.Join(execution.Tags, ",")})
	}

	utils.PrintTab(ExecutionHeaders, rows)
}

type stats struct {
	total        int64
	passed       int64
	passPercent  float32
	failed       int64
	failPercent  float32
	others       int64
	otherPercent float32
}

var statsFmt = "%s Total: %d%s, %sPassed: %d (%.2f%%)%s, %sFailed: %d (%.2f%%)%s, %sOthers: %d (%.2f%%)%s\n\n" // \n is **important

func ExecutionResultStats(projectId string, execId string, watchMode bool) {
	var cyanStyle = chalk.Cyan.NewStyle()
	var greenStyle = chalk.Green.NewStyle()
	var redStyle = chalk.Red.NewStyle()
	var yellowStyle = chalk.Yellow.NewStyle()

	if watchMode {

		fmt.Println(chalk.Blue, "\n WATCH Mode: Results will auto update every 30 secs. Ctrl+C to quit", chalk.Reset)

		GetExecution(projectId, execId)

		writer := uilive.New()

		// Registers Ctrl + C event handler
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			stopWriter(writer) // flush and stop rendering
			os.Exit(1)
		}()

		writer.Start()

		for {
			s := getResultStats(projectId, execId)

			fmt.Fprintf(writer, statsFmt,
				cyanStyle, s.total, chalk.Reset,
				greenStyle, s.passed, s.passPercent, chalk.Reset,
				redStyle, s.failed, s.failPercent, chalk.Reset,
				yellowStyle, s.others, s.otherPercent, chalk.Reset)

			time.Sleep(time.Second * 30)
		}
	} else {
		s := getResultStats(projectId, execId)
		GetExecution(projectId, execId)
		fmt.Printf(statsFmt,
			cyanStyle, s.total, chalk.Reset,
			greenStyle, s.passed, s.passPercent, chalk.Reset,
			redStyle, s.failed, s.failPercent, chalk.Reset,
			yellowStyle, s.others, s.otherPercent, chalk.Reset)
	}
}

func stopWriter(writer *uilive.Writer) {
	fmt.Print(chalk.Dim, " Signal received. Quitting...", chalk.Reset)
	writer.Stop()
}

func getResultStats(projectId string, execId string) stats {
	resp, err := api.TestraClient().Execution.
		GetExecutionResultStats(execution.NewGetExecutionResultStatsParams().
		WithProjectID(projectId).
		WithID(execId))
	checkErr(err)

	passed := resp.Payload.PassedResults
	failed := resp.Payload.FailedResults
	others := resp.Payload.OtherResults
	total := passed + failed + others
	var passPercent = float32(0)
	var failPercent = float32(0)
	var otherPercent = float32(0)

	if total != 0 {
		passPercent = float32(passed) / float32(total) * 100
		failPercent = float32(failed) / float32(total) * 100
		otherPercent = float32(others) / float32(total) * 100
	}

	return stats{
		total,
		passed,
		passPercent,
		failed,
		failPercent,
		others,
		otherPercent,
	}
}

func FilterExecutions(env []string, branch []string, tags []string) {

}
