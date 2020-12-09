package compile

import (
	"strings"

	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/base/gdb"
	"github.com/shylinux/icebergs/base/nfs"
	"github.com/shylinux/icebergs/base/web"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"

	"os"
	"path"
	"runtime"
)

const (
	COMPILE = "compile"
)
const GOLANG = "golang"

var Index = &ice.Context{Name: GOLANG, Help: "golang",
	Configs: map[string]*ice.Config{
		COMPILE: {Name: COMPILE, Help: "编译器", Value: kit.Data(
			"bootstrap", "https://dl.google.com/go/go1.4-bootstrap-20171003.tar.gz",
			"source", "https://golang.google.cn/dl/go1.15.5.src.tar.gz",

			"linux", "https://golang.google.cn/dl/go1.15.5.linux-amd64.tar.gz",
			"darwin", "https://golang.google.cn/dl/go1.15.5.darwin-amd64.tar.gz",
			"windows", "https://golang.google.cn/dl/go1.15.5.windows-amd64.zip",
		)},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		COMPILE: {Name: "compile path auto install compile source build download", Help: "编译器", Action: map[string]*ice.Action{
			web.DOWNLOAD: {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, web.DOWNLOAD, m.Conf(COMPILE, "meta.bootstrap"))
			}},
			gdb.BUILD: {Name: "build", Help: "构建", Hand: func(m *ice.Message, arg ...string) {
				cli.Follow(m, gdb.BUILD, func() {
					m.Option(cli.CMD_DIR, path.Join(m.Conf(code.INSTALL, kit.META_PATH), "go/src"))
					m.Option(cli.CMD_ENV, "CGO_ENABLE", "0", "PATH", os.Getenv("PATH"))
					m.Cmdy(cli.SYSTEM, "./all.bash")
				})
			}},

			"source": {Name: "source", Help: "源码", Hand: func(m *ice.Message, arg ...string) {
				link := m.Conf(COMPILE, kit.META_SOURCE)
				msg := m.Cmd(web.SPIDE, web.SPIDE_DEV, web.SPIDE_CACHE, web.SPIDE_GET, link)

				name := path.Base(link)
				m.Cmdy(nfs.LINK, path.Join("usr/golang", name), msg.Append(kit.MDB_FILE))

				m.Option(cli.CMD_DIR, "usr/golang")
				m.Cmd(cli.SYSTEM, "tar", "xvf", path.Base(name))
			}},
			"compile": {Name: "compile", Help: "编译", Hand: func(m *ice.Message, arg ...string) {
				cli.Follow(m, "compile", func() {
					m.Option(cli.CMD_DIR, path.Join("usr/golang", "go/src"))
					m.Option(cli.CMD_ENV, "CGO_ENABLE", "0", "PATH", kit.Path("usr/install/go/bin")+":"+os.Getenv("PATH"), "GOROOT_BOOTSTRAP", kit.Path("usr/install/go"))
					m.Cmdy(cli.SYSTEM, "./all.bash")
				})

			}},
			"install": {Name: "install", Help: "安装", Hand: func(m *ice.Message, arg ...string) {
				link := m.Conf(COMPILE, "meta."+runtime.GOOS)
				msg := m.Cmd(web.SPIDE, web.SPIDE_DEV, web.SPIDE_CACHE, web.SPIDE_GET, link)

				name := path.Base(link)
				m.Cmdy(nfs.LINK, path.Join("usr/local", name), msg.Append(kit.MDB_FILE))

				m.Option(cli.CMD_DIR, "usr/local")
				m.Cmd(cli.SYSTEM, "tar", "xvf", path.Base(name))
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Option(nfs.DIR_ROOT, path.Join("usr/local", "go"))
			if len(arg) == 0 || strings.HasSuffix(arg[0], "/") {
				m.Cmdy(nfs.DIR, kit.Select("./", arg, 0))
			} else {
				m.Cmdy(nfs.CAT, kit.Select("./", arg, 0))
			}
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
