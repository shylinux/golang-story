package main

import (
	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/misc/node"

	_ "shylinux.com/x/golang-story/src/compile"
	_ "shylinux.com/x/golang-story/src/docker"
	_ "shylinux.com/x/golang-story/src/project"
	_ "shylinux.com/x/golang-story/src/runtime"
	_ "shylinux.com/x/mysql-story/src/client"
	_ "shylinux.com/x/redis-story/src/client"
)

func main() { print(ice.Run()) }
