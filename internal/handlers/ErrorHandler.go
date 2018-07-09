package handlers

import (
	"log"
	"github.com/ttacon/chalk"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(chalk.Red, err.Error(), chalk.Reset)
	}
}
