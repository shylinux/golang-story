package cli

import (
	"shylinux.com/x/golang-story/src/project/server/cmd/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func Init(container *container.Container) {
	container.Provide(NewMainServiceCmds)

	container.Provide(NewSpaceServiceCmds)

}

type MainServiceCmds struct {
	*cmds.Cmds
}

func NewMainServiceCmds(cmds *cmds.Cmds, _ *SpaceServiceCmds) *MainServiceCmds {
	return &MainServiceCmds{cmds}
}
