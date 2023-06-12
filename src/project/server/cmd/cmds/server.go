package cmds

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/controller"
	"shylinux.com/x/golang-story/src/project/server/idl/api"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/cmds"
	"shylinux.com/x/golang-story/src/project/server/internal"
	"shylinux.com/x/golang-story/src/project/server/service"
)

type ServerCmds struct {
}

func NewServerCmds(container *container.Container, cmds *cmds.Cmds) *ServerCmds {
	cmds.Add("server", "server", func(ctx context.Context, arg ...string) {
		container.Add(controller.Init, internal.Init, service.Init, api.Init)
		container.Invoke(func(s *controller.MainController, _ *internal.InternalController) error { return s.Run() })
	})
	return &ServerCmds{}
}
