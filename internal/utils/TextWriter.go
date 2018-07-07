package utils

import (
	"fmt"
	"gopkg.in/kyokomi/emoji.v1"
)

var successCheckEmoji = emoji.Sprint(":white_check_mark: ")
var infoEmoji = emoji.Sprint(":information_source: ")
var warnEmoji = emoji.Sprint(":warning:")
var dangerEmoji = emoji.Sprint(":no_entry_sign:")

func Info(txt string) {
	wrapWithNewLines(func() { fmt.Println(infoEmoji + txt) })
}

func InfoF(txt string, a ...interface{}) {
	wrapWithNewLines(func() { fmt.Printf(infoEmoji + txt, a) })
}

func Warn(txt string) {
	wrapWithNewLines(func() { fmt.Println(warnEmoji + txt) })
}

func WarnF(txt string, a ...interface{}) {
	wrapWithNewLines(func() { fmt.Printf(warnEmoji + txt, a) })
}

func Danger(txt string) {
	wrapWithNewLines(func() { fmt.Println(dangerEmoji + txt) })
}

func DangerF(txt string, a ...interface{}) {
	wrapWithNewLines(func() { fmt.Printf(dangerEmoji + txt, a) })
}

func Success(txt string) {
	wrapWithNewLines(func() { fmt.Println(successCheckEmoji+txt) })
}

func SuccessF(txt string, a ...interface{}) {
	wrapWithNewLines(func() { fmt.Printf(successCheckEmoji+txt, a) })
}

func wrapWithNewLines(fn func()) {
	fmt.Println()
	fn()
	fmt.Println()
}
