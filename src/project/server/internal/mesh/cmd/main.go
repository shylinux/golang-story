package main

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/mysql"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
	"shylinux.com/x/golang-story/src/project/server/internal/mesh/idl"
)

func main() {
	c := container.New(idl.Init, infrastructure.Init)
	c.Provide(mysql.New)
	c.Invoke(func(s *server.MainServer, _ *idl.MainController) error { return s.Run() })
}
