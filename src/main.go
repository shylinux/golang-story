package main

import (
	"shylinux.com/x/ice"
	_ "shylinux.com/x/ice/devops"

	_ "shylinux.com/x/golang-story/src/compile"
	_ "shylinux.com/x/golang-story/src/project"
	_ "shylinux.com/x/golang-story/src/project/gitea"
	_ "shylinux.com/x/golang-story/src/project/image"
	_ "shylinux.com/x/golang-story/src/project/leetcode"
	_ "shylinux.com/x/golang-story/src/project/mail"
	_ "shylinux.com/x/golang-story/src/project/md"
	_ "shylinux.com/x/golang-story/src/project/office"
	_ "shylinux.com/x/golang-story/src/project/proto"
	_ "shylinux.com/x/golang-story/src/project/stats"
	_ "shylinux.com/x/golang-story/src/runtime"
)

func main() { print(ice.Run()) }

func init() { ice.Info.NodeIcon = "usr/icons/go.png" }
