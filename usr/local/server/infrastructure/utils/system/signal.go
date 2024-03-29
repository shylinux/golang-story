package system

import (
	"os"
	"os/signal"
	"syscall"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

var ch = make(chan os.Signal, 10)

func Watch() error {
	listen(syscall.SIGINT)
	listen(syscall.SIGQUIT)
	Printfln("signal listen %s %s", syscall.SIGINT, syscall.SIGQUIT)
	logs.Infof("signal listen %s %s", syscall.SIGINT, syscall.SIGQUIT)
	for {
		select {
		case s, ok := <-ch:
			if !ok {
				return nil
			}
			switch s {
			case syscall.SIGINT:
				os.Exit(1)
			case syscall.SIGQUIT:
				os.Exit(0)
			}
		}
	}
	return nil
}
func listen(s syscall.Signal) {
	signal.Notify(ch, s)
}
