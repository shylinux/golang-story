package data

import (
	"math/rand"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/mdb"
	kit "shylinux.com/x/toolkits"
)

type list struct {
	ice.Code
	ice.Zone
	short string `data:"zone"`
	field string `data:"time,id,data"`

	create string `name:"create zone=demo" help:"创建"`
	insert string `name:"insert zone=demo data=10" help:"添加"`
	random string `name:"create zone=demo count=30 max=100" help:"随机"`
	list   string `name:"list zone id auto" help:"表"`
}

func (s list) Random(m *ice.Message, arg ...string) {
	for i := 0; i < kit.Int(m.Option(mdb.COUNT)); i++ {
		s.Zone.Insert(m, mdb.ZONE, m.Option(mdb.ZONE), mdb.DATA, kit.Format(rand.Intn(kit.Int(m.Option("max")))))
	}
}
func (s list) List(m *ice.Message, arg ...string) {
	m.Option(mdb.CACHE_LIMIT, "-2")
	if s.Zone.List(m, arg...); len(arg) < 1 || arg[0] == "" {
		m.Action(s.Create, s.Random)
	}
	m.Display("")
}
func init() { ice.CodeCtxCmd(list{}) }
