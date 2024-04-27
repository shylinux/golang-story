package goquery

import (
	"encoding/csv"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web/html"
	kit "shylinux.com/x/toolkits"

	query "github.com/PuerkitoBio/goquery"
)

type goquery struct {
	ice.Hash
	short string `data:"path,tags"`
	field string `data:"time,hash,name,text,path,tags"`
	save  string `name:"save file*"`
}

func (s goquery) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		s.Hash.List(m, arg...)
		return
	}
	msg := s.Hash.List(m.Spawn(), arg...)
	q, e := query.NewDocument(msg.Append(nfs.PATH))
	if m.Warn(e) {
		return
	}
	q.Find(msg.Append(nfs.TAGS)).Each(func(i int, q *query.Selection) {
		switch query.NodeName(q) {
		case html.TABLE:
			s.table(m, q)
		}
	})
	m.Action(s.Save)
}
func (s goquery) Save(m *ice.Message, arg ...string) {
	f, _, e := nfs.CreateFile(m.Message, m.Option(nfs.FILE))
	if m.Warn(e) {
		return
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	s.List(m, m.Option(mdb.HASH))
	m.Table(func(value ice.Maps, index int, head []string) {
		kit.If(index == 0, func() { w.Write(head) })
		data := []string{}
		kit.For(head, func(k string) { data = append(data, value[k]) })
		w.Write(data)
	})
}

func init() { ice.CodeCtxCmd(goquery{}) }

func (s goquery) table(m *ice.Message, q *query.Selection) {
	head := []string{}
	q.Find(html.TR).Each(func(i int, q *query.Selection) {
		if i == 0 {
			if q.Find(html.TH).Each(func(i int, q *query.Selection) { head = append(head, strings.TrimSpace(q.Text())) }); len(head) == 0 {
				q.Find(html.TD).Each(func(i int, q *query.Selection) { head = append(head, strings.TrimSpace(q.Text())) })
			}
		} else {
			q.Find(html.TD).Each(func(i int, q *query.Selection) { m.Push(head[i], strings.TrimSpace(q.Text())) })
		}
	})
}
