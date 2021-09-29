package project

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
)

type project struct {
	list string `name:"list auto" help:"官方库"`
}

func (p project) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		m.Cmdy(cli.SYSTEM, "go", "list", "std")
	} else {
		m.Cmdy(cli.SYSTEM, "go", "doc", arg)
	}
}

func init() { ice.Cmd("web.code.golang.project", project{}) }
