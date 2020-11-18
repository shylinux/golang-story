package project

import (
	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"
)

const PROJECT = "project"

var Index = &ice.Context{Name: PROJECT, Help: "官方库",
	Configs: map[string]*ice.Config{
		PROJECT: {Name: PROJECT, Help: "官方库", Value: kit.Data()},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		PROJECT: {Name: "project", Help: "官方库", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			if len(arg) == 0 {
				m.Cmdy(cli.SYSTEM, "go", "list", "std")
			} else {
				m.Cmdy(cli.SYSTEM, "go", "doc", arg)
			}
			m.Set(ice.MSG_APPEND)
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
