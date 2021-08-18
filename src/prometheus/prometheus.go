package prometheus

import (
	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/gdb"
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/icebergs/core/code"
	kit "shylinux.com/x/toolkits"

	"path"
)

const PROMETHEUS = "prometheus"

var Index = &ice.Context{Name: PROMETHEUS, Help: "命令行",
	Configs: map[string]*ice.Config{
		PROMETHEUS: {Name: PROMETHEUS, Help: "命令行", Value: kit.Data(
			"source", "http://mirrors.aliyun.com/gnu/prometheus/prometheus-4.2.53.tar.gz",
		)},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) { m.Load() }},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) { m.Save() }},

		PROMETHEUS: {Name: "prometheus port path auto start build download", Help: "命令行", Action: map[string]*ice.Action{
			web.DOWNLOAD: {Name: "download", Help: "下载", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, web.DOWNLOAD, m.Conf(PROMETHEUS, kit.META_SOURCE))
			}},
			gdb.BUILD: {Name: "build", Help: "构建", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, gdb.BUILD, m.Conf(PROMETHEUS, kit.META_SOURCE))
			}},
			gdb.START: {Name: "start", Help: "启动", Hand: func(m *ice.Message, arg ...string) {
				m.Cmdy(code.INSTALL, gdb.START, m.Conf(PROMETHEUS, kit.META_SOURCE), "bin/prometheus")
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(code.INSTALL, path.Base(m.Conf(PROMETHEUS, kit.META_SOURCE)), arg)
		}},
	},
}

func init() { code.Index.Register(Index, &web.Frame{}) }