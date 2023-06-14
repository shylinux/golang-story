package java

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/deploy"
)

const JAVA = "java"

type JavaCmds struct {
	deploy *deploy.Deploy
}

func (s *JavaCmds) Install(ctx context.Context, arg ...string) {
	s.deploy.Install(JAVA)
}
func NewJavaCmds(cmds *cmds.Cmds, deploy *deploy.Deploy) *JavaCmds {
	s := &JavaCmds{deploy: deploy}
	cmds = cmds.Add(JAVA, JAVA, func(ctx context.Context, arg ...string) {})
	cmds.Add("install", "install", s.Install)
	return s
}
