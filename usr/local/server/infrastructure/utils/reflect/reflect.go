package reflect

import (
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func Bind(req interface{}, arg ...string) interface{} {
	if len(arg) == 0 {
		return req
	}
	rt, rv := reflect.TypeOf(req).Elem(), reflect.ValueOf(req).Elem()
	trans := map[string]string{}
	for i := 0; i < rv.NumField(); i++ {
		if unicode.IsUpper(rune(rt.Field(i).Name[0])) {
			trans[strings.ToLower(rt.Field(i).Name)] = rt.Field(i).Name
		}
	}
	for i := 0; i < len(arg); i += 2 {
		if fv := rv.FieldByName(trans[strings.ToLower(arg[i])]); fv.CanSet() {
			switch fv.Type().Kind() {
			case reflect.String:
				fv.SetString(arg[i+1])
			case reflect.Int64:
				v, _ := strconv.ParseInt(arg[i+1], 10, 64)
				fv.SetInt(v)
			}
		}
	}
	return req
}
func Trans(dst interface{}, src interface{}) interface{} {
	if src == nil {
		return dst
	}
	t, v := reflect.TypeOf(src).Elem(), reflect.ValueOf(src).Elem()
	if !v.IsValid() {
		return dst
	}
	call := errors.FileLine(2)
	FieldList(dst, func(name string, field Field) {
		if _, ok := t.FieldByName(name); ok {
			switch field.Kind() {
			case reflect.String:
				field.SetString(v.FieldByName(name).String())
			case reflect.Int64:
				field.SetInt(v.FieldByName(name).Int())
			case reflect.Int32:
				field.SetInt(v.FieldByName(name).Int())
			default:
				logs.Errorf("not implement covert %s %s", name, call)
			}
		} else {
			logs.Errorf("not found field %s %s", name, call)
		}
	})
	return dst
}
func TransList(list interface{}, trans interface{}, data interface{}) {
	v := reflect.ValueOf(list)
	cb := reflect.ValueOf(trans)
	target := reflect.ValueOf(data).Elem()
	for i := 0; i < v.Len(); i++ {
		res := cb.Call([]reflect.Value{v.Index(i)})
		target.Set(reflect.Append(target, res[0]))
	}
}
