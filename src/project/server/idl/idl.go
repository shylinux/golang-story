package idl

import (
	"shylinux.com/x/golang-story/src/project/server/controller"
	"shylinux.com/x/golang-story/src/project/server/service"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func Init(c *container.Container) {
	c.Provide(NewMainController)

	c.Provide(controller.NewUserController)
	c.Provide(service.NewUserService)

	c.Provide(controller.NewAuthController)
	c.Provide(service.NewAuthService)

}

type MainController struct{}

func NewMainController(

	_ controller.UserController,

	_ controller.AuthController,

) *MainController {
	return &MainController{}
}
