package reflect

import (
	"reflect"
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
