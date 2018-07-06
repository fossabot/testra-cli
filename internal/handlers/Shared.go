package handlers

import (
	"github.com/go-openapi/strfmt"
	"github.com/testra-tech/testra-cli/api/client"

	apiClient "github.com/testra-tech/testra-cli/api/client"
	httpTransport "github.com/go-openapi/runtime/client"
	"log"
)

func ApiClient() *client.Testra {

	// create the transport
	transport := httpTransport.New("localhost:8080", apiClient.DefaultBasePath, []string{"http"})

	// create the API client, with the transport
	client := apiClient.New(transport, strfmt.Default)

	return client
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
