package main

import (
	"shylinux.com/x/ice"

	_ "shylinux.com/x/golang-story/src/compile"
	_ "shylinux.com/x/golang-story/src/project"
	_ "shylinux.com/x/golang-story/src/runtime"
)

func main() { print(ice.Run()) }
