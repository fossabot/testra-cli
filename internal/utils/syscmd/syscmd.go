package syscmd

import (
	"os/exec"
	"strings"
	"os"
	"log"
)

func LessIt(content string) {
	// Could read $PAGER rather than hardcoding the path.less
	cmd := exec.Command("/usr/bin/less", []string{"-R"}...)

	// Feed it with the string you want to display.
	cmd.Stdin = strings.NewReader(content)

	// This is crucial - otherwise it will write to a null device.
	cmd.Stdout = os.Stdout

	// Fork off a process and wait for it to terminate.
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}