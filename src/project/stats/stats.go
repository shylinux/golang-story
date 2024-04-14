package stats

import (
	"time"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type stats struct {
	statsTable string `name:"statsTables" event:"stats.tables"`
	list       string `name:"list list"`
}

func (s stats) StatsTables(m *ice.Message, arg ...string) {
	fmts := func(size uint64) string { return kit.Format("%0.2f", float64(size)/float64(1<<30)) }
	vm, _ := mem.VirtualMemory()
	ds, _ := disk.Usage(kit.Path(""))
	if p, e := cpu.Percent(time.Second, false); e == nil {
		s.PushStats(m, "cpu.used", kit.Format("%0.2f", p[0]), "%", "处理器使用率")
	}
	s.PushStats(m, "vcpu.total", kit.Ignore(cpu.Counts(true)), "", "处理器线程")
	s.PushStats(m, "cpu.total", kit.Ignore(cpu.Counts(false)), "", "处理器核数")
	s.PushStats(m, "mem.used", fmts(vm.Used), "G", "内存用量")
	s.PushStats(m, "mem.total", fmts(vm.Total), "G", "内存总量")
	s.PushStats(m, "disk.used", fmts(ds.Used), "G", "磁盘用量")
	s.PushStats(m, "disk.total", fmts(ds.Total), "G", "磁盘总量")
}
func (s stats) List(m *ice.Message, arg ...string) {
	s.StatsTables(m.DisplayStory("stats.js"))
}

func init() { ice.CodeModCmd(stats{}) }

func (s stats) PushStats(m *ice.Message, name string, value ice.Any, arg ...string) {
	web.PushStats(m.Message, name, value, arg...)
}
