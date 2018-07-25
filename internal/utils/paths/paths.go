package paths

import (
	"github.com/mitchellh/go-homedir"
	"log"
)

func GetUserHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	return home
}
