package project

import (
	"context"
	"fmt"
	"path"
	"text/template"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

type ProjectCmds struct{ config *config.Config }

func (s *ProjectCmds) Create(ctx context.Context, arg ...string) {
	if !system.Exists("go.mod") {
		fmt.Println("please run go mod init")
		return
	}
	dir := ""
	if len(arg) > 0 {
		dir = arg[0]
	}
	mod := path.Join(logs.PwdModPath(), dir)
	for _, file := range templateList {
		if system.Exists(path.Join(dir, file.Path)) {
			continue
		}
		system.NewTemplateFile(path.Join(dir, file.Path), file.Text, template.FuncMap{
			"PwdModPath": func() string { return mod },
		}, map[string]interface{}{})
	}

}
func (s *ProjectCmds) List(ctx context.Context, arg ...string) {
}
func NewProjectCmds(config *config.Config, cmds *cmds.Cmds) *ProjectCmds {
	s := &ProjectCmds{config: config}
	cmds = cmds.Add("project", "project command", s.List)
	cmds.Add("create", "create path", s.Create)
	return s
}

var templateList = []struct {
	Path string
	Text string
}{
	{Path: "cmd/cmds.go", Text: `
package main

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/server"
	"{{ PwdModPath }}/idl/cli"
)

func main() {
	c := container.New(cli.Init, development.Init, infrastructure.Init)
	c.Invoke(func(s *cmds.Cmds, _ *server.ServerCmds, _ *cli.MainServiceCmds) error { return s.Run() })
}
`},
	{Path: "cmd/main.go", Text: `
package main

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
	"{{ PwdModPath }}/idl"
)

func main() {
	c := container.New(idl.Init, infrastructure.Init)
	c.Invoke(func(s *server.MainServer, _ *idl.MainController) error { return s.Run() })
}
`},
	{Path: "config/service.yaml", Text: `
logs:
  pid: log/service.pid
  path: log/service.log
  maxsize: 10 # 10M
  maxage: 30  # 30days
proxy:
  export: true
  simple: false
  root: "usr/vue-element-admin/dist/"
  port: 8081
token:
  issuer: "auth"
  secret: "auth"
  expire: "72h"
consul:
  enable: false
  addr: ":8500"
  interval: 10s
server:
  port: 9090
engine:
  queue:
    enable: true
    type: pulsar
    host: 127.0.0.1
    port: 6650
  cache:
    enable: true
    type: redis
    host: 127.0.0.1
    port: 6379
  search:
    enable: true
    type: elasticsearch
    index: demo
    host: 127.0.0.1
    port: 9200
  storage:
    type: sqlite
    username: demo
    password: demo
    database: demo
    host: 127.0.0.1
    port: 3306
`},
	{Path: "idl/idl.go", Text: `
package idl

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func Init(c *container.Container) {
	c.Provide(NewMainController)
}

type MainController struct{}

func NewMainController() *MainController {
	return &MainController{}
}
`},
	{Path: "idl/cli/cli.go", Text: `
package cli

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func Init(c *container.Container) {
	c.Provide(NewMainServiceCmds)
}

type MainServiceCmds struct{}

func NewMainServiceCmds() *MainServiceCmds {
	return &MainServiceCmds{}
}
`},
	{Path: "Makefile", Text: `
export CGO_ENABLED=1
def: matrix idl server restart
all: matrix idl server deploy restart test pack
matrix:
	go build -v -o bin/matrix cmd/cmds.go
idl:
	./bin/matrix unpack protoc
	./bin/matrix generate
server:
	go build -v -o bin/server cmd/main.go
deploy:
	./bin/matrix deploy node java git golang
	./bin/matrix deploy nginx redis consul pulsar es
restart:
	./bin/matrix server restart
test:
	go test -v --count=1 ./idl/ts
	go test -v --count=1 ./idl/test
pack:
	tar zcvf service.tar.gz bin config usr/vue-element-admin/dist
clean:
	rm -rf bin/* log/* idl/test/log/* usr/vue-element-admin/dist/*
.PHONY: matrix idl server restart test pack clean
`},
}
