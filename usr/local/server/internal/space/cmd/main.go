package main

import (
	"shylinux.com/x/golang-story/src/project/server/idl/api"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/pulsar"
	space "shylinux.com/x/golang-story/src/project/server/internal/space/src"
)

func main() {
	c := container.New(space.Init, api.Init, infrastructure.Init)
	c.Provide(pulsar.New)
	c.Invoke(func(s *space.SpaceController) error { return s.Main.Run() })
}
