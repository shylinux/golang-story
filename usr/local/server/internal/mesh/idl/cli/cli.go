package cli

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func Init(c *container.Container) {
	c.Provide(NewMainServiceCmds)

	c.Provide(NewMachineServiceCmds)

}

type MainServiceCmds struct{}

func NewMainServiceCmds(_ *MachineServiceCmds) *MainServiceCmds {
	return &MainServiceCmds{}
}
