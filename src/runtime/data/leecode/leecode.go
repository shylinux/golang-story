package leecode

import (
	"encoding/json"
	"path"
	"reflect"
	"strings"

	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/core/code"
	kit "shylinux.com/x/toolkits"
)

var Index = &ice.Context{Name: "leecode", Help: "刷题"}

func init() { code.Index.Register(Index, nil) }

type leecode struct {
	name string
	help string
	link string
	test string
	hand ice.Any
}

func Cmd(l leecode) {
	l.name = kit.Select(kit.FileName(2), l.name)
	ctx.AddFileCmd(path.Join("/require/", kit.ModPath(1), l.name+".go"), "web.code.leecode."+l.name)
	Index.MergeCommands(ice.Commands{
		l.name: {Name: l.name + " auto link", Help: l.help, Actions: ice.Actions{
			"link": {Name: "link", Help: "链接", Hand: func(m *ice.Message, arg ...string) {
				m.ProcessOpen(l.link)
			}},
		}, Hand: func(m *ice.Message, arg ...string) {
			t := reflect.TypeOf(l.hand)
			v := reflect.ValueOf(l.hand)

			for _, param := range Split(l.test) {
				args := []reflect.Value{}
				for j := 0; j < t.NumIn(); j++ {
					p := reflect.New(t.In(j))
					json.Unmarshal([]byte(param[j]), p.Interface())
					args = append(args, p.Elem())
					m.Push(kit.Format("in%d", j), param[j])
				}

				ok := true
				res := v.Call(args)
				for j, p := range param[t.NumIn():] {
					m.Push(kit.Format("out%d", j), p)
					if data := res[j].Interface(); kit.Format(data) != p {
						m.Push("res", "failure: "+kit.Format(data))
						ok = false
						break
					}
				}
				if ok {
					m.Push("res", "success")
				}
				m.StatusTimeCount()
			}
		}},
	})
}
func Split(p string) (res [][]string) {
	for _, p := range strings.Split(p, ice.NL) {
		if strings.TrimSpace(p) == "" {
			continue
		}
		if strings.HasPrefix(strings.TrimSpace(p), "# ") {
			continue
		}
		res = append(res, kit.Split(p, "\t \n", "\t \n"))
	}
	return res
}
