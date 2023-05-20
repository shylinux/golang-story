package project

import (
	"path"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/nfs"
)

type protobuf struct {
	ice.Code
	linux  string `data:"https://github.com/protocolbuffers/protobuf/releases/download/v23.1/protoc-23.1-linux-x86_64.zip"`
	darwin string `data:"https://github.com/protocolbuffers/protobuf/releases/download/v23.1/protoc-23.1-osx-x86_64.zip"`
	list   string `name:"list auto install" help:"协议"`
}

func (s protobuf) Install(m *ice.Message, arg ...string) {
	s.Code.Install(m)
	m.Cmdy(cli.SYSTEM, nfs.PUSH, path.Join("usr/install/protoc/bin"))
	m.Cmdy(cli.SYSTEM, "go", "install", "google.golang.org/protobuf/cmd/protoc-gen-go")
	m.Cmdy(cli.SYSTEM, "go", "install", "google.golang.org/grpc/cmd/protoc-gen-go-grpc")
}
func (s protobuf) List(m *ice.Message, arg ...string) {
	m.Cmdy(cli.SYSTEM, "go", "doc", "google.golang.org/grpc")
}

func init() { ice.CodeModCmd(protobuf{}) }
