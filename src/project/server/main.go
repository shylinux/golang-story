package main

import (
	"go.uber.org/dig"

	"shylinux.com/x/golang-story/src/project/server/controller"
	"shylinux.com/x/golang-story/src/project/server/domain"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/check"
	"shylinux.com/x/golang-story/src/project/server/service"
)

func main() {
	container := dig.New()
	domain.Init(container)
	service.Init(container)
	controller.Init(container)
	infrastructure.Init(container)
	check.Assert(container.Invoke(func(s *controller.MainController, config *config.Config) error { return s.Run(config.Service.Port) }))
}
