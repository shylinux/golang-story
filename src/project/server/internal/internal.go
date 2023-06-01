package internal

import (
	"go.uber.org/dig"
	space "shylinux.com/x/golang-story/src/project/server/internal/space/src"
)

func Init(container *dig.Container) {
	container.Provide(NewInternalController)
	space.Init(container)
}

type InternalController struct{}

func NewInternalController(space *space.SpaceController) *InternalController {
	return &InternalController{}
}
