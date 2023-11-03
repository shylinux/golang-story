package office

import (
	"path"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

type xls struct {
	path string `data:"usr/local/office/"`
	list string `name:"xls path auto" help:"表格"`
}

func (s xls) Upload(m *ice.Message, arg ...string) {
	ls := kit.Simple(m.Optionv(ice.MSG_UPLOAD))
	m.Cmd(web.CACHE, web.WATCH, ls[0], path.Join(m.Config(nfs.PATH), m.Option(nfs.PATH), ls[1]))
}
func (s xls) Download(m *ice.Message, arg ...string) {
	m.ProcessOpen(kit.MergeURL(path.Join(ice.SHARE_LOCAL, m.Config(nfs.PATH), m.Option(nfs.PATH)), ice.POD, m.Option(ice.MSG_USERPOD)))
}
func (s xls) Trash(m *ice.Message, arg ...string) {
	nfs.Trash(m.Message, path.Join(m.Config(nfs.PATH), m.Option(nfs.PATH)))
}
func (s xls) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 || strings.HasSuffix(arg[0], nfs.PS) {
		m.Cmdy(nfs.DIR, arg, kit.Dict(nfs.DIR_ROOT, m.Config(nfs.PATH))).Action(s.Upload).PushAction(s.Download, s.Trash)
	} else {
		m.Display("").Option(nfs.PATH, kit.MergeURL(path.Join(ice.SHARE_LOCAL, m.Config(nfs.PATH), arg[0]), ice.POD, m.Option(ice.MSG_USERPOD)))
	}
}

func init() { ice.Cmd("web.wiki.office.xls", xls{}) }
