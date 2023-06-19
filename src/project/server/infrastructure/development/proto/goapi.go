package proto

import (
	"html/template"
	"path"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func (s *GenerateCmds) GenGoAPI() {
	serviceList := []string{}
	for name, proto := range s.protos {
		serviceList = append(serviceList, proto[PACKAGE].List...)
		s.Render(path.Join(s.conf.GoPath, name+".go"), _goapi_client, proto[PACKAGE].List, template.FuncMap{
			"PwdModPath": func() string { return logs.PwdModPath() },
		})
	}
	s.Render(path.Join(s.conf.GoPath, path.Base(s.conf.GoPath)+".go"), _goapi_init, serviceList, nil)
	for i, v := range serviceList {
		serviceList[i] = strings.TrimSuffix(v, "Service")
	}
	s.Render(path.Join(s.conf.Path, "idl.go"), _idl_template, serviceList, template.FuncMap{
		"PwdModPath": func() string { return logs.PwdModPath() },
		"HasService": func() bool { return len(serviceList) > 0 },
	})
	for i, v := range serviceList {
		serviceList[i] = strings.ToLower(v)
	}
	s.Render(path.Join("config/internal.yaml"), _config_template, serviceList, template.FuncMap{
		"PwdModPath": func() string { return logs.PwdModPath() },
		"HasService": func() bool { return len(serviceList) > 0 },
	})
}

const (
	_goapi_client = `
package api

import (
	"context"

	"{{ PwdModPath }}/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/grpc"
)
{{ range $index, $service := . }}
func New{{ $service }}Client(ctx context.Context, consul consul.Consul) (pb.{{ $service }}Client, error) {
	if conn, err := grpc.NewConn(ctx, consul.Address(pb.{{ $service }}_ServiceDesc.ServiceName)); err != nil {
		return nil, err
	} else {
		return pb.New{{ $service }}Client(conn), err
	}
}
{{ end }}
`
	_goapi_init = `
package api

import "shylinux.com/x/golang-story/src/project/server/infrastructure/container"

func Init(c *container.Container) {
{{ range $index, $service := . }}
	c.Provide(New{{ $service }}Client)
{{ end }}
}
`
)
