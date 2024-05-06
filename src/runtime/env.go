package runtime

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/web/html"
	kit "shylinux.com/x/toolkits"
)

type env struct {
	list string `name:"list name auto"`
}

func (s env) List(m *ice.Message, arg ...string) {
	kit.For(kit.SplitLine(m.Cmdx(cli.SYSTEM, cli.GO, "env")), func(text string) {
		ls := kit.Split(text, "=")
		m.Push(mdb.NAME, ls[0]).Push(mdb.VALUE, ls[1])
	})
	m.Action(html.FILTER)
}

func init() { ice.Cmd("web.code.golang.env", env{}) }
