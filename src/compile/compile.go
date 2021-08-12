package compile

import (
	"os"
	"path"
	"runtime"
	"strings"

	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/base/ctx"
	"github.com/shylinux/icebergs/base/nfs"
	"github.com/shylinux/icebergs/base/web"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"
)

const (
	BOOTSTRAP = "bootstrap"
	GOLANG    = "golang"
)
const COMPILE = "compile"

var Index = &ice.Context{Name: GOLANG, Help: "golang", Configs: map[string]*ice.Config{
	COMPILE: {Name: COMPILE, Help: "编译器", Value: kit.Data(
		BOOTSTRAP, "https://dl.google.com/go/go1.4-bootstrap-20171003.tar.gz",
		cli.SOURCE, "https://golang.google.cn/dl/go1.15.5.src.tar.gz",

		cli.LINUX, "https://golang.google.cn/dl/go1.15.5.linux-amd64.tar.gz",
		cli.DARWIN, "https://golang.google.cn/dl/go1.15.5.darwin-amd64.tar.gz",
		cli.WINDOWS, "https://golang.google.cn/dl/go1.15.5.windows-amd64.zip",
	)},
}, Commands: map[string]*ice.Command{
	COMPILE: {Name: "compile auto run install compile source build download", Help: "编译器", Meta: kit.Dict(
		ice.Display("/plugin/local/code/snippet.js"),
	), Action: map[string]*ice.Action{
		web.DOWNLOAD: {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
			m.Cmdy(code.INSTALL, web.DOWNLOAD, m.Conf(COMPILE, kit.Keym(BOOTSTRAP)), path.Join(m.Conf(code.INSTALL, kit.META_PATH), BOOTSTRAP))
		}},
		cli.BUILD: {Name: "build", Help: "构建", Hand: func(m *ice.Message, arg ...string) {
			web.PushStream(m)
			m.Option(cli.CMD_DIR, path.Join(m.Conf(code.INSTALL, kit.META_PATH), BOOTSTRAP, "go/src"))
			m.Option(cli.CMD_ENV, cli.PATH, os.Getenv(cli.PATH), "CGO_ENABLE", "0")
			m.Cmd(cli.SYSTEM, "./all.bash")
			m.StatusTime()
			m.ProcessHold()
		}},
		cli.SOURCE: {Name: "source", Help: "源码", Hand: func(m *ice.Message, arg ...string) {
			m.Cmdy(code.INSTALL, web.DOWNLOAD, m.Conf(COMPILE, kit.Keym(cli.SOURCE)))
		}},
		COMPILE: {Name: "compile", Help: "编译", Hand: func(m *ice.Message, arg ...string) {
			web.PushStream(m)
			m.Option(cli.CMD_DIR, path.Join(m.Conf(code.INSTALL, kit.META_PATH), "go/src"))
			m.Option(cli.CMD_ENV, cli.HOME, os.Getenv(cli.HOME), "CGO_ENABLE", "0", "GOROOT_BOOTSTRAP", kit.Path(m.Conf(code.INSTALL, kit.META_PATH), BOOTSTRAP),
				cli.PATH, strings.Join([]string{kit.Path(m.Conf(code.INSTALL, kit.META_PATH), BOOTSTRAP, "go/bin"), os.Getenv(cli.PATH)}, ":"),
			)
			m.Cmd(cli.SYSTEM, "./all.bash")
			m.ProcessHold()
		}},
		code.INSTALL: {Name: "install", Help: "安装", Hand: func(m *ice.Message, arg ...string) {
			m.Cmdy(code.INSTALL, web.DOWNLOAD, m.Conf(COMPILE, kit.Keym(runtime.GOOS)), ice.USR_LOCAL)
		}},
		cli.RUN: {Name: "run type name text", Help: "运行", Hand: func(m *ice.Message, arg ...string) {
			file := m.Cmd(web.CACHE, web.WRITE, arg).Append(kit.MDB_FILE)
			defer os.Remove(m.Cmdx(nfs.LINK, file+".go", file))
			m.Cmdy(cli.SYSTEM, "go", cli.RUN, file+".go")
			m.StatusTime()
		}},
		code.VIMER: {Name: "vimer", Help: "编辑", Hand: func(m *ice.Message, arg ...string) {
			m.Cmdy(code.VIMER, arg)
		}},
		ctx.COMMAND: {Name: "command", Help: "运行"},
	}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
		m.Option("content", `package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}
`)
	}},
}}

func init() { code.Index.Register(Index, nil) }
