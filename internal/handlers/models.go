package handlers

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

//type TestScenarioResult struct {
//	Name string,
//	Status string,
//	Duration int32,
//
//}