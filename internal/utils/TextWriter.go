package utils

import (
	"fmt"
	"gopkg.in/kyokomi/emoji.v1"
	//"github.com/mgutz/ansi"
	"github.com/ttacon/chalk"
	"github.com/sanity-io/litter"
)

var successCheckEmoji = emoji.Sprint(":white_check_mark: ")
var infoEmoji = emoji.Sprint(":information_source: ")
var warnEmoji = emoji.Sprint(":warning:")
var dangerEmoji = emoji.Sprint(":no_entry_sign:")

var greenStyle = chalk.Green.NewStyle()
var redStyle = chalk.Red.NewStyle()
var yellowStyle = chalk.Yellow.NewStyle()

func Info(txt string) {
	wrapWithNewLines(func() { fmt.Println(infoEmoji, chalk.Green, txt, chalk.Reset) })
}

func InfoF(txt string, a ...interface{}) {
	wrapWithNewLines(func() { fmt.Printf(infoEmoji+"%s"+txt+"%s", greenStyle, a, chalk.Reset) })
}

func Warn(txt string) {
	wrapWithNewLines(func() { fmt.Println(warnEmoji, chalk.Yellow, txt, chalk.Reset) })
}

func WarnF(txt string, a ...interface{}) {
	wrapWithNewLines(func() { fmt.Printf(warnEmoji+"%s"+txt+"%s", yellowStyle, a, chalk.Reset) })
}

func Danger(txt string) {
	wrapWithNewLines(func() { fmt.Println(dangerEmoji, chalk.Red, txt, chalk.Reset) })
}

func DangerF(txt string, a ...interface{}) {
	wrapWithNewLines(func() { fmt.Printf(dangerEmoji+"%s"+txt+"%s", redStyle, a, chalk.Reset) })
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
	fmt.Print("\n", chalk.Bold, chalk.Magenta)
	litter.Dump(a)
	fmt.Print("\n")
}
