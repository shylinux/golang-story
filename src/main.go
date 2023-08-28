package main

import (
	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/core/chat/oauth"
	_ "shylinux.com/x/icebergs/misc/java"
	_ "shylinux.com/x/icebergs/misc/node"

	_ "shylinux.com/x/golang-story/src/compile"
	_ "shylinux.com/x/golang-story/src/project"
	_ "shylinux.com/x/golang-story/src/project/gitea"
	_ "shylinux.com/x/golang-story/src/project/mail"
	_ "shylinux.com/x/golang-story/src/project/markdown"
	_ "shylinux.com/x/golang-story/src/project/matrix"
	_ "shylinux.com/x/golang-story/src/runtime"
	_ "shylinux.com/x/golang-story/src/runtime/proto"

	_ "shylinux.com/x/golang-story/src/project/matrix"
)

func main() { print(ice.Run()) }
