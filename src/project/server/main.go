package main

import (
	"go.uber.org/dig"

	"shylinux.com/x/golang-story/src/project/server/application"
	"shylinux.com/x/golang-story/src/project/server/domain"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/service"
)

func main() {
	container := dig.New()
	service.Init(container)
	application.Init(container)
	domain.Init(container)
	infrastructure.Init(container)

	container.Invoke(func(s *service.MainService) error {
		return s.Run()
	})
}
