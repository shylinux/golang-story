package development

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/deploy"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/java"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/node"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/proto"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/server"
)

func Init(container *container.Container) {
	container.Provide(cmds.New)
	container.Provide(node.NewNodeCmds)
	container.Provide(java.NewJavaCmds)
	container.Provide(server.NewServerCmds)
	container.Provide(proto.NewGenerate)
	container.Provide(deploy.New)
}
