package email

import (
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
	create string `name:"create email* password*"`
	list   string `name:"list email mbox auto" help:"邮箱"`
}

func (s creds) Create(m *ice.Message, arg ...string) {
	s.maddy(m, CREDS, m.ActionKey(), "--password", m.Option(aaa.PASSWORD), m.Option(aaa.EMAIL))
	s.maddy(m, IMAP_ACCT, m.ActionKey(), m.Option(aaa.EMAIL))
}
func (s creds) Remove(m *ice.Message, arg ...string) {
	s.maddy(m, CREDS, m.ActionKey(), "-y", m.Option(aaa.EMAIL))
	s.maddy(m, IMAP_ACCT, m.ActionKey(), "-y", m.Option(aaa.EMAIL))
}
func (s creds) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		kit.For(kit.SplitLine(s.maddy(m, IMAP_ACCT, ice.LIST)), func(email string) {
			m.Push(aaa.EMAIL, email)
		})
		m.PushAction(s.Remove).Action(s.Create)
	} else if len(arg) == 1 {
		kit.For(kit.SplitLine(s.maddy(m, IMAP_MBOXES, ice.LIST, arg[0])), func(text string) {
			ls := strings.SplitN(text, "\t", 2)
			m.Push(MBOX, ls[0]).Push(ice.INFO, ls[1])
		})
	} else {
		kit.For(kit.SplitLine(s.maddy(m, IMAP_MSGS, ice.LIST, arg[0], arg[1])), func(text string) {
			if strings.TrimSpace(text) == "" {

			} else if kit.HasPrefix(text, UID) {
				ls := kit.Split(text, " :<>")
				m.Push(UID, ls[1]).Push(nfs.FROM, ls[3])
				ls = strings.SplitN(text, "-", 2)
				m.Push(web.TITLE, ls[1])
			} else {
				ls := strings.SplitN(text, ",", 2)
				t, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", strings.TrimSpace(kit.Select("", ls, -1)))
				m.Push(mdb.TIME, t.Format(ice.MOD_TIME)).Push(ice.INFO, ls[0])
			}
		})
		m.SortIntR(UID)
	}
}

func init() { ice.Cmd("web.chat.email.creds", creds{}) }

func (s creds) maddy(m *ice.Message, arg ...string) string {
	return m.System(MADDY, arg).Result()
}
