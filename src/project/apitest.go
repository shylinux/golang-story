package project

import (
	"net/http"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

type apitest struct {
	ice.Code
	ice.Hash
	short   string `data:"path"`
	field   string `data:"time,path,args"`
	create  string `name:"create path args"`
	request string `name:"request" help:"请求"`
	list    string `name:"list name client.name path auto create" help:"接口测试"`
}

func (s apitest) Inputs(m *ice.Message, arg ...string) {
	switch arg[0] {
	case nfs.PATH:
		m.Cmdy(nfs.DIR, nfs.PWD, nfs.PATH, kit.Dict(nfs.DIR_ROOT, "usr/"+m.Option(mdb.NAME)+"/api", nfs.DIR_DEEP, ice.TRUE, nfs.DIR_TYPE, nfs.TYPE_CAT))
	case "args":
		m.Push(arg[0], m.Cmdx(nfs.CAT, "usr/"+m.Option(mdb.NAME)+"/api/"+m.Option(nfs.PATH)))
	}
}
func (s apitest) Request(m *ice.Message, arg ...string) {
	if strings.HasPrefix(m.Option("args"), "{") {
		m.Cmdy(web.SPIDE, m.Option(web.CLIENT_NAME), web.SPIDE_RAW, http.MethodPost, nfs.PS+m.Option(nfs.PATH), web.SPIDE_DATA, m.Option("args"))
	} else {
		m.Cmdy(web.SPIDE, m.Option(web.CLIENT_NAME), web.SPIDE_RAW, http.MethodGet, nfs.PS+m.Option(nfs.PATH)+"?"+m.Option("args"))
	}
	m.ProcessInner()
}
func (s apitest) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		m.Cmdy(project{})
	} else if len(arg) == 1 {
		m.Cmdy(web.SPIDE, "")
	} else {
		s.Hash.List(m, arg[2:]...)
		m.PushAction(s.Request, s.Hash.Remove)
	}
}

func init() { ice.CodeModCmd(apitest{}) }
