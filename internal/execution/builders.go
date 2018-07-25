package execution

import (
	"github.com/testra-tech/testra-cli/api/models"
	"time"
	"strings"
)

type ExecutionRequestBuilder interface {
	Description(string) ExecutionRequestBuilder
	BuildRef(string) ExecutionRequestBuilder
	IsParallel(*bool) ExecutionRequestBuilder
	Environment(string) ExecutionRequestBuilder
	Branch(string) ExecutionRequestBuilder
	Tags(string) ExecutionRequestBuilder
	Host(string) ExecutionRequestBuilder
	Build() models.ExecutionRequest
}

type executionRequestBuilder struct {
	description string
	isParallel  *bool
	environment string
	branch      string
	tags        string
	host        string
	buildRef    string
}

func (erb *executionRequestBuilder) Branch(branch string) ExecutionRequestBuilder {
	erb.branch = branch
	return erb
}

func (erb *executionRequestBuilder) Description(desc string) ExecutionRequestBuilder {
	erb.description = desc
	return erb
}

func (erb *executionRequestBuilder) BuildRef(buildRef string) ExecutionRequestBuilder {
	erb.buildRef = buildRef
	return erb
}

func (erb *executionRequestBuilder) IsParallel(isParallel *bool) ExecutionRequestBuilder {
	erb.isParallel = isParallel
	return erb
}

func (erb *executionRequestBuilder) Environment(environment string) ExecutionRequestBuilder {
	erb.environment = environment
	return erb
}

func (erb *executionRequestBuilder) Tags(tags string) ExecutionRequestBuilder {
	erb.tags = tags
	return erb
}

func (erb *executionRequestBuilder) Host(host string) ExecutionRequestBuilder {
	erb.host = host
	return erb
}

func (erb *executionRequestBuilder) Build() models.ExecutionRequest {
	return models.ExecutionRequest{
		Branch:      erb.branch,
		BuildRef:    erb.buildRef,
		Description: erb.description,
		EndTime:     time.Now().UnixNano() / int64(time.Millisecond),
		Environment: erb.environment,
		Host:        &erb.host,
		Parallel:    erb.isParallel,
		Tags:        strings.Split(erb.tags, ","),
	}
}

func NewExecutionRequestBuilder() ExecutionRequestBuilder {
	return &executionRequestBuilder{}
}