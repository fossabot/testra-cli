package config

import (
	"github.com/sanity-io/litter"
	"github.com/spf13/viper"
	"github.com/testra/testra-cli/internal/constants/colors"
	"github.com/testra/testra-cli/internal/constants/strs"
		"log"
	"os"
	"github.com/testra/testra-cli/internal/utils/urls"
	"github.com/testra/testra-cli/internal/utils/paths"
)

func InitConfig() {

	InitViperGlobalConfigs()
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(colors.Red + "Missing config for testra. Run command `testra config` to initialise")
	}

	scheme, domain, _ := urls.ExtractSchemeAndDomainWithPortFromUrl(viper.GetString(BASE_URL))
	viper.Set(SCHEME, *scheme)
	viper.Set(DOMAIN, *domain)

	initLitterConfig()
}

func initLitterConfig() {
	// Strip all package names from types
	litter.Config.StripPackageNames = true

	// Sets separator used when multiple arguments are passed to Dump() or Sdump().
	litter.Config.Separator = strs.NewLine
}

func CreateConfigFileIfDoesNotExists() {
	if !doesConfigFileExist() {
		os.Create(getConfigFileAbsPath())
	}
}

func doesConfigFileExist() bool {
	_, err := os.Stat(getConfigFileAbsPath())
	return err == nil
}

func getConfigFileAbsPath() string {
	return paths.GetUserHomeDir() + "/" + CONFIG_FILE_NAME + "." + CONFIG_FILE_EXT
}

func InitViperGlobalConfigs() {
	viper.AddConfigPath(paths.GetUserHomeDir())
	viper.SetConfigType(CONFIG_FILE_EXT)
	viper.SetConfigName(CONFIG_FILE_NAME)
}
