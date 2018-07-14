// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/testra-tech/testra-cli/api/client/counter"
	"github.com/testra-tech/testra-cli/api/client/execution"
	"github.com/testra-tech/testra-cli/api/client/project"
	"github.com/testra-tech/testra-cli/api/client/result"
	"github.com/testra-tech/testra-cli/api/client/scenario"
	"github.com/testra-tech/testra-cli/api/client/test_group"
	"github.com/testra-tech/testra-cli/api/client/testcase"
)

// Default testra HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "testra-api.herokuapp.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v1/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new testra HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Testra {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new testra HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Testra {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new testra client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Testra {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Testra)
	cli.Transport = transport

	cli.Counter = counter.New(transport, formats)

	cli.Execution = execution.New(transport, formats)

	cli.Project = project.New(transport, formats)

	cli.Result = result.New(transport, formats)

	cli.Scenario = scenario.New(transport, formats)

	cli.TestGroup = test_group.New(transport, formats)

	cli.Testcase = testcase.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Testra is a client for testra
type Testra struct {
	Counter *counter.Client

	Execution *execution.Client

	Project *project.Client

	Result *result.Client

	Scenario *scenario.Client

	TestGroup *test_group.Client

	Testcase *testcase.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Testra) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Counter.SetTransport(transport)

	c.Execution.SetTransport(transport)

	c.Project.SetTransport(transport)

	c.Result.SetTransport(transport)

	c.Scenario.SetTransport(transport)

	c.TestGroup.SetTransport(transport)

	c.Testcase.SetTransport(transport)

}