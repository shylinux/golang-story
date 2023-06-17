package idl

import (
	"shylinux.com/x/golang-story/src/project/server/internal/mesh/controller"
	"shylinux.com/x/golang-story/src/project/server/internal/mesh/service"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func Init(c *container.Container) {
	c.Provide(NewMainController)

	c.Provide(controller.NewMachineController)
	c.Provide(service.NewMachineService)

}

type MainController struct{}

func NewMainController(

	_ *controller.MachineController,

) *MainController {
	return &MainController{}
}
