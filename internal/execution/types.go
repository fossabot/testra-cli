package execution

type TestExecution struct {
	ProjectName string
	Description string
	IsParallel  *bool
	StartTime   string
	EndTime     string
	Environment string
	Branch      string
	tags        string
	Host        string
	BuildRef    string
}

type EnrichedResultStats struct {
	total            int64
	passed           int64
	passPercent      float32
	failed           int64
	failPercent      float32
	expectedFailures int64
	others           int64
	otherPercent     float32
}