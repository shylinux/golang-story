package internal

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	mesh "shylinux.com/x/golang-story/src/project/server/internal/mesh/idl"
	space "shylinux.com/x/golang-story/src/project/server/internal/space/src"
)

func Init(c *container.Container) {
	c.Provide(NewInternalController)
	space.Init(c)
	mesh.Init(c)
}

type MainController struct{}

func NewInternalController(space *space.SpaceController, _ *mesh.MainController) *MainController {
	return &MainController{}
}
