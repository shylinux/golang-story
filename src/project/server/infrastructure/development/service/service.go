package service

import (
	"context"
	"fmt"
	"strings"
	"text/template"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/proto"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

const SERVICE = "service"

type ServiceCmds struct {
	config *config.Config
	name   string
}

func (s *ServiceCmds) Create(ctx context.Context, arg ...string) {
	if len(arg) == 0 {
		fmt.Println(fmt.Errorf("need params name"))
		return
	}
	for _, file := range templateList {
		file.Path = strings.ReplaceAll(file.Path, "name", arg[0])
		if system.Exists(file.Path) {
			// continue
		}
		system.NewTemplateFile(file.Path, file.Text, template.FuncMap{
			"PwdModPath": func() string { return logs.PwdModPath() },
		}, map[string]string{
			"package": s.config.Server.Name + "." + arg[0],
			"service": proto.Capital(arg[0]) + "Service",
			"name":    proto.Capital(arg[0]),
			"table":   arg[0],
		})
		system.Command("", "gofmt", "-w", file.Path)
	}

}
func (s *ServiceCmds) List(ctx context.Context, arg ...string) {}
func NewServiceCmds(cmds *cmds.Cmds, config *config.Config) *ServiceCmds {
	s := &ServiceCmds{name: SERVICE, config: config}
	cmds = cmds.Add(s.name, "service command", s.List)
	cmds.Add("create", "create path", s.Create)
	return s
}

var templateList = []struct {
	Path string
	Text string
}{
	{Path: "idl/name.proto", Text: `
syntax = "proto3";
option go_package ="./pb";
package {{ .package }};

service {{ .service }} {
    rpc Create ({{ .name }}CreateRequest) returns ({{ .name }}CreateReply);
    rpc Remove ({{ .name }}RemoveRequest) returns ({{ .name }}RemoveReply);
    rpc Info ({{ .name }}InfoRequest) returns ({{ .name }}InfoReply);
    rpc List ({{ .name }}ListRequest) returns ({{ .name }}ListReply);
}

message {{ .name }}CreateRequest {
    // length > 6
    string name = 1;
}
message {{ .name }}CreateReply {
	{{ .name }}Error error = 1;
	{{ .name }} data = 2;
}

message {{ .name }}RemoveRequest {
    // required
    int64 {{ .name }}ID = 1;
}
message {{ .name }}RemoveReply {
	{{ .name }}Error error = 1;
}

message {{ .name }}InfoRequest {
    // required
    int64 {{ .name }}ID = 1;
}
message {{ .name }}InfoReply {
	{{ .name }}Error error = 1;
	{{ .name }} data = 2;
}

message {{ .name }}ListRequest {
    // default 1
    int64 page = 1;
    // default 10
    int64 count = 2;
    string key = 3;
    string value = 4;
}
message {{ .name }}ListReply {
	{{ .name }}Error error = 1;
    repeated {{ .name }} data = 2;
    int64 total = 3;
}

message {{ .name }} {
    int64 {{ .name }}ID = 1;
    string name = 2;
}

message {{ .name }}Error {
    int64 code = 1;
    string info = 2;
}
`},
	{Path: "controller/name.go", Text: `
package controller

import (
	"context"

	dt "shylinux.com/x/golang-story/src/project/server/domain/trans"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/server"
	"{{ PwdModPath }}/domain/trans"
	"{{ PwdModPath }}/idl/pb"
	"{{ PwdModPath }}/service"
)

type {{ .name }}Controller struct {
	pb.Unimplemented{{ .name }}ServiceServer
	Main    *server.MainServer
	service *service.{{ .name }}Service
	name    string
}

func New{{ .name }}Controller(config *config.Config, server *server.MainServer, service *service.{{ .name }}Service) *{{ .name }}Controller {
	controller := &{{ .name }}Controller{Main: server, service: service, name: pb.{{ .name }}Service_ServiceDesc.ServiceName}
	if !config.Internal["{{ .table }}"].Export {
		return controller
	}
	server.Proxy.Register(controller.name, controller)
	server.Server.Register(&pb.{{ .name }}Service_ServiceDesc, controller)
	consul.Tags = append(consul.Tags, controller.name)
	return controller
}
func (s *{{ .name }}Controller) Create(ctx context.Context, req *pb.{{ .name }}CreateRequest) (*pb.{{ .name }}CreateReply, error) {
	space, err := s.service.Create(ctx, req.Name)
	return &pb.{{ .name }}CreateReply{Data: trans.{{ .name }}DTO(space)}, errors.NewCreateFailResp(err)
}
func (s *{{ .name }}Controller) Remove(ctx context.Context, req *pb.{{ .name }}RemoveRequest) (*pb.{{ .name }}RemoveReply, error) {
	return &pb.{{ .name }}RemoveReply{}, errors.NewRemoveFailResp(s.service.Remove(ctx, req.{{ .name }}ID))
}
func (s *{{ .name }}Controller) Info(ctx context.Context, req *pb.{{ .name }}InfoRequest) (*pb.{{ .name }}InfoReply, error) {
	space, err := s.service.Info(ctx, req.{{ .name }}ID)
	return &pb.{{ .name }}InfoReply{Data: trans.{{ .name }}DTO(space)}, errors.NewInfoFailResp(err)
}
func (s *{{ .name }}Controller) List(ctx context.Context, req *pb.{{ .name }}ListRequest) (*pb.{{ .name }}ListReply, error) {
	list, total, err := s.service.List(ctx, req.Page, req.Count, req.Key, req.Value)
	data := []*pb.{{ .name }}{}
	dt.ListDTO(list, trans.{{ .name }}DTO, &data)
	return &pb.{{ .name }}ListReply{Data: data, Total: total}, errors.NewListFailResp(err)
}
`},
	{Path: "service/name.go", Text: `
package service

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/uuid"
	"{{ PwdModPath }}/domain/model"
	"shylinux.com/x/golang-story/src/project/server/service"
)

type {{ .name }}Service struct {
	*uuid.Generate
	storage repository.Storage
}

func New{{ .name }}Service(generate *uuid.Generate, storage repository.Storage) *{{ .name }}Service {
	storage.AutoMigrate(&model.{{ .name }}{})
	return &{{ .name }}Service{generate, storage}
}
func (s *{{ .name }}Service) Create(ctx context.Context, name string) (*model.{{ .name }}, error) {
	space := &model.{{ .name }}{
		{{ .name }}ID: s.Generate.GenID(), Name: name,
	}
	return space, errors.NewCreateFail(s.storage.Insert(ctx, space))
}
func (s *{{ .name }}Service) Remove(ctx context.Context, spaceID int64) error {
	return errors.NewRemoveFail(s.storage.Delete(ctx, &model.{{ .name }}{
		{{ .name }}ID: spaceID,
	}))
}
func (s *{{ .name }}Service) Info(ctx context.Context, spaceID int64) (*model.{{ .name }}, error) {
	space := &model.{{ .name }}{
		{{ .name }}ID: spaceID,
	}
	return space, errors.NewInfoFail(s.storage.SelectOne(ctx, space))
}
func (s *{{ .name }}Service) List(ctx context.Context, page, count int64, key, value string) (list []*model.{{ .name }}, total int64, err error) {
	condition, arg := service.Clause(key != "" && value != "", key+" = ? and ", key, value)
	total, err = s.storage.SelectList(ctx, &model.{{ .name }}{}, &list, page, count, condition, arg...)
	return list, total, errors.NewListFail(err)
}
`},
	{Path: "domain/model/name.go", Text: `
package model

import (
	"fmt"

	"shylinux.com/x/golang-story/src/project/server/domain/model"
)

type {{ .name }} struct {
	model.Common
	{{ .name }}ID   int64
	Name string
}

func (s {{ .name }}) TableName() string { return "{{ .table }}" }
func (s {{ .name }}) GetKey() string    { return "{{ .table }}_id" }
func (s {{ .name }}) GetID() string     { return fmt.Sprintf("%d", s.{{ .name }}ID) }
`},
	{Path: "domain/trans/name.go", Text: `
package trans

import (
	"{{ PwdModPath }}/domain/model"
	"{{ PwdModPath }}/idl/pb"
)

func {{ .name }}DTO({{ .table }} *model.{{ .name }}) *pb.{{ .name }} {
	if {{ .table }} == nil {
		return nil
	}
	return &pb.{{ .name }}{
		{{ .name }}ID: {{ .table }}.{{ .name }}ID,
	}
}
`},
}
