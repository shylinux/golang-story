package email

import (
	"fmt"

	"github.com/flashmob/go-guerrilla"
	"shylinux.com/x/ice"
)

type server struct {
	ice.Hash

	list string `name:"list hash auto" help:"server"`
}

func (s server) List(m *ice.Message, arg ...string) {
	d := guerrilla.Daemon{}
	err := d.Start()
	m.Debug("what %v", err)
	if err == nil {
		fmt.Println("Server Started!")
	}
}

func init() { ice.Cmd("web.chat.project.email.server", server{}) }
