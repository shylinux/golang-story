package gogs

import (
	"path"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/icebergs/core/code"
	kit "shylinux.com/x/toolkits"
)

type gogs struct {
	source string `name:"http://mirrors.aliyun.com/gnu/gogs/gogs-4.2.53.tar.gz" help:"代码库"`
}

func (g gogs) Download(m *ice.Message, arg ...string) {
	m.Cmdy(code.INSTALL, web.DOWNLOAD, m.Conf(m.PrefixKey(), kit.META_SOURCE))
}
func (g gogs) Build(m *ice.Message, arg ...string) {
	m.Cmdy(code.INSTALL, cli.BUILD, m.Conf(m.PrefixKey(), kit.META_SOURCE))
}
func (g gogs) Start(m *ice.Message, arg ...string) {
	m.Cmdy(code.INSTALL, cli.START, m.Conf(m.PrefixKey(), kit.META_SOURCE), "bin/gogs")
}
func (p gogs) List(m *ice.Message, arg ...string) {
	m.Cmdy(code.INSTALL, path.Base(m.Conf(m.PrefixKey(), kit.META_SOURCE)), arg)
}

func init() { ice.CodeModCmd(gogs{}) }
