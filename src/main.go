package main

import (
	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/core/chat/oauth"
	_ "shylinux.com/x/icebergs/misc/java"
	_ "shylinux.com/x/icebergs/misc/node"

	_ "shylinux.com/x/golang-story/src/compile"
	_ "shylinux.com/x/golang-story/src/leetcode"
	_ "shylinux.com/x/golang-story/src/project"
	_ "shylinux.com/x/golang-story/src/project/gitea"
	_ "shylinux.com/x/golang-story/src/project/image"
	_ "shylinux.com/x/golang-story/src/project/mail"
	_ "shylinux.com/x/golang-story/src/project/md"
	_ "shylinux.com/x/golang-story/src/project/proto"
	_ "shylinux.com/x/golang-story/src/runtime"
)

func main() { print(ice.Run()) }
