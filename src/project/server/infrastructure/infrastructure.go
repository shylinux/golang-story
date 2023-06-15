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
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/goroutine"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/uuid"
)

func Init(c *container.Container) {
	c.Provide(config.New)
	c.Provide(logs.New)
	c.Provide(proxy.New)
	c.Provide(token.New)
	c.Provide(consul.New)
	c.Provide(server.New)
	c.Provide(goroutine.New)
	c.Provide(context.Background)
	c.Provide(grpc.NewServer)
	c.Provide(tests.New)
	c.Provide(uuid.New)
}
func Test(t *testing.T, cb func(*tests.Suite) interface{}) {
	container.New(Init).Invoke(func(s *tests.Suite) { s.Run(t, cb(s)) })
}
