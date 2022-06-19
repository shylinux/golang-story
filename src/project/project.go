package project

import (
	"shylinux.com/x/ice"
	kit "shylinux.com/x/toolkits"
)

type project struct {
	ice.Code
	list string `name:"list auto" help:"官方库"`
}

func (s project) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		s.Code.System(m, "", "go", "list", "std")
	} else {
		s.Code.System(m, "", kit.Simple("go", "doc", "std", arg)...)
	}
}

func init() { ice.CodeModCmd(project{}) }
