package main

import (
	"shylinux.com/x/golang-story/src/project/server/controller"
	"shylinux.com/x/golang-story/src/project/server/idl/api"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/internal"
	"shylinux.com/x/golang-story/src/project/server/service"
)

func main() {
	container := container.New(controller.Init, internal.Init, service.Init, api.Init, infrastructure.Init)
	container.Invoke(func(s *controller.MainController, _ *internal.InternalController) error { return s.Run() })
}
func init() {
	consul.Meta["repos"] = "https://shylinux.com/x/golang-story/src/project/server"
	consul.Meta["version"] = "v0.0.1"
}
