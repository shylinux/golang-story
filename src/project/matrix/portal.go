package matrix

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

const (
	ICON     = "icon"
	TITLE    = "title"
	DISPLAY  = "display"
	CHILDREN = "children"
)

type portal struct {
	Server
	list string `name:"list portal views view auto" help:"门户"`
}

func (s portal) List(m *ice.Message, arg ...string) {
	res := web.SpideGet(m.Message, m.Conf(s.Server, kit.Keym(web.DOMAIN))+"/api/menu/")
	kit.For(kit.Value(res, "Portal"), func(value ice.Map) {
		if len(arg) == 0 {
			m.Push(PORTAL, value[mdb.NAME]).Push("", value, []string{TITLE, DISPLAY, "redirect"})
			return
		} else if value[mdb.NAME] != arg[0] {
			return
		}
		kit.For(value[CHILDREN], func(value ice.Map) {
			if len(arg) == 1 {
				m.Push("views", value[mdb.NAME]).Push("", value, []string{ICON, TITLE})
				return
			} else if value[mdb.NAME] != arg[1] {
				return
			}
			kit.For(value[CHILDREN], func(value ice.Map) {
				m.Push("view", value[mdb.NAME]).Push("", value, []string{ICON, TITLE, DISPLAY})
			})
		})
	})
	m.Echo(kit.Formats(res))
}

func init() { ice.CodeCtxCmd(portal{}) }
