package project

import (
	"github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/core/code"
	"github.com/shylinux/toolkits"

	"path"
)

var Index = &ice.Context{Name: "project", Help: "官方库",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		"project": {Name: "project", Help: "官方库", Value: kit.Data(
			"source", "https://dl.google.com/go/go1.14.1.src.tar.gz",
			"target", kit.Dict(
				"linux", "https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz",
				"darwin", "https://dl.google.com/go/go1.14.2.darwin-amd64.pkg",
				"windows", "https://dl.google.com/go/go1.14.2.windows-amd64.msi",
			),
		)},

		"protobuf": {Name: "protobuf", Help: "protobuf", Value: kit.Data(
			"protoc", "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip",
			"source", "https://github.com/golang/protobuf/",
		)},
	},
	Commands: map[string]*ice.Command{
		ice.ICE_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.ICE_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		"project": {Name: "project", Help: "官方库", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			m.Cmdy(ice.CLI_SYSTEM, "go", "list", "std")
			m.Set(ice.MSG_APPEND)
		}},

		"protobuf": {Name: "protobuf install source protoc", Help: "protobuf", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			switch arg[0] {
			case "install":
				// 下载
				msg := m.Cmd(ice.WEB_SPIDE, "dev", "cache", "GET", m.Conf("protobuf", "meta.protoc"))
				p := path.Join(m.Conf("web.code.install", "meta.path"), "protoc.zip")
				m.Cmd(ice.WEB_CACHE, "watch", msg.Append("data"), p)

				// 解压
				m.Option("cmd_dir", m.Conf("web.code.install", "meta.path"))
				m.Cmd(ice.CLI_SYSTEM, "unzip", "preotoc.zip")

				// 安装
				m.Cmd("nfs.link", kit.Path("bin/protoc"), kit.Path(m.Option("cmd_dir"), "bin/protoc"))

				// 下载
				msg = m.Cmd(ice.WEB_SPIDE, "dev", "cache", "GET", m.Conf("protobuf", "meta.protoc"))
				m.Cmd(ice.CLI_SYSTEM, "git", "clone", m.Conf("web.code.install", "meta.path"))

				// 编译
				m.Option("cmd_dir", kit.Path(m.Conf("web.code.install", "meta.path"), "protobuf/protoc-gen-go"))
				m.Cmd(ice.CLI_SYSTEM, "go", "build")

				// 安装
				m.Cmd("nfs.link", kit.Path("bin/protoc-gen-go"), kit.Path(m.Option("cmd_dir"), "protoc-gen-go"))
			}
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
