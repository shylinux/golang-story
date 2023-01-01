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
	serve   string `name:"serve arg" help:"服务"`
	restart string `name:"restart" help:"重启"`
	open    string `name:"open" help:"系统"`
	drop    string `name:"drop" help:"删除"`
	df      string `name:"df" help:"磁盘"`
	prune   string `name:"prune" help:"清理"`
	list    string `name:"list IMAGE_ID CONTAINER_ID cmd auto" help:"容器"`
}

func (s client) envhost(m *ice.Message) bool {
	return nfs.ExistsFile(m.Message, strings.TrimPrefix(s.host(m), "unix://")) && s.host(m) != os.Getenv("DOCKER_HOST")
}
func (s client) cmd(m *ice.Message, p string) string {
	if s.envhost(m) {
		return kit.Format("docker --host %s exec -w /root -it %s sh", s.host(m), p)
	}
	return kit.Format("docker exec -w /root -it %s sh", p)
}
func (s client) host(m *ice.Message) string {
	return "unix://" + kit.Path("usr/install/docker/docker.sock")
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
	case tcp.PORT:
		s.List(m.Spawn(), "hi").Tables(func(value ice.Maps) {
			ls := strings.SplitN(value["PORTS"], "->", 2)
			m.Push("port", kit.Slice(strings.Split(ls[0], ":"), -1)[0])
			m.Push("bin", value["COMMAND"])
		})
		port := kit.Int(m.Cmdx(tcp.PORT, tcp.CURRENT)) + 1
		nets := m.Cmdx(cli.SYSTEM, "netstat", "-tln")
		for strings.Contains(nets, kit.Format(":%v ", port)) {
			port += 1
		}
		m.Push(tcp.PORT, port)
		m.Push(ice.BIN, "")
		m.SortIntR(arg[0])
		m.AppendTrans(func(value string, key string, index int) string {
			if index == 0 && key == "port" {
				return value + ":9020"
			}
			return value
		})
	case "dev":
		u := web.OptionUserWeb(m)
		m.Push(arg[0], tcp.PublishLocalhost(m.Message, u.Scheme+"://"+u.Hostname()+ice.DF+u.Port()))
		m.Cmd("spide").Tables(func(value ice.Maps) { m.Push(arg[0], value["client.url"]) })
	}
}
func (s client) Build(m *ice.Message, arg ...string) {
	s.Code.ToastLong(m, "构建中")
	s.Code.Module(m, path.Join(m.Option(nfs.DIR), "Dockerfile"), kit.Format(`
FROM alpine

WORKDIR /root/contexts
COPY ice.linux.amd64 bin/ice.bin

CMD ./bin/ice.bin forever start dev dev
`))
	s.image(m, BUILD, "-t", m.Option(mdb.NAME), m.Option(nfs.DIR))
	s.Code.ToastSuccess(m)
}
func (s client) Pull(m *ice.Message, arg ...string) {
	s.image(m, PULL, m.Option(IMAGE))
}
func (s client) Start(m *ice.Message, arg ...string) {
	if m.Option(ice.CMD) == "" {
		u := web.OptionUserWeb(m)
		args := []string{}
		if m.Option("port") != "" {
			args = append(args, "-p", m.Option("port"))
		}
		m.Option(CONTAINER_ID, s.container(m, kit.Simple(RUN,
			"-e", "ctx_dev="+kit.Select(u.Scheme+"://"+u.Hostname()+ice.DF+u.Port(), m.Option(ice.DEV)),
			args, "-dt", m.Option(IMAGE_ID))...))
	} else {
		m.Option(CONTAINER_ID, s.container(m, RUN, "-dt", m.Option(IMAGE_ID), m.Option(ice.CMD)))
	}
}
func (s client) Serve(m *ice.Message, arg ...string) {
	s.container(m, EXEC, m.Option(CONTAINER_ID), "wget", "-O", "/root/index.sh", web.MergeLink(m.Message, ice.PS))
	s.container(m, EXEC, "-w", "/root", "-dt", m.Option(CONTAINER_ID), nfs.SH, "/root/index.sh", "app", ice.DEV, web.MergeLink(m.Message, ice.PS))
}
func (s client) Restart(m *ice.Message, arg ...string) {
	s.container(m, RESTART, m.Option(CONTAINER_ID))
	s.Serve(m)
}
func (s client) Stop(m *ice.Message, arg ...string) {
	web.PushNoticeToast(m.Message, ice.PROCESS)
	defer web.PushNoticeToast(m.Message, ice.SUCCESS)
	if m.Option("PID") != "" { // 结束进程
		s.container(m, EXEC, m.Option(CONTAINER_ID), KILL, m.Option("PID"))
	} else { // 结束容器
		s.container(m, STOP, m.Option(CONTAINER_ID))
	}
}
func (s client) Open(m *ice.Message, arg ...string) {
	s.Code.Iframe(m, "系统页", web.MergePod(m, kit.Select(m.Option(CONTAINER_ID), arg, 1), "river", "", "storm", ""), arg...)
}
func (s client) Drop(m *ice.Message, arg ...string) {
	if m.Option(CONTAINER_ID) != "" { // 删除容器
		s.container(m, RM, m.Option(CONTAINER_ID))
	} else if m.Option(IMAGE_ID) != "" { // 删除镜像
		s.image(m, RM, m.Option(IMAGE_ID))
	}
}
func (s client) Df(m *ice.Message, arg ...string) *ice.Message {
	m.SplitIndex(s.docker(m, cli.SYSTEM, "df"))
	return m
}
func (s client) Prune(m *ice.Message, arg ...string) {
	if len(arg) > 0 { // 清理容器
		m.Echo(s.container(m, PRUNE, "-f"))
	} else { // 清理镜像
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
	if len(arg) < 1 || arg[0] == "" { // 镜像列表
		m.SplitIndex(strings.Replace(s.image(m, LS), "IMAGE ID", IMAGE_ID, 1))
		m.Cut("CREATED,IMAGE_ID,SIZE,REPOSITORY,TAG")
		m.PushAction(s.Drop).Action(s.Build, s.Pull, s.Df, s.Prune)
		m.StatusTimeCount("SIZE", s.Df(m.Spawn()).Append("SIZE"))

	} else if len(arg) < 2 || arg[1] == "" { // 容器列表
		m.SplitIndex(strings.Replace(s.container(m, LS, "-a"), "CONTAINER ID", CONTAINER_ID, 1)).RenameAppend("IMAGE", "REPOSITORY")
		m.Cut("CREATED,CONTAINER_ID,REPOSITORY,COMMAND,PORTS,STATUS,NAMES")
		m.Tables(func(value ice.Maps) {
			if strings.HasPrefix(value["STATUS"], "Up") {
				m.PushButton(s.Open, s.Vimer, s.Xterm, s.Stop)
			} else {
				m.PushButton(s.Restart, s.Drop)
			}
		}).Action(s.Start, s.Prune)
		m.EchoScript(strings.Replace(s.cmd(m, arg[0]), "exec", "run", 1))
		m.EchoScript(m.Cmdx(code.PUBLISH, ice.CONTEXTS, code.INSTALL))
		m.StatusTimeCount("SIZE", s.Df(m.Spawn()).Appendv("SIZE")[1])

	} else if len(arg) < 3 || arg[2] == "" { // 进程列表
		m.SplitIndex(s.container(m, EXEC, arg[1], PS)).PushAction(s.Stop).Action(s.Open, s.Xterm, s.Serve)
		m.EchoScript(s.cmd(m, arg[1]))
		m.StatusTimeCount()

	} else { // 执行命令
		m.Echo(s.container(m, kit.Simple(EXEC, arg[1], kit.Split(arg[2], ice.SP, ice.SP))...))
		m.StatusTimeCount()
	}
	return m
}
func (s client) Xterm(m *ice.Message, arg ...string) {
	s.Code.Xterm(m, []string{mdb.TYPE, s.cmd(m, kit.Select(m.Option(CONTAINER_ID), arg, 1)), mdb.NAME, m.Option(CONTAINER_ID)}, arg...)
}
func (s client) Vimer(m *ice.Message, arg ...string) {
	s.Code.Iframe(m, "编辑器", web.MergePodCmd(m, m.Option(CONTAINER_ID), "web.code.vimer"), arg...)
}

func init() { ice.CodeCtxCmd(client{}) }
