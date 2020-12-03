package main

import (
	ice "github.com/shylinux/icebergs"
	_ "github.com/shylinux/icebergs/base"
	_ "github.com/shylinux/icebergs/core"
	_ "github.com/shylinux/icebergs/misc"

	_ "github.com/shylinux/golang-story/src/compile"
	_ "github.com/shylinux/golang-story/src/project"
	_ "github.com/shylinux/golang-story/src/runtime"

	_ "github.com/shylinux/golang-story/src/gogs"
	_ "github.com/shylinux/golang-story/src/grafana"
	_ "github.com/shylinux/golang-story/src/leveldb"
	_ "github.com/shylinux/golang-story/src/prometheus"
	_ "github.com/shylinux/golang-story/src/rocksdb"
	_ "github.com/shylinux/golang-story/src/tcmalloc"
)

func main() { print(ice.Run()) }
