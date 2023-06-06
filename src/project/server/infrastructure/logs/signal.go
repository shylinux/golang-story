package logs

import (
	"os"
	"os/signal"
	"syscall"
)

var ch = make(chan os.Signal, 10)

func init() {
	listen(syscall.SIGINT)
	listen(syscall.SIGTERM)
	go func() {
		for {
			select {
			case s, ok := <-ch:
				if !ok {
					return
				}
				switch s {
				case syscall.SIGTERM:
					os.Exit(0)
				case syscall.SIGINT:
					os.Exit(1)
				}
			}
		}
	}()
}
func listen(s syscall.Signal) {
	signal.Notify(ch, s)
}
