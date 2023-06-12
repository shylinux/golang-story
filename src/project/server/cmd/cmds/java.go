package cmds

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/deploy"
)

const JAVA = "java"

type JavaCmds struct {
	deploy *deploy.Deploy
}

func (s *JavaCmds) Install(ctx context.Context, arg ...string) {
	s.deploy.Install(JAVA)
}
func NewJavaCmds(cmds *cmds.Cmds, deploy *deploy.Deploy) *JavaCmds {
	java := &JavaCmds{deploy: deploy}
	cmds.Register(JAVA, `java command
  install
`, java)
	return java
}
