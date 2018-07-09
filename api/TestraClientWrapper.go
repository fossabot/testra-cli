package api

import (
	"github.com/go-openapi/strfmt"
	"github.com/testra-tech/testra-cli/api/client"

	apiClient "github.com/testra-tech/testra-cli/api/client"
	httpTransport "github.com/go-openapi/runtime/client"
	"github.com/spf13/viper"
)

func TestraClient() *client.Testra {

	// create the transport
	transport := httpTransport.New(
		viper.GetString("domain"), apiClient.DefaultBasePath, []string{viper.GetString("scheme")})

	// create the API client, with the transport
	return apiClient.New(transport, strfmt.Default)
}
