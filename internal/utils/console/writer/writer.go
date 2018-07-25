package writer

import (
	"fmt"
	"github.com/sanity-io/litter"
	"github.com/testra-tech/testra-cli/internal/constants/colors"
)

func WrapWithNewLines(fn func()) {
	fmt.Println()
	fn()
	fmt.Println()
}

func DumpStruct(a interface{}) {
	fmt.Print("\n", colors.Magenta)
	litter.Dump(a)
	fmt.Print("\n")
}
