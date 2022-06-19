package kubernetes

import (
	"path"

	"shylinux.com/x/ice"
	kit "shylinux.com/x/toolkits"
)

type kubernetes struct {
	ice.Code
	linux string `data:"https://github.com/prometheus/prometheus/releases/download/v2.36.1/prometheus-2.36.1.linux-amd64.tar.gz"`
	list  string `name:"list port path auto start install" help:"可视化"`
}

func (s kubernetes) Start(m *ice.Message, arg ...string) {
	s.Code.Start(m, "", "./prometheus", func(p string) []string {
		return []string{kit.Format("--web.listen-address=:%s", path.Base(p))}
	})
}

func init() { ice.CodeModCmd(kubernetes{}) }
