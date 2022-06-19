package main

import (
	"shylinux.com/x/ice"

	_ "shylinux.com/x/golang-story/src/compile"
	_ "shylinux.com/x/golang-story/src/project"
	_ "shylinux.com/x/golang-story/src/runtime"

	_ "shylinux.com/x/golang-story/src/grafana"
	_ "shylinux.com/x/golang-story/src/kubernetes"
	_ "shylinux.com/x/golang-story/src/prometheus"
)

func main() { print(ice.Run()) }
