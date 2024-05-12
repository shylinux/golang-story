package runtime

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/icebergs/core/code"
	kit "shylinux.com/x/toolkits"
)

type runtime struct {
	list string `name:"list path name auto"`
}

func (s runtime) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		m.Cmdy(web.CODE_MOD, mdb.RENDER, nfs.MOD, ice.GO_MOD, nfs.PWD).RenameAppend(code.REQUIRE, nfs.PATH)
		kit.For(kit.Split(m.SystemCmd(cli.GO, "list", "std").Result()), func(p string) { m.Push(nfs.PATH, p) })
	} else {
		s.split(m, m.Copy(m.SystemCmd(cli.GO, "doc", arg)).Result()).StatusTimeCountStats(mdb.TYPE)
	}
}

func init() { ice.CodeModCmd(runtime{}) }

func (s runtime) split(m *ice.Message, res string) *ice.Message {
	kit.For(kit.SplitLine(res), func(text string) {
		ls := kit.Split(text, " =()[]{}", "=")
		if len(ls) < 2 {
			return
		}
		switch ls[0] {
		case mdb.TYPE:
			m.Push(mdb.TYPE, ls[2])
			m.Push(mdb.NAME, ls[1])
			m.Push(mdb.TEXT, text)
		case "func":
			m.Push(mdb.TYPE, ls[0])
			if kit.HasPrefix(text, "func (") {
				m.Push(mdb.NAME, kit.Keys(kit.TrimPrefix(ls[2], "*"), ls[3]))
			} else {
				m.Push(mdb.NAME, ls[1])
			}
			m.Push(mdb.TEXT, text)
		}
	})
	return m
}
