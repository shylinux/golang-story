package runtime

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/core/code"
	"github.com/shylinux/toolkits"
)

var Index = &ice.Context{Name: "runtime", Help: "虚拟机",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"runtime": {Name: "runtime", Help: "虚拟机", Value: kit.Data(kit.MDB_SHORT, "name")},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"runtime": {Name: "runtime", Help: "虚拟机", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(ice.CLI_SYSTEM, "go", "doc", "runtime")
			m.Set(ice.MSG_APPEND)
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
