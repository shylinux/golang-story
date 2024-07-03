package main

import (
	"shylinux.com/x/ice"

	_ "shylinux.com/x/golang-story/src/compile"
	_ "shylinux.com/x/golang-story/src/project/media"
	_ "shylinux.com/x/golang-story/src/runtime"

	_ "shylinux.com/x/golang-story/src/project/bluetooth"

	_ "shylinux.com/x/golang-story/src/project/email"
)

func main() { print(ice.Run()) }

func init() {
	ice.Info.NodeIcon = "usr/icons/go.png"
	ice.Info.NodeMain = "web.code.golang.runtime.studio"
}
