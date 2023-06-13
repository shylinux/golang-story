package cmds

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/deploy"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/proto"
)

func Init(container *container.Container) {
	container.Provide(New)
	container.Provide(cmds.New)
	container.Provide(proto.NewGenerate)
	container.Provide(deploy.New)
	container.Provide(NewServerCmds)
	container.Provide(NewJavaCmds)
	container.Provide(NewNodeCmds)
	container.Provide(NewTuiCmds)
}

type Cmds struct{ *cmds.Cmds }

func New(cmds *cmds.Cmds,
	_ *proto.Generate,
	_ *deploy.Deploy,
	_ *ServerCmds,
	_ *JavaCmds,
	_ *NodeCmds,
	_ *TuiCmds,
) *Cmds {
	return &Cmds{cmds}
}
