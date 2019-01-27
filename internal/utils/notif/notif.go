package notif

import (
	"gopkg.in/kyokomi/emoji.v1"
	"fmt"
	"github.com/testra/testra-cli/internal/constants/colors"
	consoleWriter "github.com/testra/testra-cli/internal/utils/console/writer"
)

var infoEmoji = emoji.Sprint(":information_source: ")
var warnEmoji = emoji.Sprint(":warning:")
var dangerEmoji = emoji.Sprint(":no_entry_sign:")

func Info(txt string) {
	consoleWriter.WrapWithNewLines(func() { fmt.Println(infoEmoji, colors.Green, txt, colors.Reset) })
}

func InfoF(txt string, a ...interface{}) {
	consoleWriter.WrapWithNewLines(func() { fmt.Printf(infoEmoji+"%s"+txt+"%s", colors.Green, a, colors.Reset) })
}

func Warn(txt string) {
	consoleWriter.WrapWithNewLines(func() { fmt.Println(warnEmoji, colors.Yellow, txt, colors.Reset) })
}

func WarnF(txt string, a ...interface{}) {
	consoleWriter.WrapWithNewLines(func() { fmt.Printf(warnEmoji+"%s"+txt+"%s", colors.Yellow, a, colors.Reset) })
}

func Danger(txt string) {
	consoleWriter.WrapWithNewLines(func() { fmt.Println(dangerEmoji, colors.Red, txt, colors.Reset) })
}

func DangerF(txt string, a ...interface{}) {
	consoleWriter.WrapWithNewLines(func() { fmt.Printf(dangerEmoji+"%s"+txt+"%s", colors.Yellow, a, colors.Reset) })
}
