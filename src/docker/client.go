package docker

import (
	"path"
	"runtime"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/ssh"
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
	REPOSITORY   = "REPOSITORY"
	TAG          = "TAG"
)

type client struct {
	ice.Code
	build   string `name:"build name=contexts dir=usr/publish/" help:"构建"`
	pull    string `name:"pull image=alpine" help:"下载"`
	imports string `name:"imports path*=usr/publish/alpine-dev.tar image*=alpine-dev" help:"导入"`
	exports string `name:"exports name=alpine-dev" help:"导出"`
	start   string `name:"start cmd dev port" help:"启动"`
	restart string `name:"restart" help:"重启"`
	serve   string `name:"serve arg" help:"服务"`
	save    string `name:"save name" help:"导出"`
	drop    string `name:"drop" help:"删除"`
	prune   string `name:"prune" help:"清理"`
	list    string `name:"list IMAGE_ID CONTAINER_ID auto" help:"容器"`
	df      string `name:"df" help:"磁盘"`
	open    string `name:"open" help:"系统"`
}

func (s client) host(m *ice.Message) []string {
	if nfs.ExistsFile(m.Message, "usr/install/docker/docker.sock") {
		return []string{"--host", "unix://" + kit.Path("usr/install/docker/docker.sock")}
	}
	return nil
}
func (s client) cmd(m *ice.Message, p string) string {
	return strings.Join(kit.Simple(DOCKER, s.host(m), EXEC, "-w", "/root", "-it", p, code.SH), ice.SP)
}
func (s client) docker(m *ice.Message, arg ...string) string {
	return m.Cmdx(cli.SYSTEM, DOCKER, s.host(m), arg)
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
		m.Push(arg[0], m.Option(ice.MSG_USERHOST)).Cmd(web.SPIDE, func(value ice.Maps) { m.Push(arg[0], value[web.CLIENT_URL]) })
	case ice.CMD:
		m.Push(arg[0], code.BASH, code.SH)
	case tcp.PORT:
		port := kit.Int(m.Cmdx(tcp.PORT, tcp.CURRENT)) + 1
		nets := m.Cmdx(cli.SYSTEM, "netstat", "-ln", kit.Split(kit.Select("-t", "-p tcp", runtime.GOOS == cli.DARWIN)))
		for strings.Contains(nets, kit.Format(":%v ", port)) {
			port += 1
		}
		m.Push(tcp.PORT, kit.Format("%d:9020", port)).Push(ice.BIN, "")
		m.Push(tcp.PORT, "20000:9020").Push(ice.BIN, "")
		m.Push(tcp.PORT, "20001:9020").Push(ice.BIN, "")
		s.List(m.Spawn(), cli.ALPINE).Tables(func(value ice.Maps) {
			if ls := strings.SplitN(value["PORTS"], "->", 2); len(ls) > 1 {
				m.Push(tcp.PORT, kit.Slice(strings.Split(ls[0], ice.DF), -1)[0])
				m.Push(ice.BIN, value["COMMAND"])
			}
		})
	case nfs.PATH:
		m.Option(nfs.DIR_REG, kit.ExtReg(nfs.TAR))
		if p := kit.Select(ice.USR_PUBLISH, arg, 1); strings.HasSuffix(p, ice.PS) {
			m.Cmdy(nfs.DIR, p, nfs.PATH)
		} else {
			m.Cmdy(nfs.DIR, path.Dir(p), nfs.PATH)
		}
	}
}
func (s client) Build(m *ice.Message, arg ...string) {
	defer s.Code.ToastProcess(m)()
	s.Code.Module(m, path.Join(m.Option(nfs.DIR), "Dockerfile"), nfs.Template(m, "DockerFile"))
	s.image(m, BUILD, "-t", m.Option(mdb.NAME), m.Option(nfs.DIR))
}
func (s client) Imports(m *ice.Message, arg ...string) {
	s.docker(m, mdb.IMPORT, m.Option(nfs.PATH), m.Option(IMAGE))
}
func (s client) Exports(m *ice.Message, arg ...string) {
	s.docker(m, mdb.EXPORT, m.Option(CONTAINER_ID), "-o", path.Join(ice.USR_PUBLISH, kit.Keys(m.Option(mdb.NAME), nfs.TAR)))
}
func (s client) Pull(m *ice.Message, arg ...string) {
	s.image(m, PULL, m.Option(IMAGE))
}
func (s client) Save(m *ice.Message, arg ...string) {
	s.image(m, nfs.SAVE, m.Option(IMAGE_ID), "-o", path.Join(ice.USR_PUBLISH, kit.Keys(kit.Select(m.Option(REPOSITORY), m.Option(mdb.NAME)), nfs.TAR)))
}
func (s client) Start(m *ice.Message, arg ...string) {
	image := m.Option(IMAGE_ID)
	if m.Option(REPOSITORY) != "" && m.Option(TAG) != "" {
		image = m.Option(REPOSITORY) + ice.DF + m.Option(TAG)
		defer func() { m.ProcessRewrite(IMAGE_ID, m.Option(IMAGE_ID), CONTAINER_ID, s.short(m)) }()
	}
	if args := kit.Simple("-e", "LANG=en_US.UTF-8"); m.Option(ice.CMD) == "" {
		if m.Option(ice.DEV) != "" && strings.Contains(m.Option(ice.DEV), ice.PT) {
			args = append(args, "-e", "ctx_dev="+m.Option(ice.DEV))
		} else {
			args = append(args, "-e", "ctx_dev="+m.Option(ice.MSG_USERHOST))
		}
		if m.Option(tcp.PORT) != "" && strings.Contains(m.Option(tcp.PORT), ice.DF) {
			args = append(args, "-p", m.Option(tcp.PORT))
		}
		m.Option(CONTAINER_ID, s.container(m, kit.Simple(RUN, args, "-dt", image)...))
	} else {
		m.Option(CONTAINER_ID, s.container(m, kit.Simple(RUN, args, "-dt", image, m.Option(ice.CMD))...))
	}
}
func (s client) Restart(m *ice.Message, arg ...string) {
	s.container(m, RESTART, m.Option(CONTAINER_ID))
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
func (s client) Search(m *ice.Message, arg ...string) {
	if arg[0] == mdb.FOREACH && (arg[1] == "" || arg[1] == ssh.SHELL) {
		s.List(m.Spawn(ice.Maps{ice.MSG_FIELDS: ""})).Tables(func(value ice.Maps) {
			m.PushSearch(mdb.TYPE, ssh.SHELL, mdb.NAME, value[REPOSITORY]+ice.DF+value[TAG],
				mdb.TEXT, "docker run -w /root -it "+value[REPOSITORY]+ice.DF+value[TAG]+" sh")
		})
	}
}
func (s client) List(m *ice.Message, arg ...string) *ice.Message {
	if len(arg) < 1 || arg[0] == "" {
		m.SplitIndex(strings.Replace(s.image(m, LS), "IMAGE ID", IMAGE_ID, 1))
		m.Cut("CREATED,IMAGE_ID,SIZE,REPOSITORY,TAG")
		m.PushAction(s.Start, s.Save, s.Drop).Action(s.Build, s.Imports, s.Pull, s.Df, s.Prune)
		m.StatusTimeCount("SIZE", s.Df(m.Spawn()).Append("SIZE"))
	} else if len(arg) < 2 || arg[1] == "" {
		m.SplitIndex(strings.Replace(s.container(m, LS, "-a"), "CONTAINER ID", CONTAINER_ID, 1)).RenameAppend("IMAGE", "REPOSITORY")
		m.Cut("CREATED,CONTAINER_ID,REPOSITORY,COMMAND,PORTS,STATUS,NAMES")
		m.Tables(func(value ice.Maps) {
			if strings.HasPrefix(value["STATUS"], "Up") {
				m.PushButton(s.Open, s.Vimer, s.Xterm, s.Exports, s.Stop)
			} else {
				m.PushButton(s.Restart, s.Drop)
			}
		}).Action(s.Start, s.Prune)
		m.EchoScript(strings.Replace(s.cmd(m, arg[0]), EXEC, RUN, 1)).Cmdy(code.PUBLISH, ice.CONTEXTS, ice.MISC)
		m.StatusTimeCount("SIZE", s.Df(m.Spawn()).Appendv("SIZE")[1])
	} else if len(arg) < 3 || arg[2] == "" {
		m.SplitIndex(s.container(m, EXEC, arg[1], PS)).PushAction(s.Stop).Action(s.Open, s.Vimer, s.Xterm)
		m.EchoScript(s.cmd(m, arg[1])).Cmdy(code.PUBLISH, ice.CONTEXTS, ice.MISC)
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
	s.Code.Iframe(m, "系统页", web.MergePod(m, s.short(m, arg...)), arg...)
}
func (s client) Vimer(m *ice.Message, arg ...string) {
	s.Code.Iframe(m, "编辑器", web.MergePodCmd(m, s.short(m, arg...), web.CODE_VIMER), arg...)
}
func (s client) Xterm(m *ice.Message, arg ...string) {
	ctx.ProcessField(m.Message, web.CODE_XTERM, func() []string {
		name := s.short(m, arg...)
		return []string{s.cmd(m, name), name}
	}, arg...)
}
func (s client) short(m *ice.Message, arg ...string) string {
	return kit.Select(m.Option(CONTAINER_ID), arg, 1)[:12]
}

func init() { ice.CodeCtxCmd(client{}) }
