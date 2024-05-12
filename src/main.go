package main

import (
	"shylinux.com/x/ice"

	_ "shylinux.com/x/golang-story/src/runtime"
)

func main() { print(ice.Run()) }

func init() {
	ice.Info.NodeIcon = "usr/icons/go.png"
	ice.Info.NodeMain = "web.code.golang.runtime.studio"
}
