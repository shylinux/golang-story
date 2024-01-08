package mail

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
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/icebergs/base/web/html"
	kit "shylinux.com/x/toolkits"
)

const (
	MAIL           = "mail"
	SMTP           = "smtp"
	IMAP           = "imap"
	MAILBOX        = "mailbox"
	SUBJECT        = "subject"
	CONTENT_HTML   = "content.html"
	USR_LOCAL_MAIL = "usr/local/mail/"

	INBOX   = "INBOX"
	SENT    = "Sent"
	DRAFTS  = "Drafts"
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
	ice.Code
	ice.Hash
	short   string `data:"username"`
	field   string `data:"time,username,password,smtp,imap"`
	login   string `name:"login username* password* smtp* imap*" help:"登录"`
	logout  string `name:"logout" help:"退出"`
	create  string `name:"create mailbox*"`
	rename  string `name:"rename mailbox* to*" help:"重命名"`
	spam    string `name:"spam" help:"垃圾"`
	archive string `name:"archive" help:"归档"`
	move    string `name:"move to*" help:"移动到"`
	copy    string `name:"copy to*" help:"复制到"`
	send    string `name:"send to* cc subject* content*:textarea" help:"发送"`
	reply   string `name:"send subject* content*:textarea" help:"回复"`
	forword string `name:"forword to*" help:"转发"`
	list    string `name:"list refresh" help:"邮箱" icon:"Mail.png"`
	sync    string `name:"sync" help:"同步"`
}

func (s Client) Inputs(m *ice.Message, arg ...string) {
	switch s.Hash.Inputs(m, arg...); arg[0] {
	case SMTP:
		m.Push(arg[0], kit.Keys(arg[0], kit.Select("shylinux.com", kit.Split(m.Option(aaa.USERNAME)), 1))+":25")
	case IMAP:
		m.Push(arg[0], kit.Keys(arg[0], kit.Select("shylinux.com", kit.Split(m.Option(aaa.USERNAME)), 1))+":993")
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
		kit.If(!m.Warn(c.Delete(m.Option(MAILBOX))), func() { s.Sync(m) })
		mdb.Conf(m, m.PrefixKey(), kit.KeyHash(m.Option(aaa.USERNAME), kit.KeyHash(m.Option(MAILBOX))), "")
	})
}
func (s Client) Rename(m *ice.Message, arg ...string) {
	s.client(m, m.Option(aaa.USERNAME), func(c *client.Client) {
		if !m.Warn(c.Rename(m.Option(MAILBOX), m.Option(nfs.TO))) {
			mdb.Conf(m, m.PrefixKey(), kit.KeyHash(m.Option(aaa.USERNAME), kit.KeyHash(m.Option(MAILBOX))), "")
			s.Sync(m)
		}
	})
}
func (s Client) Prune(m *ice.Message, arg ...string)   { s.Remove(m) }
func (s Client) Spam(m *ice.Message, arg ...string)    { m.Option(nfs.TO, SPAM); s.Move(m) }
func (s Client) Trash(m *ice.Message, arg ...string)   { m.Option(nfs.TO, TRASH); s.Move(m) }
func (s Client) Archive(m *ice.Message, arg ...string) { m.Option(nfs.TO, ARCHIVE); s.Move(m) }
func (s Client) Move(m *ice.Message, arg ...string) {
	s.client(m, m.Option(aaa.USERNAME), func(c *client.Client) {
		c.Select(m.Option(MAILBOX), false)
		set := &imap.SeqSet{}
		set.AddNum(uint32(kit.Int(m.Option(mdb.ID))))
		if !m.Warn(c.Move(set, m.Option(nfs.TO))) {
			mdb.Conf(m, m.PrefixKey(), kit.KeyHash(m.Option(aaa.USERNAME), kit.KeyHash(m.Option(MAILBOX))), "")
			mdb.Conf(m, m.PrefixKey(), kit.KeyHash(m.Option(aaa.USERNAME), kit.KeyHash(m.Option(nfs.TO))), "")
			s.Sync(m)
		}
	})
}
func (s Client) Copy(m *ice.Message, arg ...string) {
	s.client(m, m.Option(aaa.USERNAME), func(c *client.Client) {
		c.Select(m.Option(MAILBOX), false)
		set := &imap.SeqSet{}
		set.AddNum(uint32(kit.Int(m.Option(mdb.ID))))
		if !m.Warn(c.Copy(set, m.Option(nfs.TO))) {
			mdb.Conf(m, m.PrefixKey(), kit.KeyHash(m.Option(aaa.USERNAME), kit.KeyHash(m.Option(nfs.TO))), "")
			s.Sync(m)
		}
	})
}
func (s Client) Send(m *ice.Message, arg ...string) {
	msg := s.Hash.List(m.Spawn(), m.Option(aaa.USERNAME))
	content := []byte(kit.JoinKV(nfs.DF, lex.NL, kit.Simple(FROM, m.Option(aaa.USERNAME), m.OptionSimple(TO, CC, SUBJECT), html.ContentType, "text/plain; charset=UTF-8")...) + lex.NL + lex.NL + m.Option(nfs.CONTENT))
	auth := smtp.PlainAuth("", msg.Append(aaa.USERNAME), msg.Append(aaa.PASSWORD), kit.Split(msg.Append(SMTP), ice.DF)[0])
	m.Logs(MAIL, string(content)).Warn(smtp.SendMail(msg.Append(SMTP), auth, msg.Append(aaa.USERNAME), kit.Split(m.Option(nfs.TO)), content))
	web.ToastSuccess(m.Message)
	m.ProcessHold()
}
func (s Client) Reply(m *ice.Message, arg ...string)   {}
func (s Client) Forword(m *ice.Message, arg ...string) {}
func (s Client) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		m.OptionFields(aaa.USERNAME)
		m.Cmdy(mdb.SELECT, m.PrefixKey(), "", mdb.HASH).PushAction(s.Sync, s.Logout).Action(s.Login, s.Sync).Display("")
	} else if len(arg) == 1 {
		m.OptionFields(MAILBOX, mdb.COUNT, mdb.ORDER)
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
		m.OptionFields("time,id,subject,from,to,cc,bcc,file")
		m.Cmdy(mdb.SELECT, m.PrefixKey(), s.keys(m, arg[0], arg[1]), mdb.LIST)
		m.PushAction(s.Reply, s.Forword, s.Copy, s.Move, s.Spam, s.Trash, s.Archive)
	} else {
		m.OptionFields(mdb.DETAIL)
		m.Cmdy(mdb.SELECT, m.PrefixKey(), s.keys(m, arg[0], arg[1]), mdb.LIST, mdb.ID, arg[2]).Display("")
		if strings.HasPrefix(m.Append(html.ContentType), html.TEXT_PLAIN) {
			if p := m.Cmdx(nfs.CAT, m.Append(nfs.FILE)); kit.HasPrefix(p, nfs.PS, web.HTTP) {
				m.EchoIFrame(kit.MergeURL(p, ice.POD, m.Option(ice.MSG_USERPOD)))
			} else {
				m.Echo(p)
			}
		} else {
			m.EchoIFrame(kit.MergeURL(path.Join(nfs.REQUIRE, m.Append(nfs.FILE)), ice.POD, m.Option(ice.MSG_USERPOD)))
		}
	}
}
func (s Client) Sync(m *ice.Message, arg ...string) {
	// kit.If(len(arg) == 0, func() { arg = append(arg, m.Option(aaa.USERNAME)) })
	s.Hash.List(m, arg...)
	c, err := client.DialTLS(m.Append(IMAP), nil)
	if m.Warn(err) {
		return
	}
	defer c.Close()
	if m.Warn(c.Login(m.Append(aaa.USERNAME), m.Append(aaa.PASSWORD))) {
		return
	}
	defer c.Logout()
	m.SetAppend()
	done := make(chan error, 1)
	list := make(chan *imap.MailboxInfo, 10)
	m.Go(func() { done <- c.List("", "*", list) })
	for item := range list {
		m.Push(MAILBOX, item.Name)
	}
	if m.Warn(<-done) {
		return
	}
	m.ProcessRefresh()
	defer web.ToastSuccess(m.Message)
	m.Table(func(value ice.Maps) {
		arg := append(arg, value[MAILBOX])
		mbox, err := c.Select(arg[1], true)
		if m.Warn(err) {
			return
		}
		s.order(m, arg...)
		from, to := uint32(kit.Int(mdb.Conf(m, m.PrefixKey(), s.keys(m, arg[0], arg[1], mdb.META, mdb.COUNT)))+1), mbox.Messages
		if from > to {
			return
		}
		m.Debug("sync %s %s %v-%v", arg[0], arg[1], from, to)
		seqset := &imap.SeqSet{}
		seqset.AddRange(from, to)
		done := make(chan error, 1)
		list := make(chan *imap.Message, 10)
		m.Go(func() {
			done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchRFC822, imap.FetchRFC822Header, imap.FetchInternalDate}, list)
		})
		web.GoToast(m.Message, kit.Select(m.CommandKey()+lex.SP+m.ActionKey()), func(toast func(string, int, int)) []string {
			toast(arg[1], 0, int(to-from+1))
			for item := range list {
				var section imap.BodySectionName
				mr, err := mail.CreateReader(item.GetBody(&section))
				if m.Warn(err) {
					continue
				}
				contentType := ""
				date, _ := mr.Header.Date()
				fields := mr.Header.Fields()
				for fields.Next() {
					m.Debug("header %v %v", fields.Key(), fields.Value())
					kit.If(fields.Key() == html.ContentType, func() { contentType = fields.Value() })
				}
				subject, _ := mr.Header.Subject()
				arg := append(arg, kit.Format(item.SeqNum))
				toast(arg[1]+" "+subject, int(item.SeqNum-from), int(to-from+1))
				args := []ice.Any{mdb.TIME, date.Format(kit.MOD_TIME), mdb.ID, item.SeqNum, SUBJECT, subject}
				for _, key := range []string{FROM, TO, CC, BCC, SENDER, "Reply-To"} {
					list, err := mr.Header.AddressList(key)
					if m.Warn(err) {
						continue
					}
					_list := []string{}
					for _, item := range list {
						_list = append(_list, kit.Format("%s<%s>", item.Name, item.Address))
					}
					args = append(args, strings.ToLower(key), kit.Join(_list))
				}
				for {
					p, err := mr.NextPart()
					if err == io.EOF || m.Warn(err) {
						break
					}
					switch h := p.Header.(type) {
					case *mail.InlineHeader:
						args = append(args, nfs.FILE, m.Cmdx(nfs.SAVE, s.dirs(m, arg[0], arg[1], arg[2], "content."+path.Base(kit.Split(contentType, ";")[0])), kit.Dict(nfs.CONTENT, p.Body)))
					case *mail.AttachmentHeader:
						if filename, err := h.Filename(); m.Warn(err) || filename == "" {
							break
						} else {
							head := map[string]string{}
							for fields := h.Fields(); fields.Next(); {
								head[fields.Key()] = fields.Value()
								args = append(args, kit.KeyExtra(filename, fields.Key()), fields.Value())
							}
							args = append(args, kit.KeyExtra(filename, nfs.PATH), m.Cmdx(nfs.SAVE, s.dirs(m, arg[0], arg[1], arg[2], filename), kit.Dict(nfs.CONTENT, p.Body)))
						}
					default:
						m.Warn("unknown header")
					}
				}
				m.Cmd(mdb.INSERT, m.PrefixKey(), s.keys(m, arg[0], arg[1]), mdb.LIST, args)
			}
			return nil
		})
		m.Warn(<-done)
	})
}

func init() { ice.ChatCtxCmd(Client{}) }

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
func (s Client) dirs(m *ice.Message, arg ...string) string {
	return path.Join(USR_LOCAL_MAIL, path.Join(arg...))
}
func (s Client) keys(m *ice.Message, username, mailbox string, arg ...string) string {
	return kit.KeyHash(username, kit.KeyHash(mailbox), arg)
}
func (s Client) order(m *ice.Message, arg ...string) {
	mdb.Conf(m, m.PrefixKey(), s.keys(m, arg[0], arg[1], mdb.META, MAILBOX), arg[1])
	mdb.Conf(m, m.PrefixKey(), s.keys(m, arg[0], arg[1], mdb.META, mdb.ORDER), kit.Select("100", map[string]string{INBOX: "1", SENT: "2", DRAFTS: "3", SPAM: "4", TRASH: "5", ARCHIVE: "6"}[arg[1]]))
}
