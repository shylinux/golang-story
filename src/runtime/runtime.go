package runtime

import (
	"shylinux.com/x/ice"
	kit "shylinux.com/x/toolkits"
)

type runtime struct {
	ice.Code
	list string `name:"list pkg key auto" help:"运行时"`
}

func (s runtime) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		kit.For(kit.Split(s.Code.System(m.Spawn(), "", "go", "list", "std").Result()), func(p string) { m.Push("pkg", p) })
	} else {
		s.Code.System(m, "", kit.Simple("go", "doc", arg)...)
	}
}

func init() { ice.CodeModCmd(runtime{}) }
