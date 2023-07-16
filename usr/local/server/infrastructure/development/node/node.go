package node

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/deploy"
)

const NODE = "node"

type NodeCmds struct {
	deploy *deploy.DeployCmds
	name   string
}

func (s *NodeCmds) List(ctx context.Context, arg ...string) {
	s.deploy.Download(s.name)
	s.deploy.Unpack(s.name)
	s.deploy.Build(s.name)
	s.deploy.Start(s.name)
}
func NewNodeCmds(cmds *cmds.Cmds, deploy *deploy.DeployCmds) *NodeCmds {
	s := &NodeCmds{deploy: deploy, name: NODE}
	cmds = cmds.Add(s.name, "node runtime cli", s.List)
	return s
}
