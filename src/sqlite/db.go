package sqlite

import (
	"path"

	"shylinux.com/x/ice"
)

type db struct{ ice.Lang }

func (s db) Render(m *ice.Message, arg ...string) {
	m.EchoFields("web.code.sqlite.client", path.Join(arg[2], arg[1]))
}
func (s db) Engine(m *ice.Message, arg ...string) {
}

func init() { ice.Cmd("web.code.sqlite.db", db{}) }
