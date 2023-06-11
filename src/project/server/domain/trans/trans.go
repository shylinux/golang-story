package trans

import (
	"reflect"
)

func ListDTO(list interface{}, trans interface{}, data interface{}) {
	v := reflect.ValueOf(list)
	cb := reflect.ValueOf(trans)
	target := reflect.ValueOf(data).Elem()
	for i := 0; i < v.Len(); i++ {
		res := cb.Call([]reflect.Value{v.Index(i)})
		target.Set(reflect.Append(target, res[0]))
	}
}
