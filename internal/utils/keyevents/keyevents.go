package keyevents

import (
	"os"
	"os/signal"
	"syscall"
)

func RegisterAndExecFuncOnInterrupt(fn func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fn()
		os.Exit(1)
	}()
}
