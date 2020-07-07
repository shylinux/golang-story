package runtime

import (
	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"
)

const RUNTIME = "runtime"

var Index = &ice.Context{Name: RUNTIME, Help: "虚拟机",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		RUNTIME: {Name: RUNTIME, Help: "虚拟机", Value: kit.Data(kit.MDB_SHORT, "name")},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		RUNTIME: {Name: "runtime", Help: "虚拟机", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(cli.SYSTEM, "go", "doc", RUNTIME)
			m.Set(ice.MSG_APPEND)
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
