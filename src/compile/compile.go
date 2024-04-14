package compile

import (
	"os"
	"path"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

const (
	COMPILE = "compile"
)

type compile struct {
	ice.Code
	linux   string `data:"https://golang.google.cn/dl/go1.17.13.linux-amd64.tar.gz"`
	darwin  string `data:"https://golang.google.cn/dl/go1.17.13.darwin-amd64.tar.gz"`
	windows string `data:"https://golang.google.cn/dl/go1.17.13.windows-amd64.zip"`
	source  string `data:"https://golang.google.cn/dl/go1.17.13.src.tar.gz"`

	list string `name:"list path auto order build download install" help:"编译器"`
}

func (c compile) Install(m *ice.Message, arg ...string) {
	c.Code.Install(m, "", ice.USR_LOCAL)
	m.Cmd(cli.SYSTEM, nfs.PUSH, ice.USR_LOCAL_GO_BIN)
}
func (c compile) Build(m *ice.Message, arg ...string) {
	m.Option(cli.CMD_ENV, cli.HOME, os.Getenv(cli.HOME), "CGO_ENABLE", "0", "GOROOT_BOOTSTRAP", kit.Path(_path(m, BOOTSTRAP)),
		cli.PATH, strings.Join([]string{kit.Path(_path(m, BOOTSTRAP, "go/bin")), os.Getenv(cli.PATH)}, ice.DF),
	)
	c.Code.Stream(m, _path(m, "go/src"), "./all.bash")
}
func (c compile) Order(m *ice.Message, arg ...string) {
	c.Code.Order(m, _path(m, "go"), ice.BIN)
}
func (c compile) List(m *ice.Message, arg ...string) {
	m.Cmdy(nfs.DIR, arg, kit.Dict(nfs.DIR_ROOT, ice.USR_LOCAL_GO_BIN))
}

func init() { ice.CodeModCmd(compile{}) }

func _path(m *ice.Message, arg ...string) string {
	return path.Join(ice.USR_INSTALL, path.Join(arg...))
}
