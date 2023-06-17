package proto

import (
	"path/filepath"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

func (s *GenerateCmds) GenProto() {
	list, _ := filepath.Glob("./idl/*.proto")
	system.Command("", "usr/protoc/bin/protoc", append([]string{"--go_out=./idl", "--go-grpc_out=./idl"}, list...)...)
}
