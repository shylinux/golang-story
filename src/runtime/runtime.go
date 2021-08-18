package runtime

import (
	"shylinux.com/x/golang-story/src/compile"
	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/cli"
	kit "shylinux.com/x/toolkits"
)

const RUNTIME = "runtime"

var Index = &ice.Context{Name: RUNTIME, Help: "虚拟机", Configs: map[string]*ice.Config{
	RUNTIME: {Name: RUNTIME, Help: "虚拟机", Value: kit.Data(kit.MDB_SHORT, kit.MDB_NAME)},
}, Commands: map[string]*ice.Command{
	RUNTIME: {Name: "runtime", Help: "虚拟机", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
		m.Cmdy(cli.SYSTEM, "go", "doc", RUNTIME)
		m.Set(ice.MSG_APPEND)
	}},
}}

func init() { compile.Index.Merge(Index) }
