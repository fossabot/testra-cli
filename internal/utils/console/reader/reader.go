package reader

import (
	"bufio"
	"fmt"
	"github.com/testra/testra-cli/internal/constants/strs"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadUserInput(prompt string) string {
	fmt.Print(strs.NewLine + prompt + " : ")
	readLine, _ := reader.ReadString('\n')
	return strings.TrimSuffix(readLine, strs.NewLine)
}
