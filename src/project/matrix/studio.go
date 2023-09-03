package matrix

import (
	"shylinux.com/x/ice"
)

type studio struct {
	list string `name:"list refresh" help:"云矩阵" icon:"matrix.png"`
}

func (s studio) List(m *ice.Message, arg ...string) {
	m.DisplayStudio(portal{}, Server{}, user{})
}

func init() { ice.CodeCtxCmd(studio{}) }
