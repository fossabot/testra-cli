package utils

import (
	"github.com/caarlos0/spin"
	"fmt"
)

var s = spin.New("\033[34m%s Processing...\033[0m")

func StartSpin() {
	s.Set(spin.Spin6)
	fmt.Println()
	s.Start()
}

func StopSpin() {
	s.Stop()
}