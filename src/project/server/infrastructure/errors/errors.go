package errors

import (
	"fmt"
	"runtime"
	"strings"
)

type errorResp struct {
	last error
	code int64
	info string
}

func NewResp(err error, code int64, str string, arg ...interface{}) error {
	if err == nil {
		return nil
	}
	return &errorResp{code: code, info: fmt.Sprintf(str, arg...)}
}
func (s *errorResp) Error() string {
	return fmt.Sprintf("%d: %s %s", s.code, s.info, s.last)
}

type errors struct {
	last     error
	info     string
	fileline string
}

func New(err error, str string, arg ...interface{}) error {
	switch err.(type) {
	case nil:
		return nil
	case *errorResp:
		return err
	}
	return &errors{last: err, info: fmt.Sprintf(str, arg...), fileline: FileLine(1)}
}
func (s *errors) Error() string {
	return fmt.Sprintf("%s %s %s", s.info, s.fileline, s.last)
}

func FileLine(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	list := strings.Split(file, "/")
	if len(list) > 3 {
		list = list[len(list)-3:]
	}
	return fmt.Sprintf("%s:%s", strings.Join(list, "/"), line)
}
