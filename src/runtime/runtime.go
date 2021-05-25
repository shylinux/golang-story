package runtime

import (
	"github.com/shylinux/golang-story/src/compile"
	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	kit "github.com/shylinux/toolkits"
)

const RUNTIME = "runtime"

var Index = &ice.Context{Name: RUNTIME, Help: "虚拟机",
	Configs: map[string]*ice.Config{
		RUNTIME: {Name: RUNTIME, Help: "虚拟机", Value: kit.Data(kit.MDB_SHORT, kit.MDB_NAME)},
	},
	Commands: map[string]*ice.Command{
		RUNTIME: {Name: "runtime", Help: "虚拟机", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(cli.SYSTEM, "go", "doc", RUNTIME)
			m.Set(ice.MSG_APPEND)
		}},
	},
}

func init() { compile.Index.Merge(Index) }
