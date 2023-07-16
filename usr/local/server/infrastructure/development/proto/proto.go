package proto

import (
	"path/filepath"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

func (s *GenerateCmds) GenProto() error {
	list, _ := filepath.Glob("./idl/*.proto")
	cmd := "protoc"
	if system.Exists("usr/protoc/bin/protoc") {
		cmd = "usr/protoc/bin/protoc"
	}
	return system.CommandBuild("", cmd, append([]string{"--go_out=./idl", "--go-grpc_out=./idl"}, list...)...)
}
