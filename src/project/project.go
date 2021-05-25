package project

import (
	"github.com/shylinux/golang-story/src/compile"
	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	kit "github.com/shylinux/toolkits"
)

const PROJECT = "project"

var Index = &ice.Context{Name: PROJECT, Help: "官方库",
	Configs: map[string]*ice.Config{
		PROJECT: {Name: PROJECT, Help: "官方库", Value: kit.Data()},
	},
	Commands: map[string]*ice.Command{
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

func init() { compile.Index.Merge(Index) }
