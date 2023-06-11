package logs

import (
	"os"
	"os/signal"
	"syscall"
)

var ch = make(chan os.Signal, 10)

func watch() {
	listen(syscall.SIGINT)
	listen(syscall.SIGQUIT)
	for {
		select {
		case s, ok := <-ch:
			if !ok {
				return
			}
			switch s {
			case syscall.SIGINT:
				os.Exit(1)
			case syscall.SIGQUIT:
				os.Exit(0)
			}
		}
	}
}
func listen(s syscall.Signal) {
	signal.Notify(ch, s)
}
