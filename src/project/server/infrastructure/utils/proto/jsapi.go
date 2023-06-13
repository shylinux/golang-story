package proto

import (
	"html/template"
	"path"
	"strings"
)

func (s *Generate) GenJsAPI() {
	for name, proto := range s.protos {
		s.Render(path.Join(s.conf.JsPath, name+".js"), _jsapi_client, proto, template.FuncMap{
			"PackageName": func() string { return proto[PACKAGE].Name },
			"ServiceList": func() []string { return proto[PACKAGE].List },
			"MethodList":  func(service string) []string { return proto[service].List },
			"MethodParams": func(method string) string {
				list, params := []string{}, proto[proto[method].List[0]]
				for i := 0; i < len(params.List); i += 3 {
					list = append(list, params.List[i+2])
				}
				return strings.Join(list, ", ")
			},
		})
	}
}

const (
	_jsapi_client = `
import request from '@/utils/request'
{{ range $index, $service := ServiceList }}
export class {{ $service }} {
  static path = '/api/{{ PackageName }}.{{ $service }}/'
{{ range $index, $method := MethodList $service }}
  static async {{ $method }}({{ MethodParams $method }}) {
    return await request.post(this.path + '{{ $method }}', { {{ MethodParams $method }} })
  }
{{ end }}
}
{{ end }}
`
)
