package office

import (
	"path"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/lex"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

type xls struct {
	regexp string `data:"xlsx"`
	path   string `data:"usr/local/office/"`
	rename string `name:"rename path* newname*" help:"重命名"`
	list   string `name:"list path auto" help:"表格"`
	parse  string `name:"parse script" help:"解析"`
}

func (s xls) Init(m *ice.Message, arg ...string) {
	m.TransInput("newname", "新路径")
}
func (s xls) Inputs(m *ice.Message, arg ...string) {
	switch arg[0] {
	case nfs.SCRIPT:
		m.Cmdy(nfs.DIR, m.Config(nfs.PATH), nfs.FILE, kit.Dict(nfs.DIR_REG, kit.ExtReg(nfs.JS))).RenameAppend(nfs.FILE, arg[0])
	}
}
func (s xls) Upload(m *ice.Message, arg ...string) {
	ls := kit.Simple(m.Optionv(ice.MSG_UPLOAD))
	m.Cmd(web.CACHE, web.WATCH, ls[0], path.Join(m.Config(nfs.PATH), m.Option(nfs.PATH), ls[1]))
}
func (s xls) Download(m *ice.Message, arg ...string) {
	m.ProcessOpen(kit.MergeURL(path.Join(ice.SHARE_LOCAL, m.Config(nfs.PATH), m.Option(nfs.PATH)), ice.POD, m.Option(ice.MSG_USERPOD)))
}
func (s xls) Rename(m *ice.Message, arg ...string) {
	m.Cmd(nfs.MOVE, m.Option("newname"), m.Option(nfs.PATH), kit.Dict(nfs.DIR_ROOT, m.Config(nfs.PATH)))
}
func (s xls) Trash(m *ice.Message, arg ...string) {
	nfs.Trash(m.Message, path.Join(m.Config(nfs.PATH), m.Option(nfs.PATH)))
}
func (s xls) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 || strings.HasSuffix(arg[0], nfs.PS) {
		m.Cmdy(nfs.DIR, arg, kit.Dict(nfs.DIR_ROOT, m.Config(nfs.PATH), nfs.DIR_REG, kit.ExtReg(m.Config(lex.REGEXP)))).Action(s.Upload).PushAction(s.Download, s.Rename, s.Trash)
	} else {
		m.Action(s.Parse).Display("").Options(nfs.DIR_ROOT, m.Config(nfs.PATH), nfs.PATH, kit.MergeURL(path.Join(ice.SHARE_LOCAL, m.Config(nfs.PATH), arg[0]), ice.POD, m.Option(ice.MSG_USERPOD)))
	}
}
func (s xls) Parse(m *ice.Message, arg ...string) {
}
func init() { ice.WikiCtxCmd(xls{}) }
