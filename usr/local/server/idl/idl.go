package idl

import (
	"shylinux.com/x/golang-story/src/project/server/controller"
	"shylinux.com/x/golang-story/src/project/server/service"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func Init(c *container.Container) {
	c.Provide(NewMainController)

	c.Provide(controller.NewAuthController)
	c.Provide(service.NewAuthService)

	c.Provide(controller.NewMachineController)
	c.Provide(service.NewMachineService)

	c.Provide(controller.NewServiceController)
	c.Provide(service.NewServiceService)

	c.Provide(controller.NewUserController)
	c.Provide(service.NewUserService)

}

type MainController struct{}

func NewMainController(

	_ *controller.AuthController,

	_ *controller.MachineController,

	_ *controller.ServiceController,

	_ *controller.UserController,

) *MainController {
	return &MainController{}
}
