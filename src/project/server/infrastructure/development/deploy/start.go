package deploy

import (
	"fmt"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

func (s *Deploy) Start(name string) error {
	if err := s.Stop(name); err != nil {
		return err
	}
	arg := strings.Split(s.Config.Install.GetTarget(name).Start, " ")
	if pid, err := system.CommandStart(s.BinPath(name), arg[0], arg[1:]...); err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Printf("%d", pid)
		return nil
	}
}
func (s *Deploy) Stop(name string) error {
	arg := strings.Split(s.Config.Install.GetTarget(name).Start, " ")
	return system.CommandStop(s.BinPath(name), arg[0], arg[1:]...)
}
