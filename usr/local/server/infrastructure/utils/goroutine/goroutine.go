package goroutine

import (
	"fmt"
	"path"
	"reflect"
	"runtime"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

type Pool struct{}

func New() *Pool {
	return &Pool{}
}
func (s *Pool) Go(name string, cb func() error) {
	p := reflect.ValueOf(cb).Pointer()
	file, line := runtime.FuncForPC(p).FileLine(p)
	fileline := fmt.Sprintf("%s:%d", path.Base(file), line)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.Errorf("goroutine %s %s", fileline, err)
			}
		}()
		logs.Infof("goroutine %s %s", name, fileline)
		if err := cb(); err != nil {
			logs.Errorf("goroutine %s %s", fileline, err)
		}
	}()
}
