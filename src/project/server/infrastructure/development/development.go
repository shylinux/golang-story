package development

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/deploy"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/java"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/node"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/project"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/proto"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/server"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/service"
)

func Init(c *container.Container) {
	c.Provide(cmds.New)
	c.Provide(project.NewProjectCmds)
	c.Provide(service.NewServiceCmds)
	c.Provide(proto.NewGenerateCmds)
	c.Provide(deploy.NewDeployCmds)
	c.Provide(server.NewServerCmds)
	c.Provide(java.NewJavaCmds)
	c.Provide(node.NewNodeCmds)
}
