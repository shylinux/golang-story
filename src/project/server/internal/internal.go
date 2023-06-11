package internal

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	space "shylinux.com/x/golang-story/src/project/server/internal/space/src"
)

func Init(container *container.Container) {
	container.Provide(NewInternalController)
	space.Init(container)
}

type InternalController struct{}

func NewInternalController(space *space.SpaceController) *InternalController {
	return &InternalController{}
}
