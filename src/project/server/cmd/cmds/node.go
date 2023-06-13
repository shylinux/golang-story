package cmds

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/deploy"
)

const NODE = "node"

type NodeCmds struct {
	deploy *deploy.Deploy
}

func (s *NodeCmds) Install(ctx context.Context, arg ...string) {
	s.deploy.Install(NODE)
}
func NewNodeCmds(cmds *cmds.Cmds, deploy *deploy.Deploy) *NodeCmds {
	s := &NodeCmds{deploy: deploy}
	cmds = cmds.Add(NODE, NODE, func(ctx context.Context, arg ...string) {})
	cmds.Add("install", "install", s.Install)
	return s
}
