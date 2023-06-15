package proto

import (
	"html/template"
	"path"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func (s *Generate) GenShCLI() {
	serviceList := []string{}
	for name, proto := range s.protos {
		serviceList = append(serviceList, proto[PACKAGE].List...)
		s.Render(path.Join(s.conf.ShPath, name+".go"), _shcmd_client, proto, template.FuncMap{
			"PwdModPath":  func() string { return logs.PwdModPath() },
			"ServiceList": func() []string { return proto[PACKAGE].List },
			"ServiceCmds": func(service string) string { return strings.ToLower(strings.TrimSuffix(service, "Service")) },
			"ServiceHelp": func(service string) string {
				return strings.ToLower(strings.TrimSuffix(service, "Service")) + " service client"
			},
			"MethodList":    func(service string) []string { return proto[service].List },
			"MethodRequest": func(method string) string { return proto[method].List[0] },
		})
	}
	s.Render(path.Join(s.conf.ShPath, path.Base(s.conf.ShPath)+".go"), _shcmd_init, serviceList, nil)
}

const (
	_shcmd_client = `
package cli

import (
	"context"
	"fmt"

	"{{ PwdModPath }}/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/grpc"
)
{{ range $index, $service := ServiceList }}
type {{ $service }}Cmds struct {
	consul consul.Consul
	client pb.{{ $service }}Client
}

func New{{ $service }}Cmds(cmds *cmds.Cmds, consul consul.Consul) (*{{ $service }}Cmds, error) {
	_cmds := &{{ $service }}Cmds{consul: consul}
	cmds.Register("{{ ServiceCmds $service }}", "{{ ServiceHelp $service }}", _cmds)
	return _cmds, nil
}

func (s *{{ $service }}Cmds) conn(ctx context.Context, arg ...string) {
	if s.client != nil {
		return
	}
	conn, err := grpc.NewConn(ctx, s.consul.Address(pb.{{ $service }}_ServiceDesc.ServiceName))
	if err != nil {
		return
	}
	s.client = pb.New{{ $service }}Client(conn)
}
{{ range $index, $method := MethodList $service }}
func (s *{{ $service }}Cmds) {{ $method }}(ctx context.Context, req *pb.{{ MethodRequest $method }}) {
	s.conn(ctx)
	if res, err := s.client.{{ $method }}(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}
{{ end }}
{{ end }}
`
	_shcmd_init = `
package cli

import (
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
)

func Init(container *container.Container) {
	container.Provide(NewMainServiceCmds)
{{ range $index, $service := . }}
	container.Provide(New{{ $service }}Cmds)
{{ end }}
}

type MainServiceCmds struct {
	*cmds.Cmds
}

func NewMainServiceCmds(cmds *cmds.Cmds{{ range $index, $service := . }}, _ *{{ $service }}Cmds{{ end }}) *MainServiceCmds {
	return &MainServiceCmds{cmds}
} 
`
)
