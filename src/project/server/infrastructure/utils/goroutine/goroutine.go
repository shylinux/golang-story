package goroutine

import (
	"fmt"
	"path"
	"reflect"
	"runtime"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

type Pool struct {
}

func New() *Pool {
	return &Pool{}
}
func (s *Pool) Go(cb func() error) {
	p := reflect.ValueOf(cb).Pointer()
	file, line := runtime.FuncForPC(p).FileLine(p)
	name := fmt.Sprintf("%s:%d", path.Base(file), line)
	go func() {
		logs.Infof("goroutine %s", name)
		if err := cb(); err != nil {
			logs.Errorf("goroutine %s %s", name, err)
		}
	}()
}
