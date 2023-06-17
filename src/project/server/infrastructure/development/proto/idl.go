package proto

const (
	_config_template = `
internal:
{{ range $index, $service := . }}  {{ $service }}:
    export: true
{{ end }}
`
	_idl_template = `
package idl

import (
{{ if HasService }}
	"{{ PwdModPath }}/controller"
	"{{ PwdModPath }}/service"
{{ end }}
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
)

func Init(c *container.Container) {
	c.Provide(NewMainController)
{{ range $index, $service := . }}
	c.Provide(controller.New{{ $service }}Controller)
	c.Provide(service.New{{ $service }}Service)
{{ end }}
}

type MainController struct{}

func NewMainController(
{{ range $index, $service := . }}
	_ *controller.{{ $service }}Controller,
{{ end }}
) *MainController {
	return &MainController{}
}
`
)
