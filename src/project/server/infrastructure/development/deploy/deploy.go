package deploy

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

type Deploy struct{ *config.Config }

func New(conf *config.Config, logger logs.Logger, cmds *cmds.Cmds) *Deploy {
	s := &Deploy{conf}
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
		push := func(repos map[string]config.Target) {
			for _, v := range repos {
				if !v.Export {
					continue
				}
				if _, e := os.Stat(path.Join(USR, v.Install, v.Start)); e == nil {
					list = append(list, path.Dir(system.AbsPath(path.Join(USR, v.Install, v.Start))))
				}
			}
		}
		push(s.Config.Install.Source)
		push(s.Config.Install.Binary)
		switch runtime.GOOS {
		case "linux":
			push(s.Config.Install.Linux)
		case "darwin":
			push(s.Config.Install.Darwin)
		case "windows":
			push(s.Config.Install.Windows)
		}
		fmt.Println(strings.Join(list, "\n"))
		fmt.Println("export PATH=" + strings.Join(list, ":") + ":$PATH")
	})
	cmds.Add("deploy", "deploy command", func(ctx context.Context, arg ...string) {
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
func (s *Deploy) Path(name string) string {
	target := s.Config.Install.GetTarget(name)
	return path.Join(USR, path.Base(target.Address))
}
func (s *Deploy) SrcPath(name string) string {
	target := s.Config.Install.GetTarget(name)
	return path.Join(USR, strings.Split(target.Install, "/")[0])
}
func (s *Deploy) BinPath(name string) string {
	target := s.Config.Install.GetTarget(name)
	return path.Join(USR, target.Install)
}
func (s *Deploy) BinFile(name string) string {
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
