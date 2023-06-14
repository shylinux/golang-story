package server

import (
	"context"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"

	"shylinux.com/x/golang-story/src/project/server/controller"
	"shylinux.com/x/golang-story/src/project/server/idl/api"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/deploy"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/java"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/node"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/proto"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/internal"
	"shylinux.com/x/golang-story/src/project/server/service"
)

type ServerCmds struct {
	cmds      *cmds.Cmds
	config    *config.Config
	container *container.Container
}

func (s *ServerCmds) Run() error {
	return s.cmds.Run()
}
func (s *ServerCmds) Start(ctx context.Context, arg ...string) {
	s.container.Add(controller.Init, internal.Init, service.Init, api.Init)
	s.container.Invoke(func(s *controller.MainController, _ *internal.InternalController) error { return s.Run() })
}
func (s *ServerCmds) Restart(ctx context.Context, arg ...string) {
	if buf, err := ioutil.ReadFile(s.config.Logs.Pid); err != nil {
		logs.Errorf("restart failure %s", err)
	} else if pid, err := strconv.ParseInt(string(buf), 10, 64); err != nil {
		logs.Errorf("restart failure %s", err)
	} else if p, e := os.FindProcess(int(pid)); e != nil {
		logs.Errorf("restart failure %s", e)
	} else {
		p.Signal(syscall.SIGINT)
	}
}
func NewServerCmds(container *container.Container, config *config.Config, cmds *cmds.Cmds, _ *proto.Generate, _ *deploy.Deploy, _ *java.JavaCmds, _ *node.NodeCmds) *ServerCmds {
	s := &ServerCmds{cmds, config, container}
	cmds.Add("server", "server", s.Start)
	cmds.Add("restart", "restart", s.Restart)
	return s
}
