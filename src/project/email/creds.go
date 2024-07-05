package email

import (
	"bytes"
	"io"
	"net/mail"
	"path"
	"strings"
	"time"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/aaa"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
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
	proto
	create   string `name:"create email* password*"`
	password string `name:"password email* password*"`
	list     string `name:"list email mbox id auto" help:"邮箱管理"`
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
		m.Cut("time,id,from,status,title").SortIntR(mdb.ID)
	} else if p, e := mail.ReadMessage(bytes.NewBufferString(s.maddy(m, IMAP_MSGS, "dump", "-u", arg[0], arg[1], arg[2]))); !m.Warn(e) {
		kit.For([]string{"Subject", "Date", "From", "To"}, func(k string) {
			m.Push(mdb.NAME, k).Push(mdb.VALUE, s.decode(m, strings.TrimSpace(strings.ReplaceAll(p.Header[k][0], "\"", ""))))
		})
		nfs.Create(m.Message, path.Join(USR_LOCAL_EMAIL, arg[0], arg[1], arg[2], CONTENT_HTML), func(w io.Writer, file string) {
			msg, trans := s.parse(m.Spawn(), file, p.Header, p.Body, []func(string) string{})
			w.Write([]byte(s.trans(msg, trans)))
			m.EchoIFrame(m.Resource(file))
		})
	}
}

func init() { ice.Cmd(prefix(), creds{}) }

func (s creds) maddy(m *ice.Message, arg ...string) string {
	return m.System("sudo", MADDY, arg).Result()
}
