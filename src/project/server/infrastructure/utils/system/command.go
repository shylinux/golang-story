package system

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func Command(dir, name string, arg ...string) (string, error) {
	logs.Infof("cmd %s %s %s", dir, name, strings.Join(arg, " "))
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir
	buf, err := cmd.CombinedOutput()
	if err != nil {
		logs.Errorf("cmd failure %s %s %s %s", name, arg, string(buf), err)
		return string(buf), err
	} else {
		return string(buf), nil
	}
}
func CommandBuild(dir, name string, arg ...string) error {
	logs.Infof("cmd %s %s %s", dir, name, strings.Join(arg, " "))
	fmt.Printf("%s %s\n", name, strings.Join(arg, " "))
	cmd := exec.Command(name, arg...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		logs.Errorf("cmd failure %s %s %s", name, arg, err)
		return err
	}
	return nil
}
func CommandStart(dir, name string, arg ...string) (int, error) {
	f, e := Create(path.Join(dir, "log/service.log"))
	if e != nil {
		logs.Errorf("cmd failure %s %s %s", name, arg, e)
		return 0, e
	}
	logs.Infof("cmd %s %s %s", dir, name, strings.Join(arg, " "))
	cmd := exec.Command(name, arg...)
	cmd.Stderr = f
	cmd.Stderr = f
	cmd.Dir = dir
	if err := cmd.Start(); err != nil {
		logs.Errorf("cmd failure %s %s %s", name, arg, err)
		return 0, err
	} else {
		WriteFile(path.Join(dir, "log/service.pid"), []byte(fmt.Sprintf("%d", cmd.Process.Pid)), 0644)
		return cmd.Process.Pid, nil
	}
}
func CommandStop(dir, name string, arg ...string) error {
	if buf, err := ReadFile(path.Join(dir, "log/service.pid")); err == nil {
		if pid, err := strconv.ParseInt(string(buf), 10, 64); err == nil {
			if p, err := os.FindProcess(int(pid)); err == nil {
				if err := p.Kill(); err == nil || err == os.ErrProcessDone {
					logs.Infof("kill %v", p.Pid)
				} else {
					logs.Errorf("kill %v %v", p.Pid, err)
					return err
				}
			}
		}
	}
	return nil
}
