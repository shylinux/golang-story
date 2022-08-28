package docker

import "shylinux.com/x/ice"

type runtime struct {
	ice.Code
	linux string `data:"https://download.docker.com/linux/static/stable/x86_64/docker-19.03.9.tgz"`
	list  string `name:"list path auto order install" help:"容器"`
}

func (s runtime) List(m *ice.Message, arg ...string) {
	s.Code.Source(m, "", arg...)
}

func init() { ice.CodeCtxCmd(runtime{}) }
