package main

import (
	"shylinux.com/x/golang-story/src/project/server/cmd/cmds"
	"shylinux.com/x/golang-story/src/project/server/idl/cli"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func main() {
	// container.New(cmds.Init, infrastructure.Init).Invoke(func(cmds *cmds.Cmds) { cmds.Run() })
	container := container.New(cmds.Init, cli.Init, infrastructure.Init)
	container.Invoke(func(s *cli.MainServiceCmds) { s.Run() })
}
func init() {
	consul.Meta["repos"] = "https://shylinux.com/x/golang-story/src/project/server"
	consul.Meta["version"] = "v0.0.1"
}
