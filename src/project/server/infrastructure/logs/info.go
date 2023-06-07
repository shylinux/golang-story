package logs

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"
)

func FuncName(skip int) string {
	fun, _, _, _ := runtime.Caller(skip)
	return path.Base(runtime.FuncForPC(fun).Name())
}
func FileLine(skip interface{}) string {
	file, line := "", 0
	switch skip := skip.(type) {
	case int:
		_, file, line, _ = runtime.Caller(skip)
	case uintptr:
		file, line = runtime.FuncForPC(skip).FileLine(skip)
	default:
		fun := reflect.ValueOf(skip).Pointer()
		file, line = runtime.FuncForPC(fun).FileLine(fun)
	}
	list := strings.Split(file, "/")
	if len(list) > 2 {
		list = list[len(list)-2:]
	}
	return fmt.Sprintf("%s:%d", path.Join(list[:]...), line)
}
