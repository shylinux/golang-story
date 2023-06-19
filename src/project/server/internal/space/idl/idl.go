package idl

import (
	"shylinux.com/x/golang-story/src/project/server/internal/space/controller"
	"shylinux.com/x/golang-story/src/project/server/internal/space/service"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func Init(c *container.Container) {
	c.Provide(NewMainController)

	c.Provide(controller.NewSpaceController)
	c.Provide(service.NewSpaceService)

}

type MainController struct{}

func NewMainController(

	_ *controller.SpaceController,

) *MainController {
	return &MainController{}
}
