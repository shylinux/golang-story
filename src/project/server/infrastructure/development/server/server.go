package server

import (
	"context"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/agent"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/deploy"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/java"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/node"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/product"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/project"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/proto"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/service"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

type ServerCmds struct {
	config *config.Config
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
func (s *ServerCmds) List(ctx context.Context, arg ...string) {
}
func NewServerCmds(
	config *config.Config,
	_ *project.ProjectCmds,
	_ *service.ServiceCmds,
	_ *product.ProductCmds,
	_ *proto.GenerateCmds,
	_ *deploy.DeployCmds,
	_ *agent.AgentCmds,
	_ *java.JavaCmds,
	_ *node.NodeCmds,
	cmds *cmds.Cmds,
) *ServerCmds {
	s := &ServerCmds{config}
	cmds = cmds.Add("server", "server command", s.List)
	cmds.Add("restart", "restart", s.Restart)
	return s
}
