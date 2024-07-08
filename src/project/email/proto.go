package email

import (
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"mime/multipart"
	"mime/quotedprintable"
	"path"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

type proto struct{}

func (s proto) trans(m *ice.Message, trans []func(string) string) string {
	data := m.Result()
	for _, t := range trans {
		data = t(data)
	}
	return data
}
func (s proto) parse(m *ice.Message, file string, head map[string][]string, body io.Reader, trans []func(string) string) (*ice.Message, []func(string) string) {
	kit.For(head, func(k string, v []string) {
		m.Push(mdb.NAME, k).Push(mdb.VALUE, v[0])
		m.Debug("header: %v %v", k, v)
	})
	switch ls := kit.Split(kit.Select("", head["Content-Type"], 0), "/; ="); ls[0] {
	case "multipart":
		r := multipart.NewReader(body, kit.Select("", ls, -1))
		for {
			if p, e := r.NextPart(); e != nil || p == nil {
				m.Debug("end %v %v", e, p)
				break
			} else {
				m, trans = s.parse(m, file, p.Header, p, trans)
			}
		}
	default:
		if p := kit.Select("", kit.Split(kit.Select("", head["Content-Disposition"], 0), "; ="), -1); p == "" {
			m.SetResult(string(s.decodeBody(m, head, body)))
		} else {
			file = nfs.Create(m.Message, path.Join(path.Dir(file), s.decode(m, p)), func(w io.Writer, file string) {
				n, _ := w.Write(s.decodeBody(m, head, body))
				m.Logs(nfs.SAVE, nfs.FILE, file, nfs.SIZE, n)
				if len(head["Content-Id"]) > 0 {
					file = path.Base(file) + "?pod=" + m.Option(ice.MSG_USERPOD)
					trans = append(trans, func(data string) string {
						cid := kit.ReplaceAll(head["Content-Id"][0], "<", "", ">", "")
						m.Debug("replace cid: %s file: %s", cid, file)
						return strings.ReplaceAll(data, `src="cid:`+cid+`"`, "src=\""+file+"\"")
					})
				} else {
					m.Echo("<a href='%s' target='_blank'>%s</a>", m.Resource(file), path.Base(file))
				}
			})
		}
	}
	return m, trans
}
func (s proto) decodeBody(m *ice.Message, head map[string][]string, body io.Reader) []byte {
	b, e := ioutil.ReadAll(body)
	m.Warn(e)
	switch kit.Select("", head["Content-Transfer-Encoding"], 0) {
	case "quoted-printable":
		b, e = ioutil.ReadAll(quotedprintable.NewReader(bytes.NewBuffer(b)))
		m.Warn(e)
	case "base64":
		_b := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
		if _, e = base64.StdEncoding.Decode(_b, b); e == nil {
			b = _b
		}
	}
	return b
}
func (s proto) decode(m *ice.Message, text string) string {
	if strings.HasPrefix(strings.ToLower(text), "=?utf-8?b?") {
		b, e := base64.StdEncoding.DecodeString(text[10 : len(text)-2])
		text = string(b)
		m.Warn(e)
	} else if strings.HasPrefix(strings.ToLower(text), "=?utf-8?q?") {
		b, e := ioutil.ReadAll(quotedprintable.NewReader(bytes.NewBuffer([]byte(text[10 : len(text)-2]))))
		text = string(b)
		m.Warn(e)
	} else if strings.HasPrefix(strings.ToLower(text), "=?gbk?b?") {
		return text
		b, e := base64.StdEncoding.DecodeString(text[10 : len(text)-2])
		m.Warn(e)
		b, e = ioutil.ReadAll(transform.NewReader(bytes.NewBuffer(b), simplifiedchinese.GBK.NewDecoder()))
		text = string(b)
		m.Warn(e)
	}
	return text
}

func prefix() string { return kit.Keys("web.chat.email", kit.FileName(-1)) }
