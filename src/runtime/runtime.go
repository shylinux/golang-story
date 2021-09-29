package runtime

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
)

type runtime struct {
	list string `name:"list auto" help:"运行时"`
}

func (r runtime) List(m *ice.Message, arg ...string) {
	m.Cmdy(cli.SYSTEM, "go", "doc", "runtime")
}

func init() { ice.Cmd("web.code.golang.runtime", runtime{}) }
