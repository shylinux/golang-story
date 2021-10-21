package project

import (
	"net/http"
	"path"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/icebergs/core/code"
	kit "shylinux.com/x/toolkits"
)

type protobuf struct {
	protoc string `data:"https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip"`
	source string `data"https://github.com/golang/protobuf/"`

	list string `name:"list auto" help:"官方库"`
}

func (p protobuf) Install(m *ice.Message, arg ...string) {
	if !cli.IsSuccess(m.Cmd("protoc")) {
		// 下载
		msg := m.Cmd(web.SPIDE, "dev", web.CACHE, http.MethodGet, m.Conf(m.PrefixKey(), "meta.protoc"))
		p := path.Join(m.Conf(code.INSTALL, kit.META_PATH), "protoc.zip")
		m.Cmd(web.CACHE, web.WATCH, msg.Append(web.DATA), p)

		// 解压
		m.Option(cli.CMD_DIR, m.Conf(code.INSTALL, kit.META_PATH))
		m.Cmd(cli.SYSTEM, "unzip", "protoc.zip")

		// 安装
		m.Cmd(nfs.LINK, kit.Path("bin/protoc"), kit.Path(m.Option(cli.CMD_DIR), "bin/protoc"))
	}

	// 下载
	m.Option(cli.CMD_DIR, kit.Path(m.Conf(code.INSTALL, kit.META_PATH)))
	m.Cmd(cli.SYSTEM, "git", "clone", m.Conf(m.PrefixKey(), kit.META_SOURCE))

	// 编译
	m.Option(cli.CMD_DIR, kit.Path(m.Conf(code.INSTALL, kit.META_PATH), "protobuf/protoc-gen-go"))
	m.Cmd(cli.SYSTEM, "go", "build")

	// 安装
	m.Cmd(nfs.LINK, kit.Path("bin/protoc-gen-go"), kit.Path(m.Option(cli.CMD_DIR), "protoc-gen-go"))
}
func (p protobuf) List(m *ice.Message, arg ...string) {
	m.Cmdy(cli.SYSTEM, "go", "doc", "runtime")
}

func init() { ice.Cmd("web.code.golang.protobuf", protobuf{}) }
