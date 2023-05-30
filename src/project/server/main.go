package main

import (
	"go.uber.org/dig"

	"shylinux.com/x/golang-story/src/project/server/controller"
	"shylinux.com/x/golang-story/src/project/server/domain"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/check"
	"shylinux.com/x/golang-story/src/project/server/internal"
	"shylinux.com/x/golang-story/src/project/server/service"
)

func main() {
	container := dig.New()
	domain.Init(container)
	service.Init(container)
	internal.Init(container)
	controller.Init(container)
	check.Assert(infrastructure.Init(container).Invoke(func(s *controller.MainController) error {
		return s.Run()
	}))
}
