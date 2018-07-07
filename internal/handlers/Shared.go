package handlers

import (
	"github.com/go-openapi/strfmt"
	"github.com/testra-tech/testra-cli/api/client"

	apiClient "github.com/testra-tech/testra-cli/api/client"
	httpTransport "github.com/go-openapi/runtime/client"
	"log"
	"github.com/mgutz/ansi"
)

func ApiClient() *client.Testra {

	// create the transport
	transport := httpTransport.New("localhost:8080", apiClient.DefaultBasePath, []string{"http"})

	// create the API client, with the transport
	return apiClient.New(transport, strfmt.Default)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(ansi.Color(err.Error(), "red+b"))
	}
}
