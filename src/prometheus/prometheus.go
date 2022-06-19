package prometheus

import (
	"os"
	"path"

	"shylinux.com/x/ice"
	kit "shylinux.com/x/toolkits"
)

type prometheus struct {
	ice.Code
	linux string `data:"https://github.com/prometheus/prometheus/releases/download/v2.36.1/prometheus-2.36.1.linux-amd64.tar.gz"`
	list  string `name:"list port path auto start install" help:"可视化"`
}

func (s prometheus) Start(m *ice.Message, arg ...string) {
	s.Code.Start(m, "", "bin/prometheus", func(p string) []string {
		os.MkdirAll(path.Join(p, ice.BIN), ice.MOD_DIR)
		os.Rename(path.Join(p, "prometheus"), path.Join(p, "bin/prometheus"))
		return []string{kit.Format("--web.listen-address=:%s", path.Base(p))}
	})
}
func (s prometheus) List(m *ice.Message, arg ...string) {
	s.Code.List(m, "", arg...)
}

func init() { ice.CodeModCmd(prometheus{}) }
