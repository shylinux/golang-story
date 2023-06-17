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

	c.Provide(controller.NewUserController)
	c.Provide(service.NewUserService)

}

type MainController struct{}

func NewMainController(

	_ *controller.AuthController,

	_ *controller.UserController,

) *MainController {
	return &MainController{}
}
