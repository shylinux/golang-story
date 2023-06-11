package main

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/proto"
)

func main() {
	container.New(infrastructure.Init).Invoke(func(gen *proto.Generate) {
		gen.GenValid()
		// gen.GenTest()
		gen.GenGoAPI()
		gen.GenJsAPI()
	})
}
