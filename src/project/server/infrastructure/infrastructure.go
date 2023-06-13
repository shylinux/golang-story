package infrastructure

import (
	"context"
	"testing"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/gin"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/proxy"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/elasticsearch"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/mysql"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/pulsar"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/redis"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/tests"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/token"
)

func Init(container *container.Container) {
	container.Provide(config.New)
	container.Provide(logs.New)
	container.Provide(proxy.New)
	container.Provide(token.New)
	container.Provide(consul.New)
	container.Provide(server.New)
	container.Provide(redis.New)
	container.Provide(pulsar.New)
	container.Provide(elasticsearch.New)
	container.Provide(mysql.New)
	container.Provide(gin.NewEngine)
	container.Provide(grpc.NewServer)
	container.Provide(context.Background)
	container.Provide(tests.New)
}
func Test(t *testing.T, cb func(*tests.Suite) interface{}) {
	container.New(Init).Invoke(func(s *tests.Suite) { s.Run(t, cb(s)) })
}
