package matrix

import (
	"path"

	yml "gopkg.in/yaml.v3"
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

const (
	SERVER = "server"
	MATRIX = "matrix"
	PORTAL = "portal"

	SERVICE_CONF_REG = `service.*\.(yaml|yml)`
	CONFIG_FILE      = "config.file"
	SERVER_PORT      = "server.port"
	PROXY_PORT       = "proxy.port"
	BIN_SERVER       = "bin/server"
	BIN_MATRIX       = "bin/matrix"
	USR_MATRIX       = "usr/matrix/"
)

type Server struct {
	ice.Code
	domain string `data:"http://localhost:8090"`
	source string `data:"https://shylinux.com/x/matrix"`
	binary string `data:"bin/server"`
	action string `data:"open,iframe"`
	start  string `name:"start config.file* proxy.port*=8090 server.port*=9090"`
	list   string `name:"list path auto start build clone" help:"服务"`
}

func (s Server) Inputs(m *ice.Message, arg ...string) {
	switch arg[0] {
	case CONFIG_FILE:
		m.Cmdy(nfs.DIR, ctx.CONFIG, kit.Dict(nfs.DIR_ROOT, s.dir(m), nfs.DIR_REG, SERVICE_CONF_REG)).CutTo(nfs.PATH, arg[0])
	case SERVER_PORT, PROXY_PORT:
		data := kit.Dict()
		yml.Unmarshal([]byte(m.Cmdx(nfs.CAT, s.dir(m, m.Option(CONFIG_FILE)))), &data)
		m.Push(arg[0], kit.Value(data, arg[0]))
	default:
		s.Code.Inputs(m, arg...)
	}
}
func (s Server) Build(m *ice.Message, arg ...string) {
	s.Code.Stream(m, s.dir(m), cli.MAKE, SERVER)
	s.Code.Stream(m, s.dir(m), cli.MAKE, MATRIX)
	s.Code.Stream(m, s.dir(m), cli.MAKE, PORTAL)
}
func (s Server) Start(m *ice.Message, arg ...string) {
	s.Code.Daemon(m, s.dir(m), kit.Simple(m.Config(nfs.BINARY), s.Code.Args(m))...)
}
func (s Server) List(m *ice.Message, arg ...string) {
	s.Code.List(m, m.Config(nfs.BINARY))
	stats := []string{}
	kit.If(nfs.Exists(m, USR_MATRIX), func() { stats = append(stats, nfs.PATH, USR_MATRIX) })
	kit.If(nfs.Exists(m, USR_MATRIX+BIN_SERVER), func() { stats = append(stats, nfs.PATH, BIN_SERVER) })
	m.StatusTimeCount(stats)
}
func (s Server) Iframe(m *ice.Message, arg ...string) {
	m.EchoIFrame(m.Config(web.DOMAIN)).ProcessInner()
}
func (s Server) Open(m *ice.Message, arg ...string) {
	m.ProcessOpen(m.Config(web.DOMAIN))
}

func init() { ice.CodeCtxCmd(Server{}) }

func (s Server) dir(m *ice.Message, arg ...string) string {
	return path.Join(ice.USR, path.Base(m.Config(nfs.SOURCE)), path.Join(arg...))
}
