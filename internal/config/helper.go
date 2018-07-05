package config

import (
	"bufio"
	"os"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/mgutz/ansi"
	"github.com/spf13/viper"
	"encoding/json"
	"strings"
)

var configFileAbsolutePath string

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {

	if doesConfigExist() == false {
		fmt.Printf(ansi.Color("Missing config for testra. ", "red") +
			ansi.Color("Please run command `testra config`", "green"))
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
}

func HandleConfig() {
	if doesConfigExist() == true {

	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Testra Api url : ")
	readLine, _ := reader.ReadString('\n')
	url := strings.TrimSuffix(readLine, "\n")

	fmt.Print("Testra Project Id : ")
	readLine1, _ := reader.ReadString('\n')
	projectId := strings.TrimSuffix(readLine1, "\n")

	config := &Config{url, projectId}

	json, _ := json.Marshal(config)

	fmt.Println(string(json))

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
	home := getHomeDir()

	configFileAbsolutePath = home + "/" + ConfigFileName + "." + ConfigFileExt
	if _, err := os.Stat(configFileAbsolutePath); err == nil {
		return true
	} else {
		return false
	}
}
