package main

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/server"
	"shylinux.com/x/golang-story/src/project/server/internal/mesh/idl/cli"
)

func main() {
	c := container.New(cli.Init, development.Init, infrastructure.Init)
	c.Invoke(func(s *cmds.Cmds, _ *server.ServerCmds, _ *cli.MainServiceCmds) error { return s.Run() })
}
