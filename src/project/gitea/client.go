package gitea

import (
	"path"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/aaa"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/icebergs/core/chat/oauth"
	kit "shylinux.com/x/toolkits"
)

const (
	ORGS = "orgs"
	REPO = "repo"

	UPDATED_AT = "updated_at"

	API_V1            = "/api/v1/"
	API_V1_USER       = "/api/v1/user"
	API_V1_ORGS_REPOS = "/api/v1/orgs/%s/repos"
)

type Client struct {
	oauth.Client
	uploads string `name:"uploads name*=publish version*=latest file*" help:"上传"`
	list    string `name:"list hash orgs:text repo:text auto" help:"仓库" icon:"gitea.png"`
}

func (s Client) Inputs(m *ice.Message, arg ...string) {
	switch s.Client.Inputs(m, arg...); m.Option(ctx.ACTION) {
	case "uploads":
		switch arg[0] {
		case mdb.NAME:
			m.Push(arg[0], "publish", "install")
		case nfs.FILE:
			m.Option(nfs.DIR_TYPE, nfs.CAT)
			m.Cmdy(nfs.DIR, ice.USR_PUBLISH, nfs.PATH)
			m.Cmdy(nfs.DIR, ice.USR_INSTALL, nfs.PATH)
		}
	}
}
func (s Client) Orgs(m *ice.Message, arg ...string) {
	res := s.Get(m, kit.Select(m.Option(mdb.HASH), arg, 0), API_V1+ORGS)
	s.rewrite(m, res, ORGS, "id,avatar_url,name,description")
	m.PushAction(s.Uploads)
	m.Echo(kit.Formats(res))
}
func (s Client) Repo(m *ice.Message, arg ...string) {
	res := s.Get(m, kit.Select(m.Option(mdb.HASH), arg, 0), kit.Format(API_V1_ORGS_REPOS, kit.Select(m.Option(ORGS), arg, 1)))
	s.rewrite(m, res, REPO, "updated_at,id,avatar_url,name,language,description,html_url,website").Sort(UPDATED_AT, "str_r")
	m.PushAction(s.Open)
	m.Echo(kit.Formats(res))
}
func (s Client) List(m *ice.Message, arg ...string) {
	switch len(arg) {
	case 0:
		s.Client.List(m, arg...)
	case 1:
		s.Orgs(m, arg...)
	case 2:
		s.Repo(m, arg...)
	}
}
func (s Client) Open(m *ice.Message, arg ...string) {
	m.ProcessOpen(m.Option("html_url"))
}
func (s Client) Uploads(m *ice.Message, arg ...string) {
	msg := s.Hash.List(m, m.Option(mdb.HASH))
	m.Echo(kit.Formats(s.Put(m, m.Option(mdb.HASH), kit.Format("%s/api/packages/%s/generic/%s/%s/%s", msg.Append(web.DOMAIN), m.Option(ORGS),
		m.Option(mdb.NAME), m.Option(nfs.VERSION), path.Base(m.Option(nfs.FILE))), web.SPIDE_FILE, m.Option(nfs.FILE)))).ProcessInner()
}

func init() { ice.CodeCtxCmd(Client{}) }

func (s Client) rewrite(m *ice.Message, res ice.Any, name string, arg ...string) *ice.Message {
	kit.For(res, func(value ice.Map) { m.Push("", value, kit.Split(kit.Join(arg))) })
	m.RenameAppend(mdb.NAME, name, aaa.AVATAR_URL, aaa.AVATAR)
	m.RewriteAppend(func(value, key string, index int) string {
		if key == aaa.AVATAR {
			return kit.Format(`<img src="%s" height="32" width="32">`, value)
		}
		return value
	})
	return m
}
