package markdown

import (
	"bytes"
	"path"

	"github.com/yuin/goldmark"
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/nfs"
)

type md struct{ ice.Lang }

func (s md) Init(m *ice.Message, arg ...string) {
	s.Lang.Init(m, nfs.SCRIPT, m.Resource(""))
}
func (s md) Render(m *ice.Message, arg ...string) {
	var buf bytes.Buffer
	if !m.Warn(goldmark.Convert([]byte(m.Cmdx(nfs.CAT, path.Join(m.Option(nfs.PATH), m.Option(nfs.FILE)))), &buf)) {
		m.Echo(buf.String())
	}
}
func (s md) Engine(m *ice.Message, arg ...string) {
	var buf bytes.Buffer
	if !m.Warn(goldmark.Convert([]byte(m.Cmdx(nfs.CAT, path.Join(m.Option(nfs.PATH), m.Option(nfs.FILE)))), &buf)) {
		m.Echo(buf.String())
	}
}

func init() { ice.Cmd("web.wiki.md", md{}) }
