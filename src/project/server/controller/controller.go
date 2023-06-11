package controller

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
)

func Init(container *container.Container) {
	container.Provide(NewMainController)
	container.Provide(NewAuthController)
	container.Provide(NewUserController)
}

type MainController struct{ *server.MainServer }

func NewMainController(server *server.MainServer, auth *AuthController, user *UserController) *MainController {
	return &MainController{server}
}
