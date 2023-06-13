package system

import (
	"os/exec"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func Command(name string, arg ...string) error {
	logs.Infof("cmd %s %s", name, strings.Join(arg, " "))
	if err := exec.Command(name, arg...).Run(); err != nil {
		logs.Errorf("cmd failure %s %s %s", name, arg, err)
		return err
	}
	return nil
}
