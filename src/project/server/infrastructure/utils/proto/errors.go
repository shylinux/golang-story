package proto

import (
	"path"
	"reflect"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
)

func (s *Generate) GenErrors() {
	t, v := reflect.TypeOf(enums.Errors), reflect.ValueOf(enums.Errors)
	s.Output(path.Join(s.conf.JsPath, "errors.js"), func(echo func(str string, arg ...interface{})) {
		for i := 0; i < v.NumField(); i++ {
			echo("export const %s = %d", t.Field(i).Name, v.Field(i).Int())
		}
	})
}
