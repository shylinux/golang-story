package compile

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/core/code"
	"github.com/shylinux/toolkits"
)

var Index = &ice.Context{Name: "compile", Help: "compile",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"compile": {Name: "compile", Help: "compile", Value: kit.Data()},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"compile": {Name: "compile", Help: "compile", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Echo("hello compile world")
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
