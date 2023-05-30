package internal

import (
	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/internal/space"
)

func Init(container *dig.Container) {
	container.Provide(NewInternalController)
	container.Provide(space.NewSpaceController)
	container.Provide(space.NewSpaceService)
}

type InternalController struct{}

func NewInternalController(space *space.SpaceController) *InternalController {
	return &InternalController{}
}
