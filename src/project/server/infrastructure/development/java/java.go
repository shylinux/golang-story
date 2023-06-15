package java

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/deploy"
)

const JAVA = "java"

type JavaCmds struct {
	deploy *deploy.Deploy
	name   string
}

func (s *JavaCmds) List(ctx context.Context, arg ...string) {
	s.deploy.Download(s.name)
	s.deploy.Unpack(s.name)
	s.deploy.Start(s.name)
}
func NewJavaCmds(cmds *cmds.Cmds, deploy *deploy.Deploy) *JavaCmds {
	s := &JavaCmds{deploy: deploy, name: JAVA}
	cmds = cmds.Add(s.name, "java runtime cli", s.List)
	return s
}
