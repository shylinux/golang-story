package cli

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func Init(c *container.Container) {
	c.Provide(NewMainServiceCmds)

	c.Provide(NewSpaceServiceCmds)

}

type MainServiceCmds struct{}

func NewMainServiceCmds(_ *SpaceServiceCmds) *MainServiceCmds {
	return &MainServiceCmds{}
}
