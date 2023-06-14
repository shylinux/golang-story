package main

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/server"
)

func main() {
	c := container.New(development.Init, infrastructure.Init)
	c.Invoke(func(s *server.ServerCmds) error { return s.Run() })
}
