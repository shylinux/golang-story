package kubernetes

import "shylinux.com/x/ice"

type runtime struct {
	ice.Code
	linux string `data:"https://mirrors.tencent.com/nodejs-release/v16.15.1/node-v16.15.1-linux-x64.tar.xz"`
}

func (s runtime) List(m *ice.Message, arg ...string) {
	s.Code.Source(m, "", arg...)
}

func init() { ice.Cmd("web.code.kubernets.runtime", runtime{}) }
