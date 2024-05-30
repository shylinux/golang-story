package compile

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
)

type tools struct {
	ice.Hash
	short string `data:"name"`
	field string `data:"time,name,path,source"`
	list  string `name:"list name auto" help:"工具"`
}

func (s tools) Init(m *ice.Message, arg ...string) {
	s.Hash.Create(m, mdb.NAME, "goimports", nfs.SOURCE, "golang.org/x/tools/cmd/goimports")
	m.Cmd("").Table(func(value ice.Maps) {
		if value[nfs.PATH] == "" {
			s.Hash.Modify(m, mdb.NAME, value[mdb.NAME], nfs.PATH, cli.SystemFind(m.Message, value[mdb.NAME]))
		}
	})
}
func (s tools) List(m *ice.Message, arg ...string) {
	s.Hash.List(m, arg...).Table(func(value ice.Maps) {
		if value[nfs.PATH] == "" {
			m.PushButton(s.Install)
		} else {
			m.PushButton()
		}
	})
}
func (s tools) Install(m *ice.Message, arg ...string) {
	defer m.ToastProcess()()
	m.SystemCmd(cli.GO, m.ActionKey(), m.Option(nfs.SOURCE))
	s.Hash.Modify(m, nfs.PATH, cli.SystemFind(m.Message, m.Option(mdb.NAME)))
}

func init() { ice.CodeModCmd(tools{}) }

type Tools struct{ tools }

func init() { ice.CodeModCmd(Tools{}) }
