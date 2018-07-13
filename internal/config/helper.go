package config

import (
	"os"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"encoding/json"
	"bufio"
	"strings"
	"github.com/testra-tech/testra-cli/internal/utils"
	"regexp"
	"github.com/sanity-io/litter"
)

const (
	UrlRegex = "(https?)(://)([a-z0-9.:]+)([/.]*)"
)

var configFileAbsolutePath string

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {

	// Check for testra config in user's home directory, if not found raise an error and exit
	if doesConfigExist() == false {
		utils.Danger("Missing config for testra. Run command `testra config`")
		os.Exit(1)
	}

	// Set viper configurations
	viper.AddConfigPath(getHomeDir())
	viper.SetConfigType(ConfigFileExt)
	viper.SetConfigName(ConfigFileName)

	// read in environment variables that match
	viper.AutomaticEnv()

	// Read default config file
	viper.ReadInConfig()

	// validate config
	url := viper.GetString("BaseUrl")

	isValid, regexMatch := IsValidBaseUrl(url)

	if !isValid {
		utils.Danger("Invalid base url in config. Run `testra config` to fix it")
		os.Exit(1)
	}

	// Set extracted domain and scheme values
	// regexMatch[3] - Domain including port
	// regexMatch[1] - scheme http or https
	viper.Set("scheme", regexMatch[1])
	viper.Set("domain", regexMatch[3])

	initLitterConfig()
}

func IsValidBaseUrl(url string) (bool, []string) {
	regex := regexp.MustCompile(UrlRegex)
	regexMatch := regex.FindStringSubmatch(url)
	return len(url) != 0 && cap(regexMatch) != 0, regexMatch
}

func HandleConfig() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(NewLine + "Testra API Base Url : ")
	readLine, _ := reader.ReadString('\n')
	baseUrl := strings.TrimSuffix(readLine, NewLine)

	fmt.Print(NewLine + "Default Project Id : ")
	readLine, _ = reader.ReadString('\n')
	projectId := strings.TrimSuffix(readLine, NewLine)

	WriteConfigToFile(baseUrl, projectId)
}

func WriteConfigToFile(baseUrl string, defaultProjectId string) {
	config := &Config{baseUrl, defaultProjectId}
	json, _ := json.Marshal(*config)
	f, _ := os.Create(configFileAbsolutePath)
	defer f.Close()
	f.Write(json)
}

// User's home directory
func getHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return home
}

func doesConfigExist() bool {
	SetConfigFileAbsPath()
	if _, err := os.Stat(configFileAbsolutePath); err == nil {
		return true
	} else {
		return false
	}
}

func SetConfigFileAbsPath() {
	home := getHomeDir()
	configFileAbsolutePath = home + "/" + ConfigFileName + "." + ConfigFileExt
}

func initLitterConfig() {
	// Strip all package names from types
	litter.Config.StripPackageNames = true

	// Sets separator used when multiple arguments are passed to Dump() or Sdump().
	litter.Config.Separator = NewLine
}
