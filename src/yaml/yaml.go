package yaml

import (
	"path"

	yml "gopkg.in/yaml.v3"
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

type yaml struct {
	ice.Lang
}

func (s yaml) Init(m *ice.Message, arg ...string) {
	s.Lang.Init(m, nfs.SCRIPT, m.Resource(""))
}
func (s yaml) Render(m *ice.Message, arg ...string) {
	if kit.IsIn(kit.TrimExt(arg[1], "yml", "yaml"), "docker-stack", "docker-compose") {
		ctx.Process(m.Message, "web.code.docker.swarm", nil, arg...)
		return
	}
	data := kit.Dict()
	yml.Unmarshal([]byte(m.Cmdx(nfs.CAT, path.Join(arg[2], arg[1]))), &data)
	m.Echo(kit.Formats(data)).Display("/plugin/story/json.js")
}
func (s yaml) Engine(m *ice.Message, arg ...string) {
	if kit.IsIn(kit.TrimExt(arg[1], "yml", "yaml"), "docker-stack", "docker-compose") {
		ctx.Process(m.Message, "web.code.docker.swarm", nil, arg...)
		return
	}
	data := kit.Dict()
	yml.Unmarshal([]byte(m.Cmdx(nfs.CAT, path.Join(arg[2], arg[1]))), &data)
	defer m.StatusTimeCount()
	m.OptionFields(mdb.DETAIL)
	kit.For(kit.KeyValue(nil, "", data), func(k string, v string) { m.Push(k, v) })
}

func init() { ice.Cmd("web.code.yml", yaml{}) }
func init() { ice.Cmd("web.code.yaml", yaml{}) }
