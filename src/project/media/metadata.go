package media

import (
	"os"
	"path"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"

	"github.com/dhowden/tag"
)

type metadata struct {
	list string `name:"list path auto"`
}

func (s metadata) Parse(m *ice.Message, arg ...string) {
	m.Cmd("", arg).Table(func(value ice.Maps) {
		if f, e := os.Open(value[nfs.PATH]); e == nil {
			if t, e := tag.ReadFrom(f); e == nil {
				m.Push("file", value[nfs.PATH])
				m.Push("format", t.Format())
				m.Push("filetype", t.FileType())
				m.Push("title", t.Title())
				if p := t.Picture(); p != nil {
					m.Push("picture", p.MIMEType)
					s, _, _ := nfs.CreateFile(m.Message, path.Join("usr/cover/", kit.TrimSuffix(value[nfs.PATH], ".mp3", ".mp4")+"."+path.Base(p.MIMEType)))
					s.Write(p.Data)
					s.Close()
				} else {
					m.Push("picture", "")
				}
			}
		}
	})
}
func (s metadata) List(m *ice.Message, arg ...string) {
	m.Cmdy("web.wiki.feel", arg).Action(s.Parse)
}

func init() { ice.Cmd("web.wiki.media.metadata", metadata{}) }
