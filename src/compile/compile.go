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
	BOOTSTRAP = "bootstrap"
	COMPILE   = "compile"
)

type Compile struct {
	ice.Code
	linux     string `data:"https://golang.google.cn/dl/go1.15.5.linux-amd64.tar.gz"`
	darwin    string `data:"https://golang.google.cn/dl/go1.15.5.darwin-amd64.tar.gz"`
	windows   string `data:"https://golang.google.cn/dl/go1.15.5.windows-amd64.zip"`
	bootstrap string `data:"https://dl.google.com/go/go1.4-bootstrap-20171003.tar.gz"`
	source    string `data:"https://golang.google.cn/dl/go1.15.5.src.tar.gz"`

	c    string `name:"c" help:"源码"`
	gcc  string `name:"gcc" help:"编译"`
	list string `name:"list path auto install order build download gcc c" help:"编译器"`
}

func (c Compile) C(m *ice.Message, arg ...string) {
	c.Code.Download(m, m.Config(BOOTSTRAP), _path(m, BOOTSTRAP))
}
func (c Compile) Gcc(m *ice.Message, arg ...string) {
	m.Option(cli.CMD_ENV, cli.PATH, os.Getenv(cli.PATH), "CGO_ENABLE", "0")
	c.Code.Stream(m, _path(m, BOOTSTRAP, "go/src"), "./all.bash")
}
func (c Compile) Build(m *ice.Message, arg ...string) {
	m.Option(cli.CMD_ENV, cli.HOME, os.Getenv(cli.HOME), "CGO_ENABLE", "0", "GOROOT_BOOTSTRAP", kit.Path(_path(m, BOOTSTRAP)),
		cli.PATH, strings.Join([]string{kit.Path(_path(m, BOOTSTRAP, "go/bin")), os.Getenv(cli.PATH)}, ice.DF),
	)
	c.Code.Stream(m, _path(m, "go/src"), "./all.bash")
}
func (c Compile) Order(m *ice.Message, arg ...string) {
	c.Code.Order(m, _path(m, "go"), ice.BIN)
}
func (c Compile) Install(m *ice.Message, arg ...string) {
	c.Code.Install(m, "", ice.USR_LOCAL)
	m.Cmd(cli.SYSTEM, nfs.PUSH, ice.USR_LOCAL_GO_BIN)
}
func (c Compile) List(m *ice.Message, arg ...string) {
	m.Cmdy(nfs.DIR, arg, kit.Dict(nfs.DIR_ROOT, ice.USR_LOCAL_GO_BIN))
}

func init() { ice.CodeModCmd(Compile{}) }

func _path(m *ice.Message, arg ...string) string {
	return path.Join(ice.USR_INSTALL, path.Join(arg...))
}
