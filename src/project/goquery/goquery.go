package goquery

import (
	"encoding/csv"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"

	"github.com/PuerkitoBio/goquery"
)

type query struct {
	ice.Hash
	short string `data:"path,tags"`
	field string `data:"time,hash,path,tags"`
	save  string `name:"save file*"`
	list  string `name:"list hash auto"`
}

func (s query) Save(m *ice.Message, arg ...string) {
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
func (s query) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		s.Hash.List(m, arg...)
		return
	}
	msg := s.Hash.List(m.Spawn(), arg...)
	doc, err := goquery.NewDocument(msg.Append(nfs.PATH))
	if m.Warn(err) {
		return
	}
	doc.Find(msg.Append(nfs.TAGS)).Each(func(i int, s *goquery.Selection) {
		if goquery.NodeName(s) == "table" {
			head := []string{}
			s.Find("tr").Each(func(i int, s *goquery.Selection) {
				if i == 0 {
					s.Find("th").Each(func(i int, s *goquery.Selection) {
						head = append(head, strings.TrimSpace(s.Text()))
					})
				} else {
					s.Find("td").Each(func(i int, s *goquery.Selection) {
						m.Push(head[i], strings.TrimSpace(s.Text()))
					})
				}
			})
		}
	})
	m.Action(s.Save)
}

func init() { ice.Cmd("web.code.goquery.goquery", query{}) }
