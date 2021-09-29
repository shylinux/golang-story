package main

import (
	"shylinux.com/x/ice"

	_ "shylinux.com/x/golang-story/src/compile"
	_ "shylinux.com/x/golang-story/src/project"
	_ "shylinux.com/x/golang-story/src/runtime"

	_ "shylinux.com/x/golang-story/src/gogs"
	_ "shylinux.com/x/golang-story/src/grafana"
	_ "shylinux.com/x/golang-story/src/leveldb"
	_ "shylinux.com/x/golang-story/src/prometheus"
	_ "shylinux.com/x/golang-story/src/rocksdb"
	_ "shylinux.com/x/golang-story/src/tcmalloc"
)

func main() { print(ice.Run()) }
