package handlers

import (
	"log"
	"github.com/testra-tech/testra-cli/internal/colors"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(colors.RED, err.Error(), colors.RESET)
	}
}
