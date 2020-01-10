package main

import (
	"github.com/shylinux/icebergs"
	_ "github.com/shylinux/icebergs/base"
	_ "github.com/shylinux/icebergs/core"

	_ "github.com/shylinux/golang-story/compile"
	_ "github.com/shylinux/golang-story/network"
	_ "github.com/shylinux/golang-story/project"
	_ "github.com/shylinux/golang-story/runtime"
	_ "github.com/shylinux/golang-story/storage"
	_ "github.com/shylinux/golang-story/textual"
)

func main() {
	println(ice.Run())
}
