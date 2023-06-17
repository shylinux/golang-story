package deploy

import (
	"context"
	"fmt"
	"path"
	"strconv"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

type DeployCmds struct{ *config.Config }

func NewDeployCmds(conf *config.Config, logger logs.Logger, cmds *cmds.Cmds) *DeployCmds {
	s := &DeployCmds{conf}
	check := func(arg []string, action ...func(string) error) {
		if len(arg) == 0 {
			system.Printfln(system.MarshalIndent(conf.Install))
		} else {
			for _, action := range action {
				if err := action(arg[0]); err != nil {
					break
				}
			}
		}
	}
	cmds.Add("download", "deploy download", func(ctx context.Context, arg ...string) {
		check(arg, s.Download)
	})
	cmds.Add("unpack", "deploy unpack", func(ctx context.Context, arg ...string) {
		check(arg, s.Download, s.Unpack)
	})
	cmds.Add("build", "deploy build", func(ctx context.Context, arg ...string) {
		check(arg, s.Download, s.Unpack, s.Build)
	})
	cmds.Add("start", "deploy start", func(ctx context.Context, arg ...string) {
		check(arg, s.Download, s.Unpack, s.Build, s.Start)
	})
	cmds.Add("stop", "deploy stop", func(ctx context.Context, arg ...string) {
		check(arg, s.Stop)
	})
	cmds.Add("env", "deploy env", func(ctx context.Context, arg ...string) {
		list := []string{}
		s.Config.Install.ForEach(func(name string, target config.Target) {
			if p := s.BinFile(name); target.Export && system.Exists(p) {
				list = append(list, path.Dir(system.AbsPath(p)))
			}
		})
		fmt.Println(strings.Join(list, "\n"))
		fmt.Println("export PATH=" + strings.Join(list, ":") + ":$PATH")
	})
	cmds.Add("deploy", "deploy command", func(ctx context.Context, arg ...string) {
		if len(arg) == 0 {
			s.Config.Install.ForEach(func(name string, target config.Target) {
				buf, err := system.ReadFile(path.Join(s.BinPath(name), "log/service.pid"))
				if err != nil {
					return
				}
				pid, _ := strconv.ParseInt(string(buf), 10, 64)
				system.Printfln("%s %d", name, pid)
			})
			return
		}
		args := []string{"-lh"}
		for _, v := range arg {
			check([]string{v}, s.Download, s.Unpack, s.Build)
			args = append(args, s.BinFile(v))
		}
		res, _ := system.Command("", "ls", args...)
		fmt.Printf(res)
	})
	return s
}
func (s *DeployCmds) Path(name string) string {
	target := s.Config.Install.GetTarget(name)
	return path.Join(USR, path.Base(target.Address))
}
func (s *DeployCmds) SrcPath(name string) string {
	target := s.Config.Install.GetTarget(name)
	return path.Join(USR, strings.Split(target.Install, "/")[0])
}
func (s *DeployCmds) BinPath(name string) string {
	target := s.Config.Install.GetTarget(name)
	return path.Join(USR, target.Install)
}
func (s *DeployCmds) BinFile(name string) string {
	target := s.Config.Install.GetTarget(name)
	return path.Join(USR, target.Install, strings.Split(target.Start, " ")[0])
}

const (
	CMD = "cmd"
	IDL = "idl"
	SRC = "src"
	USR = "usr"
	LOG = "log"
)
