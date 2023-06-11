package main

import (
	"shylinux.com/x/golang-story/src/project/server/idl/api"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	space "shylinux.com/x/golang-story/src/project/server/internal/space/src"
)

func main() {
	container := container.New(space.Init, api.Init, infrastructure.Init)
	container.Invoke(func(s *space.SpaceController) error { return s.Main.Run() })
}
func init() {
	consul.Meta["repos"] = "https://shylinux.com/x/golang-story/src/project/server/internal/space"
	consul.Meta["version"] = "v0.0.1"
}
