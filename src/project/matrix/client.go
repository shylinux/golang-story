package matrix

import (
	"os"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

const (
	LIST   = "list"
	INFO   = "info"
	DATA   = "data"
	TOTAL  = "total"
	IDNAME = "idname"
)

type Client struct{ ice.Code }

func (s Client) Create(m *ice.Message, arg ...string) {
	s.Cmdx(m, "", "", arg...)
}
func (s Client) Remove(m *ice.Message, arg ...string) {
	s.Cmdx(m, "", "", m.OptionSimple(m.Config(IDNAME))...)
}
func (s Client) List(m *ice.Message, arg ...string) {
	if len(arg) == 0 {
		res := s.Cmdx(m, "", LIST)
		kit.For(kit.Value(res, DATA), func(value ice.Map) { m.Push("", value) })
		m.Action(s.Create).StatusTimeCountTotal(kit.Value(res, TOTAL))
	} else {
		res := s.Cmdx(m, "", INFO, m.Config(IDNAME), arg[0])
		m.Push(mdb.DETAIL, kit.Value(res, DATA))
	}
	m.PushAction(mdb.REMOVE)
}
func (s Client) Cmdx(m *ice.Message, mod, cmd string, arg ...string) ice.Any {
	kit.If(cmd == "", func() { cmd = m.ActionKey() })
	kit.If(mod == "", func() { mod = m.CommandKey() })
	msg := s.Code.System(m.Spawn(), kit.Select(USR_MATRIX, os.Getenv("MATRIX_HOME")), kit.Simple(BIN_MATRIX, mod, cmd, arg)...)
	if !cli.IsSuccess(msg.Message) {
		web.ToastFailure(m.Message, msg.Append(cli.CMD_ERR))
		m.ProcessHold()
	}
	if !kit.HasPrefix(msg.Append(cli.CMD_OUT), "{") {
		web.ToastFailure(m.Message, msg.Append(cli.CMD_OUT))
		m.ProcessHold()
	}
	return kit.UnMarshal(msg.Append(cli.CMD_OUT))
}
