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
	"shylinux.com/x/icebergs/misc/git"
	kit "shylinux.com/x/toolkits"
)

type project struct {
	ice.Code
	ice.Hash
	short    string `data:"name"`
	field    string `data:"time,name,path,repos,bin,pid,port"`
	create   string `name:"create repos*='https://shylinux.com/x/demo'"`
	generate string `name:"generate" help:"生成"`
	start    string `name:"start port=11000"`
	list     string `name:"list name auto" help:"项目管理"`
}

func (s project) Create(m *ice.Message, arg ...string) {
	p := path.Join(ice.USR, path.Base(m.Option(nfs.REPOS))) + nfs.PS
	s.Hash.Create(m, append(arg, mdb.NAME, path.Base(p), nfs.PATH, p)...)
	m.Cmd(nfs.DEFS, path.Join(p, ice.SRC_MAIN_GO), nfs.Template(m.Message, ice.SRC_MAIN_GO))
	m.Cmd(nfs.DEFS, path.Join(p, ice.GO_MOD), kit.Format(nfs.Template(m.Message, ice.GO_MOD), strings.TrimPrefix(m.Option(nfs.REPOS), "https://")))
	m.Cmd(nfs.DEFS, path.Join(p, ice.GO_SUM), nfs.Template(m.Message, ice.GO_SUM))
	git.ReposInit(m.Message, m.Option(nfs.REPOS), p)
}
func (s project) Generate(m *ice.Message, arg ...string) {
	list := m.Cmd(nfs.DIR, "idl/", kit.Dict(nfs.DIR_ROOT, m.Option(nfs.PATH), nfs.DIR_REG, kit.ExtReg("proto"))).Appendv(nfs.PATH)
	m.Option(cli.CMD_DIR, m.Option(nfs.PATH))
	m.Cmdy(cli.SYSTEM, "protoc", "--go_out=./idl", list)
}
func (s project) Build(m *ice.Message, arg ...string) {
	defer web.ToastProcess(m.Message)()
	web.PushStream(m.Message)
	s.System(m, m.Option(nfs.PATH), "go", "build", "-o", nfs.BIN+m.Option(mdb.NAME), "main.go")
	if cli.IsSuccess(m.Message) {
		s.Hash.Modify(m, kit.Simple(m.OptionSimple(mdb.NAME), mdb.TIME, m.Time(), ice.BIN, nfs.BIN+m.Option(mdb.NAME))...)
		m.ProcessRefresh()
	} else {
		m.ProcessInner()
	}
}
func (s project) Start(m *ice.Message, arg ...string) {
	s.Daemon(m, m.Option(nfs.PATH), nfs.BIN+m.Option(mdb.NAME), "--service.addr="+nfs.DF+m.Option(tcp.PORT))
	if cli.IsSuccess(m.Message) {
		s.Hash.Modify(m, kit.Simple(m.OptionSimple(mdb.NAME, tcp.PORT), cli.PID, m.Result())...)
		m.Cmd(web.SPIDE, mdb.CREATE, m.Option(mdb.NAME), "http://localhost:"+m.Option(tcp.PORT))
		m.ProcessRefresh()
	}
}
func (s project) Test(m *ice.Message, arg ...string) {
	ctx.ProcessField(m.Message, "xterm", []string{"sh", "", "make test", m.Option(nfs.PATH)}, arg...)
}
func (s project) List(m *ice.Message, arg ...string) {
	s.Hash.List(m, arg...).PushAction(s.Generate, s.Build, s.Start, s.Test, s.Remove)
}

func init() { ice.CodeModCmd(project{}) }
