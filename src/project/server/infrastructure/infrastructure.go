package infrastructure

import (
	"context"

	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/gin"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/mysql"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/pulsar"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/redis"
)

func Init(container *dig.Container) *dig.Container {
	container.Provide(logs.New)
	container.Provide(config.New)
	container.Provide(consul.New)
	container.Provide(pulsar.New)
	container.Provide(redis.New)
	container.Provide(mysql.New)
	container.Provide(gin.NewEngine)
	container.Provide(grpc.NewServer)
	container.Provide(context.Background)
	container.Provide(NewMainServer)
	container.Provide(NewProxy)
	return container
}
