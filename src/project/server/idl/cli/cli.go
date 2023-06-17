package cli

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func Init(c *container.Container) {
	c.Provide(NewMainServiceCmds)

	c.Provide(NewAuthServiceCmds)

	c.Provide(NewUserServiceCmds)

}

type MainServiceCmds struct{}

func NewMainServiceCmds(_ *AuthServiceCmds, _ *UserServiceCmds) *MainServiceCmds {
	return &MainServiceCmds{}
}
