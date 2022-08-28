package docker

import (
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

const (
	DOCKER    = "docker"
	IMAGE     = "image"
	CONTAINER = "container"

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

type runtime struct {
	ice.Code
	pull    string `name:"pull image=alpine" help:"下载"`
	start   string `name:"start cmd=/bin/sh" help:"启动"`
	serve   string `name:"serve arg" help:"服务"`
	restart string `name:"restart" help:"重启"`
	drop    string `name:"drop" help:"删除"`
	prune   string `name:"prune" help:"清理"`
	list    string `name:"list IMAGE_ID CONTAINER_ID cmd auto" help:"容器"`
}

func (s runtime) host(m *ice.Message) string {
	return "unix://" + kit.Path("usr/install/docker/docker.sock")
}
func (s runtime) docker(m *ice.Message, arg ...string) string {
	if nfs.ExistsFile(m.Message, s.host(m)) {
		m.Option(cli.CMD_ENV, "DOCKER_HOST", s.host(m))
	}
	return m.Cmdx(cli.SYSTEM, DOCKER, arg)
}
func (s runtime) image(m *ice.Message, arg ...string) string {
	return s.docker(m, kit.Simple(IMAGE, arg)...)
}
func (s runtime) container(m *ice.Message, arg ...string) string {
	return s.docker(m, kit.Simple(CONTAINER, arg)...)
}

func (s runtime) Inputs(m *ice.Message, arg ...string) {
	switch arg[0] {
	case IMAGE:
		m.Push(arg[0], "busybox", "alpine", "centos", "ubuntu")
	}
}
func (s runtime) Pull(m *ice.Message, arg ...string) {
	s.image(m, PULL, m.Option(IMAGE))
}
func (s runtime) Start(m *ice.Message, arg ...string) {
	m.Option(CONTAINER_ID, s.container(m, RUN, "-dt", m.Option(IMAGE_ID), m.Option(ice.CMD)))
	s.Serve(m)
}
func (s runtime) Serve(m *ice.Message, arg ...string) {
	s.container(m, EXEC, m.Option(CONTAINER_ID), "wget", "-O", "/root/index.sh", web.MergeLink(m.Message, ice.PS))
	s.container(m, EXEC, "-w", "/root", "-dt", m.Option(CONTAINER_ID), "sh", "/root/index.sh", "app", "dev", web.MergeLink(m.Message, ice.PS))
}
func (s runtime) Restart(m *ice.Message, arg ...string) {
	s.container(m, RESTART, m.Option(CONTAINER_ID))
	s.Serve(m)
}
func (s runtime) Stop(m *ice.Message, arg ...string) {
	web.PushNoticeToast(m.Message, "process")
	defer web.PushNoticeToast(m.Message, "success")
	if m.Option("PID") != "" { // 结束进程
		s.container(m, EXEC, m.Option(CONTAINER_ID), KILL, m.Option("PID"))
	} else { // 结束容器
		s.container(m, STOP, m.Option(CONTAINER_ID))
	}
}
func (s runtime) Open(m *ice.Message, arg ...string) {
	m.ProcessOpen(web.MergePod(m, m.Option(CONTAINER_ID)))
}
func (s runtime) Drop(m *ice.Message, arg ...string) {
	if m.Option(CONTAINER_ID) != "" { // 删除容器
		s.container(m, RM, m.Option(CONTAINER_ID))
	} else if m.Option(IMAGE_ID) != "" { // 删除镜像
		s.image(m, RM, m.Option(IMAGE_ID))
	}
}
func (s runtime) Prune(m *ice.Message, arg ...string) {
	if len(arg) > 0 { // 清理容器
		m.Echo(s.container(m, PRUNE, "-f"))
	} else { // 清理镜像
		m.Echo(s.image(m, PRUNE, "-f"))
	}
}
func (s runtime) Modify(m *ice.Message, arg ...string) {
	switch arg[0] {
	case "NAMES":
		s.container(m, RENAME, m.Option(CONTAINER_ID), arg[1])
	}
}
func (s runtime) List(m *ice.Message, arg ...string) {
	if len(arg) < 1 || arg[0] == "" { // 镜像列表
		m.SplitIndex(strings.Replace(s.image(m, LS), "IMAGE ID", IMAGE_ID, 1))
		m.Cut("CREATED,IMAGE_ID,SIZE,REPOSITORY,TAG")
		m.PushAction(s.Drop).Action(s.Pull, s.Prune)

	} else if len(arg) < 2 || arg[1] == "" { // 容器列表
		m.SplitIndex(strings.Replace(s.container(m, LS, "-a"), "CONTAINER ID", CONTAINER_ID, 1)).RenameAppend("IMAGE", "REPOSITORY")
		m.Cut("CREATED,CONTAINER_ID,REPOSITORY,COMMAND,PORTS,STATUS,NAMES")
		m.Tables(func(value ice.Maps) {
			if strings.HasPrefix(value["STATUS"], "Up") {
				m.PushButton(s.Open, s.Stop)
			} else {
				m.PushButton(s.Restart, s.Drop)
			}
		}).Action(s.Start, s.Prune)

	} else if len(arg) < 3 || arg[2] == "" { // 进程列表
		m.SplitIndex(s.container(m, EXEC, arg[1], PS)).PushAction(s.Stop).Action(s.Xterm, s.Serve)
		if nfs.ExistsFile(m.Message, s.host(m)) {
			m.EchoScript(kit.Format("docker --host %s exec -w /root -it %s sh", s.host(m), arg[1]))
		} else {
			m.EchoScript(kit.Format("docker exec -w /root -it %s sh", arg[1]))
		}

	} else { // 执行命令
		m.Echo(s.container(m, kit.Simple(EXEC, arg[1], kit.Split(arg[2], " ", " "))...))
	}
	m.StatusTimeCount()
}
func (s runtime) Xterm(m *ice.Message, arg ...string) {
	if nfs.ExistsFile(m.Message, s.host(m)) {
		s.Code.Xterm(m, kit.Format("docker --host %s exec -w /root -it %s sh", s.host(m), m.Option(CONTAINER_ID)), arg...)
	} else {
		s.Code.Xterm(m, kit.Format("docker exec -w /root -it %s sh", m.Option(CONTAINER_ID)), arg...)
	}
}

func init() { ice.CodeCtxCmd(runtime{}) }
