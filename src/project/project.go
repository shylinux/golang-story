package project

import (
	"path"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/tcp"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

type project struct {
	ice.Code
	ice.Hash
	short    string `data:"name"`
	field    string `data:"time,name,path,repos,bin,pid,port"`
	create   string `name:"create repos='shylinux.com/x/space'" help:"创建"`
	generate string `name:"generate" help:"生成"`
	list     string `name:"list name auto create" help:"项目管理"`
}

func (s project) Create(m *ice.Message, arg ...string) {
	p := path.Join(ice.USR, path.Base(m.Option(nfs.REPOS)))
	m.Cmd(nfs.DIR, nfs.PWD, kit.Dict(nfs.DIR_DEEP, ice.TRUE, nfs.DIR_TYPE, nfs.TYPE_CAT, nfs.DIR_ROOT, "src/project/server/"), func(value ice.Maps) {
		p := path.Join(ice.USR, path.Base(m.Option(nfs.REPOS)), value[nfs.PATH])
		m.Cmd(nfs.COPY, p, path.Join("src/project/server/", value[nfs.PATH]))
		if strings.Contains(p, "/bin/") {
			return
		}
		nfs.Rewrite(m.Message, p, func(p string) string {
			if strings.Contains(p, "shylinux.com/x/golang-story/src/project/server") {
				return strings.Replace(p, "shylinux.com/x/golang-story/src/project/server", m.Option(nfs.REPOS), 1)
			}
			return p
		})
	})
	m.Cmd("web.code.git.repos", "init", "origin", m.Option(nfs.REPOS))
	s.Hash.Create(m, append(arg, mdb.NAME, path.Base(p), nfs.PATH, p)...)
}
func (s project) Generate(m *ice.Message, arg ...string) {
	list := m.Cmd(nfs.DIR, "idl/", kit.Dict(nfs.DIR_ROOT, m.Option(nfs.PATH), nfs.DIR_REG, kit.ExtReg("proto"))).Appendv(nfs.PATH)
	m.Option(cli.CMD_DIR, m.Option(nfs.PATH))
	m.Cmdy(cli.SYSTEM, "protoc", "--go_out=./idl", list)
}
func (s project) Build(m *ice.Message, arg ...string) {
	defer web.ToastProcess(m.Message)()
	web.PushStream(m.Message)
	s.System(m, m.Option(nfs.PATH), "go", "build", "-o", "bin/"+m.Option(mdb.NAME), "main.go")
	s.Hash.Modify(m, kit.Simple(m.OptionSimple(mdb.NAME), mdb.TIME, m.Time(), ice.BIN, "bin/"+m.Option(mdb.NAME))...)
}
func (s project) Start(m *ice.Message, arg ...string) {
	s.Daemon(m, m.Option(nfs.PATH), "bin/"+m.Option(mdb.NAME), "--service.addr="+":"+m.Option(tcp.PORT))
	s.Hash.Modify(m, kit.Simple(m.OptionSimple(mdb.NAME, tcp.PORT), "pid", m.Result())...)
	m.Cmd(web.SPIDE, mdb.CREATE, m.Option(mdb.NAME), "http://localhost:"+m.Option(tcp.PORT))
}
func (s project) Test(m *ice.Message, arg ...string) {
	ctx.ProcessField(m.Message, "apitest", []string{m.Option(mdb.NAME), m.Option(mdb.NAME)}, arg...)
}
func (s project) List(m *ice.Message, arg ...string) {
	s.Hash.List(m, arg...)
	m.PushAction(s.Generate, s.Build, s.Start, s.Test, s.Remove)
}

func init() { ice.CodeModCmd(project{}) }
