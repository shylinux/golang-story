package email

import (
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"mime/multipart"
	"mime/quotedprintable"
	"net/mail"
	"path"
	"strings"
	"time"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/aaa"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/toolkits"
)

const (
	MADDY       = "maddy"
	CREDS       = "creds"
	IMAP_ACCT   = "imap-acct"
	IMAP_MBOXES = "imap-mboxes"
	IMAP_MSGS   = "imap-msgs"
	MBOX        = "mbox"
	UID         = "UID"
)

type creds struct {
	create   string `name:"create email* password*"`
	password string `name:"password email* password*"`
	list     string `name:"list email mbox id auto" help:"邮箱"`
}

func (s creds) Create(m *ice.Message, arg ...string) {
	s.maddy(m, CREDS, m.ActionKey(), "-p", m.Option(aaa.PASSWORD), m.Option(aaa.EMAIL))
	s.maddy(m, IMAP_ACCT, m.ActionKey(), m.Option(aaa.EMAIL))
}
func (s creds) Remove(m *ice.Message, arg ...string) {
	s.maddy(m, CREDS, m.ActionKey(), "-y", m.Option(aaa.EMAIL))
	s.maddy(m, IMAP_ACCT, m.ActionKey(), "-y", m.Option(aaa.EMAIL))
}
func (s creds) Password(m *ice.Message, arg ...string) {
	s.maddy(m, CREDS, m.ActionKey(), "-p", m.Option(aaa.PASSWORD), m.Option(aaa.EMAIL))
}
func (s creds) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		kit.For(kit.SplitLine(s.maddy(m, IMAP_ACCT, ice.LIST)), func(email string) {
			m.Push(aaa.EMAIL, email)
		})
		m.PushAction(s.Password, s.Remove).Action(s.Create)
	} else if len(arg) == 1 {
		kit.For(kit.SplitLine(s.maddy(m, IMAP_MBOXES, ice.LIST, arg[0])), func(text string) {
			ls := strings.SplitN(text, "\t", 2)
			m.Push(MBOX, ls[0]).Push(mdb.STATUS, kit.Join(kit.Split(ls[1], " [\\]"), " "))
		})
	} else if len(arg) == 2 {
		kit.For(kit.SplitLine(s.maddy(m, IMAP_MSGS, ice.LIST, arg[0], arg[1])), func(text string) {
			if strings.TrimSpace(text) == "" {

			} else if kit.HasPrefix(text, UID) {
				ls := kit.Split(text, " :<>")
				m.Push(mdb.ID, ls[1]).Push(nfs.FROM, ls[3])
				ls = strings.SplitN(text, "-", 2)
				m.Push(web.TITLE, s.decode(m, strings.TrimSpace(ls[1])))
			} else {
				ls := strings.SplitN(text, ",", 2)
				t, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", strings.TrimSpace(kit.Select("", ls, -1)))
				m.Push(mdb.TIME, t.Format(ice.MOD_TIME)).Push(mdb.STATUS, kit.Join(kit.Split(ls[0], " [\\]"), " "))
			}
		})
		m.Cut("time,id,from,status,title,what").SortIntR(mdb.ID)
	} else if p, e := mail.ReadMessage(bytes.NewBufferString(s.maddy(m, IMAP_MSGS, "dump", "-u", arg[0], arg[1], arg[2]))); !m.Warn(e) {
		kit.For([]string{"Subject", "Date", "From", "To"}, func(k string) {
			m.Push(mdb.NAME, k).Push(mdb.VALUE, s.decode(m, strings.TrimSpace(strings.ReplaceAll(p.Header[k][0], "\"", ""))))
		})
		nfs.Create(m.Message, path.Join("usr/local/email/", arg[0], arg[1], arg[2], "content.html"), func(w io.Writer, file string) {
			msg, trans := s.parse(m.Spawn(), file, p.Header, p.Body, []func(string) string{})
			data := msg.Result()
			for _, t := range trans {
				data = t(data)
			}
			w.Write([]byte(data))
			m.EchoIFrame("/p/" + file)
			// m.Copy(msg)
		})
	}
}

func init() { ice.Cmd("web.chat.email.creds", creds{}) }

func (s creds) maddy(m *ice.Message, arg ...string) string {
	return m.System(MADDY, arg).Result()
}
func (s creds) parse(m *ice.Message, file string, head map[string][]string, body io.Reader, trans []func(string) string) (*ice.Message, []func(string) string) {
	kit.For(head, func(k string, v []string) { m.Push(mdb.NAME, k).Push(mdb.VALUE, v[0]) })
	switch ls := kit.Split(head["Content-Type"][0], "/; ="); ls[0] {
	case "multipart":
		r := multipart.NewReader(body, kit.Select("", ls, -1))
		for {
			if p, e := r.NextPart(); e != nil || p == nil {
				break
			} else {
				m, trans = s.parse(m, file, p.Header, p, trans)
			}
		}
	case "image":
		nfs.Create(m.Message, path.Join(path.Dir(file), kit.Select("some.png", kit.Split(head["Content-Disposition"][0], "; ="), -1)), func(w io.Writer, file string) {
			w.Write(s.decodeBody(m, head, body))
			file = path.Base(file) + "?pod=" + m.Option(ice.MSG_USERPOD)
			trans = append(trans, func(data string) string {
				return strings.ReplaceAll(data, `src="cid:`+head["Content-Id"][0]+`"`, "src=\""+file+"\"")
			})
		})
	default:
		m.SetResult(string(s.decodeBody(m, head, body)))
	}
	return m, trans
}
func (s creds) decodeBody(m *ice.Message, head map[string][]string, body io.Reader) []byte {
	b, _ := ioutil.ReadAll(body)
	switch kit.Select("", head["Content-Transfer-Encoding"], 0) {
	case "quoted-printable":
		b, _ = ioutil.ReadAll(quotedprintable.NewReader(bytes.NewBuffer(b)))
	case "base64":
		b, _ = base64.StdEncoding.DecodeString(string(b))
	}
	return b
}
func (s creds) decode(m *ice.Message, text string) string {
	if strings.HasPrefix(text, "=?utf-8?") {
		if strings.HasPrefix(strings.ToLower(text), "=?utf-8?b?") {
			b, _ := base64.StdEncoding.DecodeString(text[10 : len(text)-2])
			text = string(b)
		}
		if strings.HasPrefix(strings.ToLower(text), "=?utf-8?q?") {
			b, _ := ioutil.ReadAll(quotedprintable.NewReader(bytes.NewBuffer([]byte(text[10 : len(text)-2]))))
			text = string(b)
		}
	}
	return text
}
