package runtime

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	kit "shylinux.com/x/toolkits"
)

type runtime struct {
	ice.Code
	list string `name:"list pkg key auto"`
}

func (s runtime) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		kit.For(kit.Split(m.SystemCmd(cli.GO, "list", "std").Result()), func(p string) { m.Push("pkg", p) })
	} else {
		m.Copy(m.SystemCmd(cli.GO, "doc", arg))
	}
}

func init() { ice.CodeModCmd(runtime{}) }
