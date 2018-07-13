package utils

import (
	"fmt"
	"gopkg.in/kyokomi/emoji.v1"

	"github.com/sanity-io/litter"
	"github.com/testra-tech/testra-cli/internal/colors"
)

var successCheckEmoji = emoji.Sprint(":white_check_mark: ")
var infoEmoji = emoji.Sprint(":information_source: ")
var warnEmoji = emoji.Sprint(":warning:")
var dangerEmoji = emoji.Sprint(":no_entry_sign:")

func Info(txt string) {
	wrapWithNewLines(func() { fmt.Println(infoEmoji, colors.GREEN, txt, colors.RESET) })
}

func InfoF(txt string, a ...interface{}) {
	wrapWithNewLines(func() { fmt.Printf(infoEmoji+"%s"+txt+"%s", colors.GREEN, a, colors.RESET) })
}

func Warn(txt string) {
	wrapWithNewLines(func() { fmt.Println(warnEmoji, colors.YELLOW, txt, colors.RESET) })
}

func WarnF(txt string, a ...interface{}) {
	wrapWithNewLines(func() { fmt.Printf(warnEmoji+"%s"+txt+"%s", colors.YELLOW, a, colors.RESET) })
}

func Danger(txt string) {
	wrapWithNewLines(func() { fmt.Println(dangerEmoji, colors.RED, txt, colors.RESET) })
}

func DangerF(txt string, a ...interface{}) {
	wrapWithNewLines(func() { fmt.Printf(dangerEmoji+"%s"+txt+"%s", colors.YELLOW, a, colors.RESET) })
}

func Success(txt string) {
	Info(txt)
}

func SuccessF(txt string, a ...interface{}) {
	InfoF(txt, a)
}

func wrapWithNewLines(fn func()) {
	fmt.Println()
	fn()
	fmt.Println()
}

func DumpStruct(a interface{}) {
	fmt.Print("\n", colors.MAGENTA)
	litter.Dump(a)
	fmt.Print("\n")
}
