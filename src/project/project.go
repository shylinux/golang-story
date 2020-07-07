package project

import (
	"net/http"

	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/base/nfs"
	"github.com/shylinux/icebergs/base/web"
	"github.com/shylinux/icebergs/core/code"
	kit "github.com/shylinux/toolkits"

	"path"
)

const (
	PROJECT = "project"
)
const (
	PROTOBUF = "protobuf"
)

var Index = &ice.Context{Name: PROJECT, Help: "官方库",
	Caches: map[string]*ice.Cache{},
	Configs: map[string]*ice.Config{
		PROJECT: {Name: PROJECT, Help: "官方库", Value: kit.Data(
			"source", "https://dl.google.com/go/go1.14.1.src.tar.gz",
			"target", kit.Dict(
				"linux", "https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz",
				"darwin", "https://dl.google.com/go/go1.14.2.darwin-amd64.pkg",
				"windows", "https://dl.google.com/go/go1.14.2.windows-amd64.msi",
			),
		)},

		PROTOBUF: {Name: PROTOBUF, Help: "protobuf", Value: kit.Data(
			"protoc", "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip",
			"source", "https://github.com/golang/protobuf/",
		)},
	},
	Commands: map[string]*ice.Command{
		ice.CTX_INIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},
		ice.CTX_EXIT: {Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {}},

		PROJECT: {Name: "project", Help: "官方库", Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
			if len(arg) == 0 {
				m.Cmdy(cli.SYSTEM, "go", "list", "std")
			} else {
				m.Cmdy(cli.SYSTEM, "go", "doc", arg)
			}
			m.Set(ice.MSG_APPEND)
		}},

		PROTOBUF: {Name: "protobuf", Help: "protobuf", Action: map[string]*ice.Action{
			"install": {Hand: func(m *ice.Message, arg ...string) {
				if m.Cmd("protoc").Append(cli.CMD_CODE) != "0" {
					// 下载
					msg := m.Cmd(web.SPIDE, "dev", web.CACHE, http.MethodGet, m.Conf(PROTOBUF, "meta.protoc"))
					p := path.Join(m.Conf("web.code.install", "meta.path"), "protoc.zip")
					m.Cmd(web.CACHE, web.WATCH, msg.Append(web.DATA), p)

					// 解压
					m.Option(cli.CMD_DIR, m.Conf("web.code.install", "meta.path"))
					m.Cmd(cli.SYSTEM, "unzip", "protoc.zip")

					// 安装
					m.Cmd(nfs.LINK, kit.Path("bin/protoc"), kit.Path(m.Option(cli.CMD_DIR), "bin/protoc"))
				}

				// 下载
				m.Option(cli.CMD_DIR, kit.Path(m.Conf("web.code.install", "meta.path")))
				m.Cmd(cli.SYSTEM, "git", "clone", m.Conf(PROTOBUF, "meta.source"))

				// 编译
				m.Option(cli.CMD_DIR, kit.Path(m.Conf("web.code.install", "meta.path"), "protobuf/protoc-gen-go"))
				m.Cmd(cli.SYSTEM, "go", "build")

				// 安装
				m.Cmd(nfs.LINK, kit.Path("bin/protoc-gen-go"), kit.Path(m.Option(cli.CMD_DIR), "protoc-gen-go"))
			}},
		}, Hand: func(m *ice.Message, c *ice.Context, cmd string, arg ...string) {
		}},
	},
}

func init() { code.Index.Register(Index, nil) }
