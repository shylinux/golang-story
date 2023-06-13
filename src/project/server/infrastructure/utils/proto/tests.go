package proto

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

func (s *Generate) GenTests() {
	for name, proto := range s.protos {
		s.Render(path.Join(s.conf.TsPath, name+"_test.go"), _ts_cases, proto, template.FuncMap{
			"PwdModPath":  func() string { return logs.PwdModPath() },
			"ServiceList": func() []string { return proto[PACKAGE].List },
			"MethodList": func(service string) map[string]string {
				res := map[string]string{}
				for _, method := range proto[service].List {
					res[method] = proto[method].List[0]
				}
				return res
			},
			"CaseStruct": func(method string) interface{} {
				res := []string{"		OK bool `yaml:\"ok\"`"}
				list := proto[proto[method].List[0]].List
				for i := 0; i < len(list); i += 3 {
					res = append(res, fmt.Sprintf("		%s %s `yaml:\"%s\"`", Capital(list[i+2]), list[i+1], list[i+2]))
				}
				return template.HTML(strings.Join(res, "\n"))
			},
			"CaseConfig": func(service, method string) interface{} {
				p := fmt.Sprintf("testdata/%s/%s.yaml", service, method)
				if _, e := os.Stat(path.Join(s.conf.TsPath, p)); os.IsNotExist(e) {
					system.WriteFile(path.Join(s.conf.TsPath, p), []byte("- ok: false\n"), 0644)
				}
				return template.HTML(fmt.Sprintf("\"%s\"", p))
			},
			"CaseParams": func(method string) string {
				res := []string{}
				list := proto[proto[method].List[0]].List
				for i := 0; i < len(list); i += 3 {
					res = append(res, fmt.Sprintf("%s: c.%s", Capital(list[i+2]), Capital(list[i+2])))
				}
				return strings.Join(res, ", ")
			},
		})
	}
}

const (
	_ts_cases = `
package test

import (
	"context"
	"testing"

	"{{ PwdModPath }}/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/tests"
)
{{ range $index, $service := ServiceList }}
type {{ $service }}TestSuite struct {
	*tests.Suite
	ctx      context.Context
	client   pb.{{ $service }}Client
}

func (s *{{ $service }}TestSuite) SetupTest() {
	s.client = pb.New{{ $service }}Client(s.Conn(s.ctx, pb.{{ $service }}_ServiceDesc.ServiceName))
}
{{ range $method, $request := MethodList $service }}
func (s *{{ $service }}TestSuite) Test{{ $method }}() {
	cases := []struct {
{{ CaseStruct $method }}
	}{}
	s.Load({{ CaseConfig $service $method }}, &cases)
	for i, c := range cases {
		res, err := s.client.{{ $method }}(s.ctx, &pb.{{ $request }}{ {{ CaseParams $method }} })
		s.ConveySo(i, c.OK, c, res, err)
	}
}
{{ end }}
func Test{{ $service }}TestSuite(t *testing.T) {
	infrastructure.Test(t, func(suite *tests.Suite) interface{} {
		return &{{ $service }}TestSuite{Suite: suite, ctx: suite.Context()}
	})
}
{{ end }}
`
)
