package docker

import (
	"os"
	"path"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/tcp"
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/icebergs/core/code"
	kit "shylinux.com/x/toolkits"
)

const (
	DOCKER    = "docker"
	IMAGE     = "image"
	CONTAINER = "container"

	BUILD   = "build"
	PULL    = "pull"
	EXEC    = "exec"
	STOP    = "stop"
	RUN     = "run"
	LS      = "ls"
	PS      = "ps"
	RM      = "rm"
	KILL    = "kill"
	PRUNE   = "prune"
	RENAME  = "rename"
	RESTART = "restart"

	IMAGE_ID     = "IMAGE_ID"
	CONTAINER_ID = "CONTAINER_ID"
)

type client struct {
	ice.Code
	build   string `name:"build name=contexts dir=usr/publish/" help:"构建"`
	pull    string `name:"pull image=alpine" help:"下载"`
	start   string `name:"start cmd dev port" help:"启动"`
	restart string `name:"restart" help:"重启"`
	serve   string `name:"serve arg" help:"服务"`
	drop    string `name:"drop" help:"删除"`
	prune   string `name:"prune" help:"清理"`
	list    string `name:"list IMAGE_ID CONTAINER_ID auto" help:"容器"`
	df      string `name:"df" help:"磁盘"`
	open    string `name:"open" help:"系统"`
}

func (s client) envhost(m *ice.Message) bool {
	return nfs.ExistsFile(m.Message, strings.TrimPrefix(s.host(m), "unix://")) && s.host(m) != os.Getenv("DOCKER_HOST")
}
func (s client) host(m *ice.Message) string {
	return "unix://" + kit.Path("usr/install/docker/docker.sock")
}
func (s client) cmd(m *ice.Message, p string) string {
	if s.envhost(m) {
		return kit.Format("docker --host %s exec -w /root -it %s sh", s.host(m), p)
	}
	return kit.Format("docker exec -w /root -it %s sh", p)
}
func (s client) docker(m *ice.Message, arg ...string) string {
	if s.envhost(m) {
		return m.Cmdx(cli.SYSTEM, DOCKER, "--host", s.host(m), arg)
	}
	return m.Cmdx(cli.SYSTEM, DOCKER, arg)
}
func (s client) image(m *ice.Message, arg ...string) string {
	return s.docker(m, kit.Simple(IMAGE, arg)...)
}
func (s client) container(m *ice.Message, arg ...string) string {
	return s.docker(m, kit.Simple(CONTAINER, arg)...)
}

func (s client) Inputs(m *ice.Message, arg ...string) {
	switch arg[0] {
	case IMAGE:
		m.Push(arg[0], cli.BUSYBOX, cli.ALPINE, cli.CENTOS, cli.UBUNTU)
	case ice.DEV:
		u := web.OptionUserWeb(m)
		m.Push(arg[0], tcp.PublishLocalhost(m.Message, u.Scheme+"://"+u.Hostname()+ice.DF+u.Port()))
		m.Cmd(web.SPIDE).Tables(func(value ice.Maps) { m.Push(arg[0], value[web.CLIENT_URL]) })
	case tcp.PORT:
		s.List(m.Spawn(), "hi").Tables(func(value ice.Maps) {
			ls := strings.SplitN(value["PORTS"], "->", 2)
			if len(ls) > 1 {
				m.Push(tcp.PORT, kit.Slice(strings.Split(ls[0], ice.DF), -1)[0])
			} else {
				m.Push(tcp.PORT, "")
			}
			m.Push(ice.BIN, value["COMMAND"])
		})
		nets := m.Cmdx(cli.SYSTEM, "netstat", "-tln")
		port := kit.Int(m.Cmdx(tcp.PORT, tcp.CURRENT)) + 1
		for strings.Contains(nets, kit.Format(":%v ", port)) {
			port += 1
		}
		m.Push(tcp.PORT, port).Push(ice.BIN, "").SortIntR(arg[0])
		m.AppendTrans(func(value string, key string, index int) string {
			return value + kit.Select("", ":9020", index == 0 && key == "port")
		})
	}
}
func (s client) Build(m *ice.Message, arg ...string) {
	defer s.Code.ToastProcess(m)()
	s.Code.Module(m, path.Join(m.Option(nfs.DIR), "Dockerfile"), _dockerfile)
	s.image(m, BUILD, "-t", m.Option(mdb.NAME), m.Option(nfs.DIR))
}
func (s client) Pull(m *ice.Message, arg ...string) {
	s.image(m, PULL, m.Option(IMAGE))
}
func (s client) Start(m *ice.Message, arg ...string) {
	args := []string{"-e", "LANG=en_US.UTF-8"}
	if m.Option(ice.CMD) == "" {
		if m.Option(ice.DEV) != "" && strings.Contains(m.Option(ice.DEV), ice.PT) {
			args = append(args, "-e", "ctx_dev="+m.Option(ice.DEV))
		} else {
			args = append(args, "-e", "ctx_dev="+m.Option(ice.MSG_USERHOST))
		}
		if m.Option(tcp.PORT) != "" && strings.Contains(m.Option(tcp.PORT), ice.DF) {
			args = append(args, "-p", m.Option(tcp.PORT))
		}
		// m.Option(CONTAINER_ID, s.container(m, kit.Simple(RUN, "-e", "ctx_dev="+m.Option(ice.DEV), args, "-dt", m.Option(IMAGE_ID))...))
		m.Option(CONTAINER_ID, s.container(m, kit.Simple(RUN, args, "-dt", m.Option(IMAGE_ID))...))
	} else {
		m.Option(CONTAINER_ID, s.container(m, kit.Simple(RUN, args, "-dt", m.Option(IMAGE_ID), m.Option(ice.CMD))...))
	}
}
func (s client) Restart(m *ice.Message, arg ...string) {
	s.container(m, RESTART, m.Option(CONTAINER_ID))
	s.Serve(m)
}
func (s client) Serve(m *ice.Message, arg ...string) {
	s.container(m, EXEC, m.Option(CONTAINER_ID), "wget", "-O", "/root/index.sh", m.Option(ice.MSG_USERHOST))
	s.container(m, EXEC, "-w", "/root", "-e", "ctx_dev="+m.Option(ice.MSG_USERHOST), "-dt", m.Option(CONTAINER_ID), nfs.SH, "/root/index.sh", "app")
}
func (s client) Stop(m *ice.Message, arg ...string) {
	defer s.Code.ToastProcess(m)()
	if m.Option("PID") != "" {
		s.container(m, EXEC, m.Option(CONTAINER_ID), KILL, m.Option("PID"))
	} else {
		s.container(m, STOP, m.Option(CONTAINER_ID))
	}
}
func (s client) Drop(m *ice.Message, arg ...string) {
	if m.Option(CONTAINER_ID) != "" {
		s.container(m, RM, m.Option(CONTAINER_ID))
	} else if m.Option(IMAGE_ID) != "" {
		s.image(m, RM, m.Option(IMAGE_ID))
	}
}
func (s client) Prune(m *ice.Message, arg ...string) {
	if len(arg) > 0 {
		m.Echo(s.container(m, PRUNE, "-f"))
	} else {
		m.Echo(s.image(m, PRUNE, "-f"))
	}
}
func (s client) Modify(m *ice.Message, arg ...string) {
	switch arg[0] {
	case "NAMES":
		s.container(m, RENAME, m.Option(CONTAINER_ID), arg[1])
	}
}
func (s client) List(m *ice.Message, arg ...string) *ice.Message {
	if len(arg) < 1 || arg[0] == "" {
		m.SplitIndex(strings.Replace(s.image(m, LS), "IMAGE ID", IMAGE_ID, 1))
		m.Cut("CREATED,IMAGE_ID,SIZE,REPOSITORY,TAG")
		m.PushAction(s.Drop).Action(s.Build, s.Pull, s.Df, s.Prune)
		m.StatusTimeCount("SIZE", s.Df(m.Spawn()).Append("SIZE"))
	} else if len(arg) < 2 || arg[1] == "" {
		m.SplitIndex(strings.Replace(s.container(m, LS, "-a"), "CONTAINER ID", CONTAINER_ID, 1)).RenameAppend("IMAGE", "REPOSITORY")
		m.Cut("CREATED,CONTAINER_ID,REPOSITORY,COMMAND,PORTS,STATUS,NAMES")
		m.Tables(func(value ice.Maps) {
			if strings.HasPrefix(value["STATUS"], "Up") {
				m.PushButton(s.Open, s.Vimer, s.Xterm, s.Stop)
			} else {
				m.PushButton(s.Restart, s.Drop)
			}
		}).Action(s.Start, s.Prune)
		m.EchoScript(strings.Replace(s.cmd(m, arg[0]), EXEC, RUN, 1))
		m.Cmdy(code.PUBLISH, ice.CONTEXTS, ice.MISC)
		m.StatusTimeCount("SIZE", s.Df(m.Spawn()).Appendv("SIZE")[1])
	} else if len(arg) < 3 || arg[2] == "" {
		m.SplitIndex(s.container(m, EXEC, arg[1], PS)).PushAction(s.Stop).Action(s.Open, s.Vimer, s.Xterm, s.Serve)
		m.EchoScript(s.cmd(m, arg[1]))
		m.StatusTimeCount()
	} else {
		m.Echo(s.container(m, kit.Simple(EXEC, arg[1], kit.Split(arg[2], ice.SP, ice.SP))...))
		m.StatusTimeCount()
	}
	return m
}
func (s client) Df(m *ice.Message, arg ...string) *ice.Message {
	m.SplitIndex(s.docker(m, cli.SYSTEM, "df"))
	return m
}
func (s client) Open(m *ice.Message, arg ...string) {
	s.Code.Iframe(m, "系统页", web.MergePod(m, kit.Select(m.Option(CONTAINER_ID), arg, 1)), arg...)
}
func (s client) Vimer(m *ice.Message, arg ...string) {
	s.Code.Iframe(m, "编辑器", web.MergePodCmd(m, m.Option(CONTAINER_ID), web.CODE_VIMER), arg...)
}
func (s client) Xterm(m *ice.Message, arg ...string) {
	s.Code.Xterm(m, []string{mdb.TYPE, s.cmd(m, kit.Select(m.Option(CONTAINER_ID), arg, 1)), mdb.NAME, m.Option(CONTAINER_ID)}, arg...)
}

func init() { ice.CodeCtxCmd(client{}) }

const _dockerfile = `
FROM alpine

WORKDIR /root/contexts
COPY ice.linux.amd64 bin/ice.bin

CMD ./bin/ice.bin forever start dev dev
`
