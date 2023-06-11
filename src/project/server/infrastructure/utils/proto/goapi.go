package proto

import (
	"os"
	"path"
	"strings"
)

func (s *Generate) GenGoAPI() {
	service := []string{}
	s.OpenProto(func(file *os.File, name string) {
		s.ScanProto(file, func(ls []string, text string) {
			if strings.HasPrefix(text, SERVICE) {
				service = append(service, ls[1])
				s.Render(path.Join(s.Config.Generate.GoPath, name+".go"), _goapi_client, map[string]string{SERVICE: ls[1]})
			}
		})
	})
	s.Render(path.Join(s.Config.Generate.GoPath, "api.go"), _goapi_init, map[string][]string{SERVICE: service})
}

const (
	_goapi_client = `
package api

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
)

func New{{ .service }}Client(ctx context.Context, consul consul.Consul) (pb.{{ .service }}Client, error) {
	if conn, err := grpc.NewConn(ctx, consul.Address(pb.{{ .service }}_ServiceDesc.ServiceName)); err != nil {
		return nil, err
	} else {
		client := pb.New{{ .service }}Client(conn)
		return client, err
	}
}
`
	_goapi_init = `
package api

import "shylinux.com/x/golang-story/src/project/server/infrastructure/container"

func Init(container *container.Container) {
{{ range $index, $item := .service }}
	container.Provide(New{{ $item }}Client)
{{ end }}
}
`
)
