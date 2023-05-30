package infrastructure

import (
	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/gin"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/log"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/mysql"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/pulsar"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/redis"
)

func Init(container *dig.Container) *dig.Container {
	container.Provide(log.New)
	container.Provide(config.New)
	container.Provide(consul.New)
	container.Provide(pulsar.New)
	container.Provide(redis.New)
	container.Provide(mysql.New)
	container.Provide(grpc.NewServer)
	container.Provide(gin.NewEngine)
	return container
}
