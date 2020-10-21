package compile

import (
	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/base/gdb"
	"github.com/shylinux/icebergs/base/web"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"

	"os"
	"path"
)

const (
	GOLANG  = "golang"
	COMPILE = "compile"
)

var Index = &ice.Context{Name: GOLANG, Help: "golang",
	Configs: map[string]*ice.Config{
		COMPILE: {Name: COMPILE, Help: "编译器", Value: kit.Data(
			"source", "https://golang.google.cn/dl/go1.15.3.src.tar.gz",
			// "source", "https://dl.google.com/go/go1.4-bootstrap-20171003.tar.gz",
			// "linux", "https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz",
			// "darwin", "https://dl.google.com/go/go1.14.2.darwin-amd64.pkg",
			// "windows", "https://dl.google.com/go/go1.14.2.windows-amd64.msi",
		)},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		COMPILE: {Name: "compile port path auto start build download", Help: "编译器", Action: map[string]*ice.Action{
			web.DOWNLOAD: {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, web.DOWNLOAD, m.Conf(COMPILE, kit.META_SOURCE))
			}},
			gdb.BUILD: {Name: "build", Help: "构建", Hand: func(m *ice.Message, arg ...string) {
				if cli.Follow(m) {
					return
				}

				p := path.Join(m.Conf(code.INSTALL, kit.META_PATH), "go/src")
				m.Option(cli.CMD_ENV,
					"GOBIN", kit.Path(path.Join(path.Dir(p), "install/bin")),
					"GOROOT_FINAL", kit.Path(path.Dir(p)),
					"PATH", os.Getenv("PATH"),
					"HOME", os.Getenv("HOME"),
					"GOCACHE", os.Getenv("GOCACHE"),
					"GOPROXY", "https://goproxy.cn,direct",
					"GOPRIVATE", "github.com",
				)
				m.Option(cli.CMD_DIR, p)
				m.Go(func() { m.Cmdy(cli.SYSTEM, "./all.bash", "--no-clean") })
			}},
			gdb.START: {Name: "start", Help: "启动", Hand: func(m *ice.Message, arg ...string) {
				m.Optionv("prepare", func(p string) []string {
					m.Option(cli.CMD_DIR, p)
					return []string{}
				})
				m.Cmdy(code.INSTALL, "start", "go", "bin/go")
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(code.INSTALL, "go", arg)
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
