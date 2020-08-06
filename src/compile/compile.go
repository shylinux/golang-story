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
	Configs: map[string]*ice.Config{
		COMPILE: {Name: COMPILE, Help: "编译器", Value: kit.Data(
			"source", "https://dl.google.com/go/go1.14.1.src.tar.gz",
			"target", kit.Dict(
				"linux", "https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz",
				"darwin", "https://dl.google.com/go/go1.14.2.darwin-amd64.pkg",
				"windows", "https://dl.google.com/go/go1.14.2.windows-amd64.msi",
			),
		)},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		COMPILE: {Name: "compile subcmd auto", Help: "编译器", Action: map[string]*ice.Action{
			"install": {Name: "install", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(cli.SYSTEM, "go", arg).Set(ice.MSG_APPEND)
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
