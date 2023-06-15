package deploy

import (
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

func (s *Deploy) Build(name string) error {
	if system.Exists(s.BinPath(name)) {
		return nil
	}
	src := s.SrcPath(name)
	for _, cmd := range s.Config.Install.GetTarget(name).Build {
		system.Printfln(cmd)
		args := strings.Split(strings.ReplaceAll(cmd, "$PWD", system.AbsPath(src)), " ")
		if err := system.CommandBuild(src, args[0], args[1:]...); err != nil {
			return err
		}
	}
	return nil
}
