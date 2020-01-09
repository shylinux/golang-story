package runtime

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/toolkits"
)

var Index = &ice.Context{Name: "runtime", Help: "runtime",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"runtime": {Name: "runtime", Help: "runtime", Value: kit.Data(kit.MDB_SHORT, "name")},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"runtime": {Name: "runtime", Help: "runtime", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
            m.Echo("hello world")
		}},
	},
}

func init() { cli.Index.Register(Index, nil) }

