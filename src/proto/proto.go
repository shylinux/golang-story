package proto

import (
	"path"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/lex"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/core/code"
	kit "shylinux.com/x/toolkits"
)

type proto struct{ ice.Lang }

func (s proto) Init(m *ice.Message, arg ...string) {

	s.Lang.Init(m, code.PREPARE, kit.Dict(code.KEYWORD, kit.List("message")))
}
func (s proto) parse(m *ice.Message, arg ...string) ice.Any {
	block, comment := "", ""
	list := kit.Dict()
	m.Cmd(lex.SPLIT, path.Join(m.Option(nfs.PATH), m.Option(nfs.FILE)), func(ls []string) {
		switch ls[0] {
		case "service":
			list[ls[0]] = append(kit.List(list[ls[0]]), ls[1])
			fallthrough
		case "message":
			block, list[ls[1]] = ls[1], []ice.Any{}
		case "rpc":
			list[block] = append(kit.List(list[block]), ls[1])
			list[ls[1]] = []string{ls[3], ls[7]}
		case "}":
		case "//":
			comment = strings.Join(ls[1:], " ")
		default:
			if block == "" {
				comment = ""
				break
			}
			if ls[0] == "repeated" {
				list[block] = append(kit.List(list[block]), kit.Dict(mdb.TYPE, "[]"+ls[1], mdb.NAME, ls[2], mdb.TEXT, comment))
			} else {
				list[block] = append(kit.List(list[block]), kit.Dict(mdb.TYPE, ls[0], mdb.NAME, ls[1], mdb.TEXT, comment))
			}
			comment = ""
		}
	})
	data := kit.Dict()
	for _, service := range kit.List(list["service"]) {
		service := kit.Format(service)
		serviceDict := kit.Dict()
		for _, api := range kit.List(list[service]) {
			api := kit.Format(api)
			request := kit.Value(list, kit.Keys(api, 0))
			reply := kit.Value(list, kit.Keys(api, 1))
			serviceDict[api] = kit.Dict(request, kit.Value(list, request), reply, kit.Value(list, reply))
		}
		data[service] = serviceDict
	}
	return data
}

func (s proto) Render(m *ice.Message, arg ...string) {
	m.Echo(kit.Formats(s.parse(m, arg...))).Display("/plugin/story/json.js")
}
func (s proto) Engine(m *ice.Message, arg ...string) {
	m.OptionFields(mdb.DETAIL)
	kit.For(kit.KeyValue(nil, "", s.parse(m, arg...)), func(k string, v string) { m.Push(k, v) })
	m.StatusTimeCount()
}

func init() { ice.CodeCmd(proto{}) }
