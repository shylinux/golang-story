package controller

import (
	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/internal"
	"shylinux.com/x/golang-story/src/project/server/service"
)

func Init(container *dig.Container) {
	container.Provide(NewMainController)
	container.Provide(NewUserController)
	internal.Init(container)
	service.Init(container)
}

type MainController struct{ *infrastructure.MainServer }

func NewMainController(mainServer *infrastructure.MainServer, user *UserController, internal *internal.InternalController) *MainController {
	return &MainController{mainServer}
}
