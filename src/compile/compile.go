package compile

import (
	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"
)

const (
	COMPILE = "compile"
)

var Index = &ice.Context{Name: COMPILE, Help: "编译器",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		COMPILE: {Name: COMPILE, Help: "编译器", Value: kit.Data()},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		COMPILE: {Name: "compile", Help: "编译器", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(cli.SYSTEM, "go", arg)
			m.Set(ice.MSG_APPEND)
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
