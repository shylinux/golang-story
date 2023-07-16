package internal

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	space "shylinux.com/x/golang-story/src/project/server/internal/space/src"
)

func Init(c *container.Container) {
	c.Provide(NewInternalController)
	space.Init(c)
}

type MainController struct{}

func NewInternalController(space *space.SpaceController) *MainController {
	return &MainController{}
}
