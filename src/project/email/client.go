package email

import (
	"io"
	"net/smtp"
	"path"
	"strings"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/aaa"
	"shylinux.com/x/icebergs/base/lex"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web/html"
	kit "shylinux.com/x/toolkits"
)

const (
	MAIL            = "mail"
	IMAP            = "imap"
	SMTP            = "smtp"
	MAILBOX         = "mailbox"
	SUBJECT         = "subject"
	CONTENT_HTML    = "content.html"
	USR_LOCAL_EMAIL = "usr/local/email/"

	INBOX   = "INBOX"
	DRAFTS  = "Drafts"
	SENT    = "Sent"
	SPAM    = "Spam"
	TRASH   = "Trash"
	ARCHIVE = "Archive"

	SEND   = "send"
	FROM   = "from"
	TO     = "to"
	CC     = "cc"
	BCC    = "bcc"
	SENDER = "sender"
)

type Client struct {
	ice.Hash
	proto
	important string `data:"true"`
	short     string `data:"username"`
	field     string `data:"time,username,password,imap,smtp"`
	login     string `name:"login username* password* imap* smtp*"`
	create    string `name:"create mailbox*"`
	rename    string `name:"rename mailbox* to*"`
	copy      string `name:"copy to*"`
	move      string `name:"move to*"`
	forword   string `name:"forword to*"`
	reply     string `name:"reply subject* content*:textarea"`
	send      string `name:"send to* cc subject* content*:textarea"`
	list      string `name:"list refresh" help:"邮件" icon:"Mail.png"`
}

func (s Client) Inputs(m *ice.Message, arg ...string) {
	domain := kit.Keys(arg[0], kit.Select("shylinux.com", kit.Split(m.Option(aaa.USERNAME), "@"), 1))
	switch s.Hash.Inputs(m, arg...); arg[0] {
	case IMAP:
		m.Push(arg[0], domain+":993")
		m.Push(arg[0], domain+":143")
	case SMTP:
		m.Push(arg[0], domain+":465")
		m.Push(arg[0], domain+":25")
	case MAILBOX, nfs.TO:
		m.Cmdy("", m.Option(aaa.USERNAME)).CutTo(MAILBOX, arg[0])
	}
}
func (s Client) Login(m *ice.Message, arg ...string)  { s.Hash.Create(m, arg...); s.Sync(m) }
func (s Client) Logout(m *ice.Message, arg ...string) { s.Hash.Remove(m, arg...) }
func (s Client) Create(m *ice.Message, arg ...string) {
	s.client(m, m.Option(aaa.USERNAME), func(c *client.Client) {
		kit.If(!m.Warn(c.Create(m.Option(MAILBOX))), func() { s.Sync(m) })
	})
}
func (s Client) Remove(m *ice.Message, arg ...string) {
	s.client(m, m.Option(aaa.USERNAME), func(c *client.Client) {
		kit.If(!m.Warn(c.Delete(m.Option(MAILBOX))), func() {
			s.clear(m, m.Option(MAILBOX))
			s.Sync(m)
		})
	})
}
func (s Client) Rename(m *ice.Message, arg ...string) {
	s.client(m, m.Option(aaa.USERNAME), func(c *client.Client) {
		if !m.Warn(c.Rename(m.Option(MAILBOX), m.Option(nfs.TO))) {
			s.clear(m, m.Option(MAILBOX))
			s.Sync(m)
		}
	})
}
func (s Client) Prune(m *ice.Message, arg ...string) { s.Remove(m) }
func (s Client) Copy(m *ice.Message, arg ...string) {
	s.client(m, m.Option(aaa.USERNAME), func(c *client.Client) {
		c.Select(m.Option(MAILBOX), false)
		if !m.Warn(c.Copy(s.seqs(m, m.Option(mdb.ID)), m.Option(nfs.TO))) {
			s.Sync(m)
		}
	})
}
func (s Client) Move(m *ice.Message, arg ...string) {
	s.client(m, m.Option(aaa.USERNAME), func(c *client.Client) {
		c.Select(m.Option(MAILBOX), false)
		if !m.Warn(c.Move(s.seqs(m, m.Option(mdb.ID)), m.Option(nfs.TO))) {
			s.clear(m, m.Option(MAILBOX))
			s.Sync(m)
		}
	})
}
func (s Client) Spam(m *ice.Message, arg ...string)    { s.Move(m.Options(nfs.TO, SPAM)) }
func (s Client) Trash(m *ice.Message, arg ...string)   { s.Move(m.Options(nfs.TO, TRASH)) }
func (s Client) Archive(m *ice.Message, arg ...string) { s.Move(m.Options(nfs.TO, ARCHIVE)) }
func (s Client) Forword(m *ice.Message, arg ...string) {}
func (s Client) Reply(m *ice.Message, arg ...string)   {}
func (s Client) Send(m *ice.Message, arg ...string) {
	defer m.ToastProcess()()
	msg := s.Hash.List(m.Spawn(), m.Option(aaa.USERNAME))
	content := []byte(kit.JoinKV(nfs.DF, lex.NL, kit.Simple(FROM, m.Option(aaa.USERNAME), m.OptionSimple(TO, CC, SUBJECT), html.ContentType, "text/plain; charset=UTF-8")...) + lex.NL + lex.NL + m.Option(nfs.CONTENT))
	auth := smtp.PlainAuth("", msg.Append(aaa.USERNAME), msg.Append(aaa.PASSWORD), kit.Split(msg.Append(SMTP), ice.DF)[0])
	m.Logs(MAIL, string(content)).Warn(smtp.SendMail(msg.Append(SMTP), auth, msg.Append(aaa.USERNAME), kit.Split(m.Option(nfs.TO)), content))
	m.ProcessHold()
}
func (s Client) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		m.OptionFields(aaa.USERNAME)
		m.Cmdy(mdb.SELECT, m.PrefixKey(), "", mdb.HASH).PushAction(s.Sync, s.Logout).Action(s.Login, s.Sync).Display("").DisplayCSS("")
		kit.If(m.Length() == 0, func() { m.EchoInfoButton("please login", aaa.LOGIN) })
	} else if len(arg) == 1 {
		m.OptionFields(MAILBOX, mdb.ORDER, mdb.COUNT)
		m.Cmdy(mdb.SELECT, m.PrefixKey(), kit.KeyHash(arg[0]), mdb.HASH).Sort("order,mailbox")
		m.Table(func(value ice.Maps) {
			switch value[MAILBOX] {
			case INBOX:
				m.PushButton(s.Create)
			case SENT:
				m.PushButton(s.Send)
			case DRAFTS, SPAM, TRASH:
				m.PushButton(s.Prune)
			case ARCHIVE:
				m.PushButton(s.Create)
			default:
				m.PushButton(s.Create, s.Rename, s.Prune, s.Remove)
			}
		})
	} else if len(arg) == 2 {
		m.OptionDefault(mdb.CACHE_LIMIT, "-1")
		m.OptionFields("time,id,subject,from,to,cc,bcc,file")
		m.Cmdy(mdb.SELECT, m.PrefixKey(), s.keys(m, arg[0], arg[1]), mdb.LIST)
		m.PushAction(s.Forword, s.Reply, s.Copy, s.Move, s.Spam, s.Trash, s.Archive)
	} else {
		m.OptionFields(mdb.DETAIL)
		m.Cmdy(mdb.SELECT, m.PrefixKey(), s.keys(m, arg[0], arg[1]), mdb.LIST, mdb.ID, arg[2])
		m.EchoIFrame(m.Resource(m.Append(nfs.FILE)))
	}
}
func (s Client) Sync(m *ice.Message, arg ...string) {
	if m.Option(aaa.USERNAME) == "" {
		s.Hash.List(m.Spawn()).Table(func(value ice.Maps) { s.Sync(m.Spawn(value)) })
		return
	}
	kit.If(len(arg) == 0, func() { arg = append(arg, m.Option(aaa.USERNAME)) })
	s.client(m, arg[0], func(c *client.Client) {
		list := make(chan *imap.MailboxInfo, 10)
		m.Go(func() { c.List("", "*", list) })
		for item := range list {
			s.check(m, c, append(arg, item.Name)...)
		}
	})
}

func init() { ice.Cmd(prefix(), Client{}) }

func (s Client) client(m *ice.Message, username string, cb func(c *client.Client)) {
	msg := s.Hash.List(m.Spawn(), username)
	c, err := client.DialTLS(msg.Append(IMAP), nil)
	if m.Warn(err) {
		return
	}
	defer c.Close()
	if m.Warn(c.Login(msg.Append(aaa.USERNAME), msg.Append(aaa.PASSWORD))) {
		return
	}
	defer c.Logout()
	cb(c)
}
func (s Client) check(m *ice.Message, c *client.Client, arg ...string) {
	mbox, err := c.Select(arg[1], true)
	if m.Warn(err) {
		return
	}
	s.order(m, arg...)
	from, to := uint32(kit.Int(mdb.Conf(m.Message, m.PrefixKey(), s.keys(m, arg[0], arg[1], mdb.META, mdb.COUNT)))+1), mbox.Messages
	if from > to {
		return
	}
	m.Debug("sync %s %s %v-%v", arg[0], arg[1], from, to)
	list := make(chan *imap.Message, 10)
	m.ToastProcess()
	m.Go(func() {
		c.Fetch(s.seqs(m, kit.Format(from), kit.Format(to)), []imap.FetchItem{imap.FetchRFC822, imap.FetchRFC822Header, imap.FetchInternalDate}, list)
	})
	s.fetch(m, list, from, to, arg...)
}
func (s Client) fetch(m *ice.Message, list chan *imap.Message, from, to uint32, arg ...string) {
	m.GoToast(func(toast func(string, int, int)) {
		toast(arg[1], 0, int(to-from+1))
		for item := range list {
			var section imap.BodySectionName
			r, err := mail.CreateReader(item.GetBody(&section))
			if m.Warn(err) {
				continue
			}
			for fields := r.Header.Fields(); fields.Next(); {
				m.Debug("header %v %v", fields.Key(), fields.Value())
			}
			date, _ := r.Header.Date()
			subject, _ := r.Header.Subject()
			args := []ice.Any{mdb.TIME, date.Format(kit.MOD_TIME), mdb.ID, item.SeqNum, SUBJECT, subject}
			for _, key := range []string{FROM, TO, CC, BCC, SENDER, "Reply-To"} {
				if list, err := r.Header.AddressList(key); !m.Warn(err) {
					_list := []string{}
					for _, item := range list {
						_list = append(_list, kit.Format("%s<%s>", s.decode(m, item.Name), item.Address))
					}
					args = append(args, strings.ToLower(key), kit.Join(_list))
				}
			}
			arg := append(arg, kit.Format(item.SeqNum))
			toast(arg[1]+" "+subject, int(item.SeqNum-from), int(to-from+1))
			msg, file, trans := m.Spawn(), s.dirs(m, arg[0], arg[1], arg[2], CONTENT_HTML), []func(string) string{}
			for {
				p, err := r.NextPart()
				if err == io.EOF || m.Warn(err) {
					break
				}
				switch h := p.Header.(type) {
				case *mail.InlineHeader:
					msg, trans = s.parse(msg, file, h.Header.Map(), p.Body, trans)
				case *mail.AttachmentHeader:
					msg, trans = s.parse(msg, file, h.Header.Map(), p.Body, trans)
				default:
					m.ErrorNotImplement(h)
				}
			}
			args = append(args, nfs.FILE, m.Cmdx(nfs.SAVE, file, kit.Dict(nfs.CONTENT, s.trans(msg, trans))))
			m.Cmd(mdb.INSERT, m.PrefixKey(), s.keys(m, arg[0], arg[1]), mdb.LIST, args)
		}
	})
}
func (s Client) order(m *ice.Message, arg ...string) {
	mdb.Conf(m.Message, m.PrefixKey(), s.keys(m, arg[0], arg[1], mdb.META, MAILBOX), arg[1])
	mdb.Conf(m.Message, m.PrefixKey(), s.keys(m, arg[0], arg[1], mdb.META, mdb.ORDER), kit.Select("100", map[string]string{
		INBOX: "1", DRAFTS: "2", SENT: "3", SPAM: "4", TRASH: "5", ARCHIVE: "6",
	}[arg[1]]))
}
func (s Client) clear(m *ice.Message, box string) {
	mdb.Conf(m.Message, m.PrefixKey(), s.keys(m, m.Option(aaa.USERNAME), box), "")
}
func (s Client) seqs(m *ice.Message, arg ...string) *imap.SeqSet {
	seqs := &imap.SeqSet{}
	if len(arg) == 1 {
		seqs.AddNum(uint32(kit.Int(arg[0])))
	} else if len(arg) == 2 {
		seqs.AddRange(uint32(kit.Int(arg[0])), uint32(kit.Int(arg[1])))
	}
	return seqs
}
func (s Client) keys(m *ice.Message, username, mailbox string, arg ...string) string {
	return kit.KeyHash(username, kit.KeyHash(mailbox), arg)
}
func (s Client) dirs(m *ice.Message, arg ...string) string {
	return path.Join(USR_LOCAL_EMAIL, path.Join(arg...))
}
