package gotags

import (
	"path"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/core/code"
	"shylinux.com/x/icebergs/core/wiki"
	kit "shylinux.com/x/toolkits"
)

type chain struct {
	ice.Code
	list string `name:"list path file line auto" help:"编译器"`
}

func (s chain) Find(m *ice.Message, arg ...string) {
	if !nfs.ExistsFile(m, path.Join(m.Option(nfs.PATH), nfs.TAGS)) {
		m.Cmd(cli.SYSTEM, "gotags", "-R", kit.Dict(cli.CMD_DIR, m.Option(nfs.PATH)))
	}

	if msg := m.Cmd(code.INNER, nfs.TAGS, arg[0]); msg.Append(nfs.FILE) != "" {
		ctx.ProcessFloat(m.Message, m.PrefixKey(), msg.Append(nfs.PATH), msg.Append(nfs.FILE), msg.Append(nfs.LINE))
		return
	}

	if !strings.HasSuffix(arg[0], ".go") {
		if msg := m.Cmd(cli.SYSTEM, "go", "doc", arg[0]); cli.IsSuccess(msg) {
			ctx.ProcessFloat(m.Message, m.PrefixKey(), "doc", arg[0])
			return
		}
	}
	if nfs.ExistsFile(m, path.Join(m.Option(nfs.PATH), arg[0])) {
		ctx.ProcessFloat(m.Message, m.PrefixKey(), m.Option(nfs.PATH), arg[0], "1")
		return
	}
}
func (s chain) Doc(m *ice.Message, arg ...string) {
	m.Cmdy(cli.SYSTEM, "go", "doc", arg[0])
	ctx.Display(m, "/plugin/local/code/inner.js")
}
func (s chain) List(m *ice.Message, arg ...string) {
	if strings.HasSuffix(arg[0], ice.PS) && !strings.Contains(arg[0], ice.NL) {
		m.Cmdy(code.INNER, arg)
		ctx.Display(m, "/plugin/local/code/inner.js")
		return
	}
	m.Cmdy(wiki.CHART, wiki.CHAIN, arg, kit.Dict(ctx.INDEX, m.PrefixKey()))
}

func init() { ice.CodeCtxCmd(chain{}) }
