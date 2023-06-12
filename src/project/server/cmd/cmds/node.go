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
	node := &NodeCmds{deploy: deploy}
	cmds.Register(NODE, `node command
  install
`, node)
	return node
}
