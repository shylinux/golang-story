package main

import (
	"os"

	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/controller"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/check"
	"shylinux.com/x/golang-story/src/project/server/internal"
	"shylinux.com/x/golang-story/src/project/server/service"
)

func main() {
	container := dig.New()
	service.Init(container)
	internal.Init(container)
	controller.Init(container)
	check.Assert(infrastructure.Init(container).Invoke(func(s *controller.MainController) error { return s.Run() }))
}
func init() {
	wd, _ := os.Getwd()
	consul.Meta["repos"] = "https://shylinux.com/x/golang-story/src/project/server"
	consul.Meta["version"] = "v0.0.1"
	consul.Meta["binary"] = os.Args[0]
	consul.Meta["dir"] = wd
}
