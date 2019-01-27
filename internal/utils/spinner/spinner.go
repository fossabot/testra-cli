package spinner

import (
	"fmt"
	"github.com/caarlos0/spin"
	"github.com/testra/testra-cli/internal/constants/colors"
)

var s = spin.New(colors.Blue + "%s Loading..." + colors.Reset)

func Start() {
	s.Set(spin.Spin6)
	fmt.Println()
	s.Start()
}

func Stop() {
	s.Stop()
}

func LoadWithSpinner(fn func()) {
	Start()
	defer Stop()
	fn()
}