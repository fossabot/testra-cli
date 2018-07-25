package execution

import (
	"fmt"
	"github.com/testra-tech/testra-cli/internal/check"
	"github.com/testra-tech/testra-cli/internal/constants/colors"
	consolewriter "github.com/testra-tech/testra-cli/internal/utils/console/writer"
	"github.com/gosuri/uilive"
	"github.com/testra-tech/testra-cli/internal/utils/keyevents"
	"time"
	"os"
	"io"
	"github.com/testra-tech/testra-cli/internal/commons"
	"github.com/spf13/cobra"
	"github.com/testra-tech/testra-cli/internal/constants/flags"
	"github.com/testra-tech/testra-cli/api/models"
	"github.com/testra-tech/testra-cli/internal/utils/console/reader"
	"strconv"
)

var ExecutionHeaders = []string{"Id", "Start Time", "End Time", "Environment", "Branch", "Tags"}

func HandleCreateExecution(cmd *cobra.Command) {
	isInteractive, _ := cmd.Flags().GetBool(flags.INTERACTIVE)

	if isInteractive {
		handleInteractiveCreateExecution()
	} else {
		handleNonInteractiveCreateExecution(cmd)
	}
}

func createExecution(projectId string, execRequest *models.ExecutionRequest) {
	resp, err := CreateExecution(projectId, execRequest)
	check.Err(err)
	fmt.Println(*resp.Payload.ID)
}

func handleNonInteractiveCreateExecution(cmd *cobra.Command) {
	projectId := commons.GetResolvedProjectId(cmd)
	description, _ := cmd.Flags().GetString("description")
	isParallel, _ := cmd.Flags().GetBool("parallel")
	host, _ := cmd.Flags().GetString("host")
	environment, _ := cmd.Flags().GetString("env")
	branch, _ := cmd.Flags().GetString("branch")
	tags, _ := cmd.Flags().GetString("tags")
	buildRef, _ := cmd.Flags().GetString("buildRef")

	executionRequest := NewExecutionRequestBuilder().
		Branch(branch).
		BuildRef(buildRef).
		Description(description).
		Environment(environment).
		Host(host).
		IsParallel(&isParallel).
		Tags(tags).Build()

	createExecution(projectId, &executionRequest)
}

func handleInteractiveCreateExecution() {
	projectId := reader.ReadUserInput("Project Id")
	description := reader.ReadUserInput("Description")
	isParallel, _ := strconv.ParseBool(reader.ReadUserInput("Is tests run in parallel"))
	host := reader.ReadUserInput("Host IP")
	environment := reader.ReadUserInput("Environment")
	branch := reader.ReadUserInput("Branch")
	tags := reader.ReadUserInput("Tags")
	buildRef := reader.ReadUserInput("Build Reference")

	executionRequest := NewExecutionRequestBuilder().
		Branch(branch).
		BuildRef(buildRef).
		Description(description).
		Environment(environment).
		Host(host).
		IsParallel(&isParallel).
		Tags(tags).Build()

	createExecution(projectId, &executionRequest)
}

func HandleDeleteExecution(cmd *cobra.Command) {
	projectId, executionId := getProjectIdAndExecutionIdFromCmd(cmd)

	_, err := DeleteExecution(projectId, executionId)
	check.Err(err)
	fmt.Println("OK")
}

func HandleGetExecution(cmd *cobra.Command) {
	projectId, executionId := getProjectIdAndExecutionIdFromCmd(cmd)

	handleGetExecution(projectId, executionId)
}

func handleGetExecution(projectId string, executionId string) {
	resp, err := GetExecution(projectId, executionId)
	check.Err(err)
	consolewriter.DumpStruct(MapToTestExecution(projectId, resp.Payload))
}

func HandleListExecutions(cmd *cobra.Command) {
	projectId := commons.GetResolvedProjectId(cmd)

	resp, err := ListExecutions(projectId)
	check.Err(err)

	consolewriter.PrintTable(ExecutionHeaders, MapTestExecutionsToTableRows(resp.Payload))
}

func HandleExecutionStats(cmd *cobra.Command) {
	projectId, executionId := getProjectIdAndExecutionIdFromCmd(cmd)
	watchMode, _ := cmd.Flags().GetBool("watch")

	if watchMode {
		handleExecutionStatsInWatchMode(projectId, executionId)
	} else {
		handleExecutionNonInteractiveStats(projectId, executionId)
	}
}

func getProjectIdAndExecutionIdFromCmd(cmd *cobra.Command) (string, string) {
	projectId := commons.GetResolvedProjectId(cmd)
	executionId, _ := cmd.Flags().GetString(flags.EXECUTION_ID)
	return projectId, executionId
}

func handleExecutionNonInteractiveStats(projectId string, executionId string) {
	handleGetExecution(projectId, executionId)
	enrichedResultStats := getEnrichedResultStats(projectId, executionId)
	writeStatsToConsole(os.Stdout, enrichedResultStats)
}

func handleExecutionStatsInWatchMode(projectId string, executionId string) {
	handleGetExecution(projectId, executionId)

	fmt.Println(colors.Blue, "\n WATCH Mode: Results will auto update every 30 secs. Ctrl+C to quit", colors.Reset)

	liveConsoleWriter := initLiveConsoleWriter()
	writeStatsToLiveConsoleWriter(projectId, executionId, liveConsoleWriter)
}

func initLiveConsoleWriter() *uilive.Writer {
	writer := uilive.New()
	keyevents.RegisterAndExecFuncOnInterrupt(func() {
		fmt.Print(" signal received. quitting...")
		writer.Stop()
	})
	writer.Start()
	return writer
}

func writeStatsToLiveConsoleWriter(projectId string, executionId string, liveConsoleWriter *uilive.Writer) {
	for {
		enrichedResultStats := getEnrichedResultStats(projectId, executionId)
		writeStatsToConsole(liveConsoleWriter, enrichedResultStats)

		time.Sleep(time.Second * 30)
	}
}

func writeStatsToConsole(writer io.Writer, enrichedResultStats EnrichedResultStats) {
	fmt.Fprintf(writer, StatsFormat,
		colors.Cyan, enrichedResultStats.total, colors.Reset,
		colors.Green, enrichedResultStats.passed, enrichedResultStats.passPercent, colors.Reset,
		colors.Red, enrichedResultStats.failed, enrichedResultStats.failPercent, colors.Reset,
		colors.LightRed, enrichedResultStats.expectedFailures, colors.Reset,
		colors.Yellow, enrichedResultStats.others, enrichedResultStats.otherPercent, colors.Reset)
}

func getEnrichedResultStats(projectId string, execId string) EnrichedResultStats {
	resp, err := GetResultStats(projectId, execId)
	check.Err(err)

	return MapToEnrichedResultStats(resp.Payload)
}
