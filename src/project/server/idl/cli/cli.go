package cli

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
)

func Init(container *container.Container) {
	container.Provide(NewMainServiceCmds)

	container.Provide(NewUserServiceCmds)

	container.Provide(NewAuthServiceCmds)

}

type MainServiceCmds struct {
	*cmds.Cmds
}

func NewMainServiceCmds(cmds *cmds.Cmds, _ *UserServiceCmds, _ *AuthServiceCmds) *MainServiceCmds {
	return &MainServiceCmds{cmds}
}
