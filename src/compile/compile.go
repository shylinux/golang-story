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
	linux   string `data:"https://golang.google.cn/dl/go1.22.2.linux-amd64.tar.gz"`
	darwin  string `data:"https://golang.google.cn/dl/go1.22.2.darwin-amd64.tar.gz"`
	windows string `data:"https://golang.google.cn/dl/go1.22.2.windows-amd64.zip"`
	source  string `data:"https://golang.google.cn/dl/go1.22.2.src.tar.gz"`

	list string `name:"list cmd auto"`
}

func (s compile) Install(m *ice.Message, arg ...string) {
	s.Code.Install(m, "", ice.USR_LOCAL)
	m.Cmd(cli.SYSTEM, nfs.PUSH, ice.USR_LOCAL_GO_BIN)
}
func (c compile) Build(m *ice.Message, arg ...string) {
	m.Option(cli.CMD_ENV, cli.HOME, os.Getenv(cli.HOME), "CGO_ENABLE", "0", "GOROOT_BOOTSTRAP", kit.Path(_path(m, BOOTSTRAP)),
		cli.PATH, strings.Join([]string{kit.Path(_path(m, BOOTSTRAP, "go/bin")), os.Getenv(cli.PATH)}, ice.DF),
	)
	c.Code.Stream(m, _path(m, "go/src"), "./all.bash")
}
func (s compile) Order(m *ice.Message, arg ...string) {
	s.Code.Order(m, _path(m, cli.GO), ice.BIN)
}
func (s compile) List(m *ice.Message, arg ...string) {
	if cli.SystemFind(m.Message, cli.GO) == "" {
		s.Code.List(m, "", arg...)
	} else {
		m.SystemCmdy(cli.GO, "help", arg)
	}
}

func init() { ice.CodeModCmd(compile{}) }

func _path(m *ice.Message, arg ...string) string {
	return path.Join(ice.USR_INSTALL, path.Join(arg...))
}
