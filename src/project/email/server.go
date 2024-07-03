package email

import (
	"shylinux.com/x/ice"
)

type server struct {
	ice.Hash
	list string `name:"list hash auto" help:"服务器"`
}

func (s server) List(m *ice.Message, arg ...string) {
	m.Echo(m.System("maddy", "version").Result())
}

func init() { ice.Cmd("web.chat.email.server", server{}) }
