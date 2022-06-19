package runtime

import (
	"shylinux.com/x/ice"
)

type runtime struct {
	ice.Code
	list string `name:"list auto" help:"运行时"`
}

func (s runtime) List(m *ice.Message, arg ...string) {
	s.Code.System(m, "", "go", "doc", "runtime")
}

func init() { ice.CodeModCmd(runtime{}) }
