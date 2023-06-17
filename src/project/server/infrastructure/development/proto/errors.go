package proto

import (
	"path"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/reflect"
)

func (s *GenerateCmds) GenErrors() {
	s.Output(path.Join(s.conf.JsPath, "errors.js"), func(echo func(str string, arg ...interface{})) {
		reflect.FieldList(&enums.Errors, func(name string, field reflect.Field) {
			echo("export const %s = %d", name, field.Int())
		})
	})
}
