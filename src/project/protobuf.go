package project

import (
	"net/http"
	"path"

	ice "github.com/shylinux/icebergs"
	"github.com/shylinux/icebergs/base/cli"
	"github.com/shylinux/icebergs/base/nfs"
	"github.com/shylinux/icebergs/base/web"
	kit "github.com/shylinux/toolkits"
)

const PROTOBUF = "protobuf"

func init() {
	Index.Merge(&ice.Context{Name: PROJECT, Help: "官方库",
		Configs: map[string]*ice.Config{
			PROTOBUF: {Name: PROTOBUF, Help: "protobuf", Value: kit.Data(
				"protoc", "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip",
				"source", "https://github.com/golang/protobuf/",
			)},
		},
		Commands: map[string]*ice.Command{
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
	})
}
