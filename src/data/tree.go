package data

import (
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/core/wiki"
	kit "shylinux.com/x/toolkits"
)

type treeNode struct {
	data ice.Any
	list []*treeNode
	sup  *treeNode
}

type tree struct {
	ice.Zone
	short string `data:"name"`
	field string `data:"time,id,sup,data"`

	create string `name:"create name" help:"创建"`
	insert string `name:"insert name sup data" help:"添加"`
	list   string `name:"list name auto" help:"树"`
}

func (s tree) show(m *ice.Message, deep int, node *treeNode) {
	m.Echo("%s%v\n", strings.Repeat("  ", deep), node.data)
	for _, v := range node.list {
		s.show(m, deep+1, v)
	}
}
func (s tree) Init(m *ice.Message, arg ...string) {
	m.OptionFields("name")
	s.Hash.List(m).Tables(func(val ice.Maps) {
		list := map[string]*treeNode{}
		list[""] = &treeNode{}
		m.Confv("", kit.KeyHash(val[mdb.NAME], mdb.META, mdb.TARGET), list[""])
		m.OptionFields("id,sup,data")
		s.Zone.List(m, val[mdb.NAME]).Tables(func(value ice.Maps) {
			root := list[value["sup"]]
			node := &treeNode{data: value[mdb.DATA], sup: root}
			list[value[mdb.ID]] = node
			root.list = append(root.list, node)
			m.Confv("", kit.KeyHash(val[mdb.NAME], mdb.LIST, kit.Int(value[mdb.ID])-1, mdb.TARGET), list[""])
		})
	})
}
func (s tree) Insert(m *ice.Message, arg ...string) {
	data := &treeNode{data: m.Option(mdb.DATA)}

	s.Hash.Target(m, m.Option(mdb.NAME), func() ice.Any {
		root := &treeNode{list: []*treeNode{data}}
		data.sup = root
		return root
	})
	if m.Option("sup") != "" {
		m.OptionCB(mdb.SELECT, func(value ice.Map) {
			root := value[mdb.TARGET].(*treeNode)
			root.list = append(root.list, data)
			data.sup = root
		})
		s.Zone.List(m, m.Option(mdb.NAME), m.Option("sup"))
	}
	m.Option(mdb.TARGET, data)
	s.Zone.Insert(m, arg...)
}
func (s tree) List(m *ice.Message, arg ...string) {
	if len(arg) < 1 || arg[0] == "" {
		m.OptionFields("time,name")
		s.Hash.List(m, arg...)
		// m.Action(s.Create)
		m.Display("")
	} else if len(arg) < 2 || arg[1] == "" {
		m.OptionFields("time,id,sup,data")
		s.Zone.List(m, arg...)
		if node, ok := s.Hash.Target(m, arg[0], nil).(*treeNode); ok && len(node.list) > 0 {
			s.show(m, 0, node.list[0])
			m.Cmdy(wiki.CHART, wiki.CHAIN, m.Result())
		}
		m.Action(s.Insert).Render("")
	} else {
		m.OptionFields("detail")
		s.Zone.List(m, arg...)
	}
}

func init() { ice.CodeCtxCmd(tree{}) }
