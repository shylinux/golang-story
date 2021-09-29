package grafana

import (
	"path"
	"runtime"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/icebergs/core/code"
	kit "shylinux.com/x/toolkits"
)

type grafana struct {
	linux string `data:"https://dl.grafana.com/oss/release/grafana-7.3.4.linux-amd64.tar.gz"`
}

func (g grafana) Install(m *ice.Message, arg ...string) {
	m.Cmdy(code.INSTALL, web.DOWNLOAD, m.Conf(m.PrefixKey(), kit.Keym(runtime.GOOS)))
}
func (g grafana) Start(m *ice.Message, arg ...string) {
	m.Cmdy(code.INSTALL, cli.START, m.Conf(m.PrefixKey(), kit.META_SOURCE), "bin/grafana")
}
func (p grafana) List(m *ice.Message, arg ...string) {
	m.Cmdy(code.INSTALL, path.Base(m.Conf(m.PrefixKey(), kit.META_SOURCE)), arg)
}

func init() { ice.Cmd("web.code.golang.project", grafana{}) }
