package cli

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func Init(c *container.Container) {
	c.Provide(NewMainServiceCmds)

	c.Provide(NewUserServiceCmds)

	c.Provide(NewAuthServiceCmds)

}

type MainServiceCmds struct{}

func NewMainServiceCmds(_ *UserServiceCmds, _ *AuthServiceCmds) *MainServiceCmds {
	return &MainServiceCmds{}
}
