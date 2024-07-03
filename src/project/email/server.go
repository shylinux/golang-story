package email

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/nfs"
)

type server struct {
	list string `name:"list hash auto" help:"服务器"`
}

func (s server) List(m *ice.Message, arg ...string) {
	m.Echo(m.System(MADDY, nfs.VERSION).Result())
}

func init() { ice.Cmd("web.chat.email.server", server{}) }
