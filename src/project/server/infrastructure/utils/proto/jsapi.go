package proto

import (
	"os"
	"path"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func (s *Generate) GenJsAPI() {
	s.OpenProto(func(file *os.File, name string) {
		module, serviceList, serviceRequest, requestList := "", map[string][]string{}, map[string]string{}, map[string][]string{}
		block, service, request, comment, action := "", "", "", "", map[string]string{}
		s.ScanProto(file, func(ls []string, text string) {
			if strings.HasPrefix(text, PACKAGE) {
				module = strings.TrimSuffix(ls[1], ";")
			} else if strings.HasPrefix(text, SERVICE) {
				block, service = ls[0], ls[1]
			} else if strings.HasPrefix(text, MESSAGE) && strings.HasSuffix(text, "Request {") {
				block, request = ls[0], ls[1]
			} else if block == "" {

			} else if strings.HasPrefix(text, "}") {
				block = ""
			} else if strings.HasPrefix(text, "//") {
				comment = text
			} else {
				if comment != "" && block == MESSAGE {
					comment, action[ls[1]] = "", strings.TrimSpace(strings.TrimPrefix(comment, "//"))
				}
				if block == SERVICE {
					serviceList[service] = append(serviceList[service], ls[1])
					serviceRequest[ls[1]] = strings.TrimPrefix(strings.TrimSuffix(ls[2], ")"), "(")
				} else {
					requestList[request] = append(requestList[request], ls[1])
				}
			}
		})
		logs.Infof("  module: %v %+v", module, serviceList)
		s.Output(path.Join(s.Config.Generate.JsPath, name+".js"), func(echo func(string, ...interface{})) {
			echo(_jsapi_import)
			for service, api := range serviceList {
				echo(s.Template(_jsapi_header, map[string]string{"module": module, SERVICE: service}))
				for _, api := range api {
					echo(s.Template(_jsapi_template, map[string]string{
						"api": api, "params": strings.Join(requestList[serviceRequest[api]], ", "),
					}))
				}
				echo(_jsapi_footer)
			}
		})
	})
}

const (
	_jsapi_import = `
import request from '@/utils/request'

`
	_jsapi_header = `
export default class {{ .service }} {
  static path = '/api/{{ .module}}.{{ .service }}/'
`
	_jsapi_template = `
  static async {{ .api }}({{ .params }}) {
    return await request.post(this.path + '{{ .api }}', { {{ .params }} })
  }
`
	_jsapi_footer = `
}
`
)
