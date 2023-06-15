package reflect

import (
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

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
