package controller

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
)

func Init(c *container.Container) {
	c.Provide(NewMainController)
	c.Provide(NewAuthController)
	c.Provide(NewUserController)
}

type MainController struct{ *server.MainServer }

func NewMainController(server *server.MainServer, auth *AuthController, user *UserController) *MainController {
	return &MainController{server}
}
