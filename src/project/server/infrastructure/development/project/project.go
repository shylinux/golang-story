package project

import (
	"context"
	"fmt"
	"path"
	"text/template"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

const PROJECT = "project"

type ProjectCmds struct {
	name string
}

func (s *ProjectCmds) Create(ctx context.Context, arg ...string) {
	if len(arg) == 0 {
		fmt.Println(fmt.Errorf("need params path"))
		return
	} else if system.Exists(arg[0]) {
		fmt.Println(fmt.Errorf("project already exists"))
		return
	}
	for _, file := range templateList {
		if system.Exists(path.Join(arg[0], file.Path)) {
			continue
		}
		system.NewTemplateFile(path.Join(arg[0], file.Path), file.Text, template.FuncMap{
			"PwdModPath": func() string { return path.Join(logs.PwdModPath(), arg[0]) },
		}, nil)
	}

}
func (s *ProjectCmds) List(ctx context.Context, arg ...string) {
}
func NewProjectCmds(cmds *cmds.Cmds) *ProjectCmds {
	s := &ProjectCmds{name: PROJECT}
	cmds = cmds.Add(s.name, "project command", s.List)
	cmds = cmds.Add("create", "create path", s.Create)
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
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository/mysql"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
	"{{ PwdModPath }}/idl"
)

func main() {
	c := container.New(idl.Init, infrastructure.Init)
	c.Provide(mysql.New)
	c.Invoke(func(s *server.MainServer, _ *idl.MainController) error { return s.Run() })
}
`},
	{Path: "config/service.yaml", Text: `
logs:
  pid: log/service.pid
  path: log/service.log
  maxsize: 10 # 10M
  maxage: 30  # 30days
  stdout: false
consul:
  enable: true
  addr: ":8500"
  interval: 10s
  workid: 2
server:
  name: demo
  port: 9090
engine:
  storage:
    name: mysql
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
