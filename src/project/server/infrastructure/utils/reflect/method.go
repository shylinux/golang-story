package reflect

import (
	"reflect"
	"unicode"
)

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
