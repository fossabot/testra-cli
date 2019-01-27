package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/testra/testra-cli/internal/constants/colors"
	"github.com/testra/testra-cli/internal/constants/flags"
	"log"
	consolereader "github.com/testra/testra-cli/internal/utils/console/reader"
	"github.com/testra/testra-cli/internal/utils/urls"
)

func HandleCreate() {
	baseUrl := consolereader.ReadUserInput("Testra API Base Url")
	projectId := consolereader.ReadUserInput("Default Project Id")

	setBaseUrlConfig(baseUrl)
	setDefaultProjectId(projectId)

	CreateConfigFileIfDoesNotExists()

	if err := viper.WriteConfig(); err == nil {
		log.Println("Configs are saved")
	} else {
		log.Fatal(err)
	}
}

func HandleSet(cmd *cobra.Command) {
	baseUrl, _ := cmd.Flags().GetString(BASE_URL)
	projectId, _ := cmd.Flags().GetString(flags.PROJECT_ID)

	if len(baseUrl) > 0 {
		setBaseUrlConfig(baseUrl)
	}

	if len(projectId) > 0 {
		setDefaultProjectId(projectId)
	}

	viper.WriteConfig()
}

func setDefaultProjectId(projectId string) {
	viper.Set("defaultprojectid", projectId)
}

func setBaseUrlConfig(baseUrl string) {
	if urls.IsValidBaseUrl(baseUrl) {
		viper.Set(BASE_URL, baseUrl)
	} else {
		log.Fatalf("Invalid base url. Given: %s", baseUrl)
	}
}

func HandleList() {
	fmt.Println(colors.Magenta)
	for _, key := range viper.AllKeys() {
		if key != SCHEME && key != DOMAIN {
			fmt.Printf("%s : %s\n", key, viper.Get(key))
		}
	}
}
