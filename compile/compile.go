package compile

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/ctx"
	"github.com/shylinux/toolkits"

	"bufio"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"unicode"
)

func isKey(key string) bool {
	return map[string]bool{
		"if":     true,
		"else":   true,
		"return": true,
	}[key]
}
func isWord(key string) bool {
	for i, k := range []rune(key) {
		if k == '_' || unicode.IsLetter(k) {
			continue
		}
		if i == 0 {
			return false
		}
		if unicode.IsNumber(k) {
			continue
		}
		return false
	}
	return true
}

type check func(*ice.Message, string, int, string, []string, []string) (string, string)

func ctags(m *ice.Message, file string, line int, text string, last []string, list []string) (string, string) {
	switch list[0] {
	case "type":
		return "struct", list[1]
	case "func":
		return "function", list[1]
	}
	return "", ""
}
func scan(m *ice.Message, root string, name string, cb check) {
	s, e := os.Stat(path.Join(root, name))
	if m.Warn(e != nil, "%s", e) {
		return
	}

	if s.IsDir() {
		fs, e := ioutil.ReadDir(path.Join(root, name))
		if m.Warn(e != nil, "%s", e) {
			return
		}

		for _, f := range fs {
			scan(m, path.Join(root, name), f.Name(), cb)
		}
		return
	}

	switch path.Ext(name) {
	case ".go":
	default:
		return
	}

	f, e := os.Open(path.Join(root, name))
	if m.Warn(e != nil, "%s", e) {
		return
	}
	defer f.Close()

	key := m.Rich("tag", nil, kit.Dict("meta", kit.Dict(
		kit.MDB_TYPE, path.Ext(name),
		kit.MDB_NAME, path.Join(root, name),
		kit.MDB_TEXT, name,
		"file", path.Join(root, name),
	)))
	m.Log(ice.LOG_IMPORT, "%s: %s", key, name)

	stat := map[string]int{}
	last := []string{}
	bio := bufio.NewScanner(f)
	text := ""
	for line := 1; bio.Scan(); line++ {
		text = bio.Text()
		list := kit.Split(text, "\t \n", "{[(*:,;)]}")
		if len(list) == 0 {
			continue
		}

		t, n := cb(m, path.Join(root, name), line, text, last, list)
		if t == "" {
			continue
		}
		switch t {
		case "":
			continue
		}

		m.Grow("tag", nil, kit.Dict(
			kit.MDB_TYPE, t, kit.MDB_NAME, n, kit.MDB_TEXT, text,
			"file", path.Join(root, name), "line", line,
		))
		m.Grow("tag", kit.Keys("hash", key), kit.Dict(
			kit.MDB_TYPE, t, kit.MDB_NAME, n, kit.MDB_TEXT, text,
			"file", path.Join(root, name), "line", line,
		))
		m.Grow("tag", kit.Keys("find", n), kit.Dict(
			kit.MDB_TYPE, t, kit.MDB_NAME, n, kit.MDB_TEXT, text,
			"file", path.Join(root, name), "line", line,
		))

		stat[t]++
		m.Push("file", path.Join(root, name))
		m.Push("line", line)
		m.Push("type", t)
		m.Push("name", n)
		m.Push("text", text)
		text = ""
	}
	m.Richs("tag", nil, key, func(key string, value map[string]interface{}) {
		kit.Value(value, "meta.struct", stat["struct"])
		kit.Value(value, "meta.function", stat["function"])
	})
}
func view(m *ice.Message, detail map[string]interface{}) string {
	switch detail["type"] {
	case "struct", "function":
		return m.Cmdx(ice.CLI_SYSTEM, "sed", "-n", kit.Format("%v,/^}/p", kit.Int(detail["line"])), detail["file"])
	}
	return kit.Format(detail["text"])
}

var Index = &ice.Context{Name: "compile", Help: "编译器",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"compile": {Name: "compile", Help: "编译器", Value: kit.Data(kit.MDB_SHORT, "name")},

		"tag": {Name: "tag", Help: "代码分析", Value: kit.Data(kit.MDB_SHORT, "name")},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"tag": {Name: "tag find key index", Help: "tag", Meta: kit.Dict(
			"detail", []string{"精通", "掌握", "熟悉", "了解", "未知"},
		), List: kit.List(
			kit.MDB_INPUT, "text", "name", "file", "action", "auto",
			kit.MDB_INPUT, "text", "name", "id", "action", "auto",
			kit.MDB_INPUT, "button", "name", "查看",
			kit.MDB_INPUT, "button", "name", "返回", "cb", "Last",
		), Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			if len(arg) > 1 && arg[0] == "action" {
				switch arg[1] {
				case "精通":
				case "掌握":
				case "熟悉":
				case "了解":
				}
				m.Richs(cmd, nil, m.Option("file"), func(key string, val map[string]interface{}) {
					m.Grows(cmd, kit.Keys("hash", key), "id", m.Option("id"), func(index int, value map[string]interface{}) {
						value["status"] = arg[1]
					})
				})
				m.Grows(cmd, kit.Keys("find", m.Option("name")), "", "", func(index int, value map[string]interface{}) {
					value["status"] = arg[1]
				})
				return
			}

			if len(arg) == 0 {
				m.Richs(cmd, nil, "*", func(key string, value map[string]interface{}) {
					m.Push(key, value["meta"], []string{"type", "count", "struct", "function", "file"})
				})
				m.Sort("count", "int_r")
				return
			}

			var rest []string
			var detail map[string]interface{}

			switch arg[0] {
			case "add":
				m.Conf(cmd, nil, kit.Data(kit.MDB_SHORT, "name"))
				for _, v := range arg[1:] {
					scan(m, "", v, ctags)
				}

				if m.Option("o") != "" {
					ioutil.WriteFile(m.Option("o"), []byte(m.Table().Result()), 0777)
				}
			case "find":
				if len(arg) == 1 {
					m.Grows(cmd, nil, "", "", func(index int, value map[string]interface{}) {
						m.Push("", value, []string{"id", "type", "name", "text", "line", "file"})
					})
					break
				}
				if len(arg) == 2 {
					m.Grows(cmd, kit.Keys("find", arg[1]), "", "", func(index int, value map[string]interface{}) {
						m.Push("", value, []string{"id", "type", "name", "text", "line", "file"})
					})
					break
				}

				m.Grows(cmd, kit.Keys("find", arg[1]), "id", arg[2], func(index int, value map[string]interface{}) {
					detail = value
				})
				rest = arg[3:]

			case "rand":
				m.Option("cache.limit", -2)
				for i := 0; i < kit.Int(kit.Select("5", arg, 1)); i++ {
					m.Grows(cmd, nil, "id", kit.Format(rand.Intn(kit.Int(m.Conf(cmd, "meta.count")))+1), func(index int, value map[string]interface{}) {
						m.Richs(cmd, nil, value["file"], func(key string, val map[string]interface{}) {
							m.Grows(cmd, kit.Keys("hash", key), "line", kit.Format(value["line"]), func(index int, value map[string]interface{}) {
								m.Push("", value, []string{"id", "type", "name"})
								m.Push("text", view(m, value))
								m.Push("", value, []string{"line", "file"})
							})
						})
					})
				}
			case "stat":
				stat := map[string]int{}
				m.Option("cache.limit", -2)
				m.Richs(cmd, nil, "*", func(key string, val map[string]interface{}) {
					m.Grows(cmd, kit.Keys("hash", key), "", "", func(index int, value map[string]interface{}) {
						stat[kit.Select("未知", value["status"])]++
						stat["总数"]++
					})
				})
				for _, k := range []string{"未知", "了解", "熟悉", "掌握", "精通", "总数"} {
					m.Push(k, stat[k])
					m.Push(k, kit.Format("%d%%", stat[k]*100/stat["总数"]))
				}

			default:
				m.Richs(cmd, nil, arg[0], func(key string, value map[string]interface{}) {
					if len(arg) == 1 {
						m.Grows(cmd, kit.Keys("hash", key), "", "", func(index int, value map[string]interface{}) {
							m.Push("", value, []string{"id", "status", "type", "name", "text", "line"})
						})
						return
					}

					m.Grows(cmd, kit.Keys("hash", key), "id", arg[1], func(index int, value map[string]interface{}) {
						detail = value
						rest = arg[2:]
					})
				})
			}

			if detail != nil {
				if len(rest) > 0 {
					begin, _ := kit.Render(rest[0], detail)
					end, _ := kit.Render(rest[1], detail)
					m.Cmdy(ice.CLI_SYSTEM, "sed", "-n", kit.Format("%s,%sp", string(begin), string(end)), detail["file"])
					m.Set(ice.MSG_APPEND)
					return
				}

				m.Echo(view(m, detail))
			}
		}},
		"compile": {Name: "compile", Help: "编译器", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(ice.CLI_SYSTEM, "go", "help")
			m.Set(ice.MSG_APPEND)
		}},
	},
}

func init() { ctx.Index.Register(Index, nil) }
