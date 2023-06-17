package main

import (
	"shylinux.com/x/golang-story/src/project/server/idl"
	"shylinux.com/x/golang-story/src/project/server/idl/api"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/elasticsearch"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/mysql"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/pulsar"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/redis"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
	"shylinux.com/x/golang-story/src/project/server/internal"
)

func main() {
	c := container.New(idl.Init, internal.Init, api.Init, infrastructure.Init)
	c.Provide(redis.New, pulsar.New, mysql.New, elasticsearch.New)
	c.Invoke(func(s *server.MainServer, _ *idl.MainController, _ *internal.MainController) error { return s.Run() })
}
func init() {
	consul.Meta["repos"] = "https://shylinux.com/x/golang-story/src/project/server"
	consul.Meta["version"] = "v0.0.1"
}
