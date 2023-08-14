package main

import (
	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/misc/java"
	_ "shylinux.com/x/icebergs/misc/node"

	_ "shylinux.com/x/golang-story/src/compile"
	_ "shylinux.com/x/golang-story/src/markdown"
	_ "shylinux.com/x/golang-story/src/project"
	_ "shylinux.com/x/golang-story/src/proto"
	_ "shylinux.com/x/golang-story/src/runtime"
	_ "shylinux.com/x/golang-story/src/sqlite"
)

func main() { print(ice.Run()) }
