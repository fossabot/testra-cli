package check

import (
	"github.com/testra/testra-cli/internal/constants/colors"
	"log"
)

func Err(err error) {
	if err != nil {
		log.Fatal(colors.Red, err.Error(), colors.Reset)
	}
}
