package deploy

import (
	"fmt"
	"path"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

func (s *DeployCmds) Start(name string) error {
	if system.Exists(path.Join(s.BinPath(name), system.LOG_SERVICE_PID)) {
		return nil
	}
	target := s.Config.Install.GetTarget(name)
	arg := strings.Split(target.Start, " ")
	if !target.Daemon {
		if err := system.CommandBuild(s.BinPath(name), arg[0], arg[1:]...); err != nil {
			fmt.Println(err)
			return err
		} else {
			return nil
		}
	}
	if err := s.Stop(name); err != nil {
		return err
	}
	if pid, err := system.CommandStart(s.BinPath(name), arg[0], arg[1:]...); err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Printf("%d", pid)
		return nil
	}
}
func (s *DeployCmds) Stop(name string) error {
	arg := strings.Split(s.Config.Install.GetTarget(name).Start, " ")
	return system.CommandStop(s.BinPath(name), arg[0], arg[1:]...)
}
