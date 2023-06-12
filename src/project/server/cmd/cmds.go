package main

import (
	"shylinux.com/x/golang-story/src/project/server/cmd/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func main() {
	container.New(cmds.Init, infrastructure.Init).Invoke(func(cmds *cmds.Cmds) { cmds.Run() })
}
