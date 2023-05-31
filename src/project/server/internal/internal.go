package internal

import (
	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/internal/space"
)

func Init(container *dig.Container) {
	container.Provide(NewInternalController)
	container.Provide(space.NewSpaceController)
	container.Provide(space.NewSpaceService)
	container.Provide(space.NewUserConsumer)
}

type InternalController struct{}

func NewInternalController(space *space.SpaceController) *InternalController {
	return &InternalController{}
}
