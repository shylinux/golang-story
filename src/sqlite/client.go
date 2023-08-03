package sqlite

import (
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/nfs"
)

type client struct {
	list string `name:"list path auto"`
}

func (s client) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 || strings.HasSuffix(arg[0], nfs.PS) {
		m.Cmdy(nfs.DIR, arg)
		return
	}
	db, err := gorm.Open(sqlite.Open(arg[0]), &gorm.Config{})
	if m.Warn(err) {
		return
	}
	db = db.Exec(".tables")
	m.Echo("%v", db)
}

func init() { ice.Cmd("web.code.sqlite.client", client{}) }
