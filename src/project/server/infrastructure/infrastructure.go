package infrastructure

import (
	"context"
	"testing"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/tests"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/proxy"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/token"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/gin"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/goroutine"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/uuid"
)

func Init(container *container.Container) {
	container.Provide(config.New)
	container.Provide(logs.New)
	container.Provide(proxy.New)
	container.Provide(token.New)
	container.Provide(consul.New)
	container.Provide(server.New)
	container.Provide(goroutine.New)
	container.Provide(context.Background)
	container.Provide(grpc.NewServer)
	container.Provide(gin.NewEngine)
	container.Provide(tests.New)
	container.Provide(uuid.New)
}
func Test(t *testing.T, cb func(*tests.Suite) interface{}) {
	container.New(Init).Invoke(func(s *tests.Suite) { s.Run(t, cb(s)) })
}
