package deploy

import (
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

func (s *DeployCmds) Build(name string) error {
	target := s.Config.Install.GetTarget(name)
	for _, p := range target.Plugin {
		args := strings.Split(p, " ")
		system.CommandBuild("", args[0], args[1:]...)
	}
	if system.Exists(s.BinPath(name)) {
		return nil
	}
	src := s.SrcPath(name)
	for _, cmd := range target.Build {
		system.Printfln(cmd)
		args := strings.Split(strings.ReplaceAll(cmd, "$PWD", system.AbsPath(src)), " ")
		if err := system.CommandBuild(src, args[0], args[1:]...); err != nil {
			return err
		}
	}
	return nil
}
