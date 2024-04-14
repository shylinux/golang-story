package compile

import (
	"os"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/nfs"
)

const (
	BOOTSTRAP = "bootstrap"
)

type bootstrap struct {
	ice.Code
	source string `data:"https://dl.google.com/go/go1.4-bootstrap-20171003.tar.gz"`
	list   string `name:"list path auto" help:"编译器"`
}

func (c bootstrap) Download(m *ice.Message, arg ...string) {
	c.Code.Download(m, m.Config(nfs.SOURCE), _path(m, BOOTSTRAP))
}
func (c bootstrap) Build(m *ice.Message, arg ...string) {
	m.Option(cli.CMD_ENV, cli.PATH, os.Getenv(cli.PATH), "CGO_ENABLE", "0")
	c.Code.Stream(m, _path(m, BOOTSTRAP, "go/src"), "./all.bash")
}
func (c bootstrap) List(m *ice.Message, arg ...string) {
	c.Code.List(m, BOOTSTRAP, arg...)
}

func init() { ice.CodeModCmd(bootstrap{}) }
