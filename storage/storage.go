package storage

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/toolkits"
)

var Index = &ice.Context{Name: "storage", Help: "序列化",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"storage": {Name: "storage", Help: "storage", Value: kit.Data(kit.MDB_SHORT, "name")},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"storage": {Name: "storage", Help: "storage", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Echo("hello world")
		}},
	},
}

func init() { cli.Index.Register(Index, nil) }
