package project

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/ctx"
	"github.com/shylinux/toolkits"
)

var Index = &ice.Context{Name: "project", Help: "官方库",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"project": {Name: "project", Help: "官方库", Value: kit.Data(kit.MDB_SHORT, "name")},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"project": {Name: "project", Help: "官方库", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(ice.CLI_SYSTEM, "go", "list", "std")
			m.Set(ice.MSG_APPEND)
		}},
	},
}

func init() { ctx.Index.Register(Index, nil) }
