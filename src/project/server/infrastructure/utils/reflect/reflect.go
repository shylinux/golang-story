package reflect

import (
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type Field struct{ reflect.Value }

func FieldList(target interface{}, cb func(string, Field)) {
	t, v := reflect.TypeOf(target).Elem(), reflect.ValueOf(target).Elem()
	for i := 0; i < v.NumField(); i++ {
		if unicode.IsLower(rune(t.Field(i).Name[0])) {
			continue
		}
		cb(t.Field(i).Name, Field{v.Field(i)})
	}
}

type Method struct{ reflect.Value }

func MethodList(target interface{}, cb func(string, Method)) {
	t, v := reflect.TypeOf(target), reflect.ValueOf(target)
	for i := 0; i < v.NumMethod(); i++ {
		if unicode.IsLower(rune(t.Method(i).Name[0])) {
			continue
		}
		cb(t.Method(i).Name, Method{v.Method(i)})
	}
}

func (s Method) Call(arg ...interface{}) (res []interface{}) {
	args := []reflect.Value{}
	for _, v := range arg {
		args = append(args, reflect.ValueOf(v))
	}
	list := s.Value.Call(args)
	for _, v := range list {
		res = append(res, v.Interface())
	}
	return
}
func (s Method) NewParam(n int) interface{} {
	return reflect.New(s.Type().In(n).Elem()).Interface()
}
func (s Method) NewResult(n int) interface{} {
	return reflect.New(s.Type().Out(n).Elem()).Interface()
}

func Bind(req interface{}, arg ...string) interface{} {
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
