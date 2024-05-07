package runtime

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

type studio struct {
	list string `name:"list refresh" icon:"go.png"`
}

func (s studio) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		m.Cmdy(runtime{}).Display("").DisplayCSS("")
		kit.If(m.Config(ctx.TOOLS) == "", func() { m.Toolkit(env{}) })
		m.StatusTimeCount(nfs.VERSION, kit.TrimPrefix(m.SystemCmdx(cli.GO, nfs.VERSION), "go version "))
	} else {
		m.ProcessField(runtime{}, arg)
	}
}

func init() { ice.CodeModCtxCmd(studio{}) }
