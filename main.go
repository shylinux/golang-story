package main

import (
	"github.com/shylinux/icebergs"
	_ "github.com/shylinux/icebergs/base"
	_ "github.com/shylinux/icebergs/core"
	_ "github.com/shylinux/icebergs/misc"

	_ "github.com/shylinux/golang-story/compile"
	_ "github.com/shylinux/golang-story/runtime"
)

func main() {
	println(ice.Run())
}
