package errors

import (
	"fmt"
	"runtime"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
)

type errorResp struct {
	code     int
	info     string
	fileline string
	last     error
}

func (s *errorResp) Error() string {
	return fmt.Sprintf("%d: %s %s %s", s.code, s.info, s.fileline, s.last.Error())
}
func newResp(err error, code int, str string, arg ...interface{}) error {
	if err == nil {
		return nil
	}
	return &errorResp{code: code, info: fmt.Sprintf(str, arg...), fileline: FileLine(3), last: err}
}
func NewResp(err error, code int, str string, arg ...interface{}) error {
	return newResp(err, code, str, arg...)
}
func NewCreateFailResp(err error) error {
	return newResp(err, enums.Errors.ModelCreate, "model create failure")
}
func NewRemoveFailResp(err error) error {
	return newResp(err, enums.Errors.ModelRemove, "model remove failure")
}
func NewInfoFailResp(err error) error {
	return newResp(err, enums.Errors.ModelInfo, "model info failure")
}
func NewListFailResp(err error) error {
	return newResp(err, enums.Errors.ModelList, "model list failure")
}

type errors struct {
	info     string
	fileline string
	last     error
}

func (s *errors) Error() string {
	return fmt.Sprintf("%s %s %s", s.info, s.fileline, s.last.Error())
}
func newError(err error, str string, arg ...interface{}) error {
	switch err.(type) {
	case nil:
		return nil
	case *errorResp:
		return err
	}
	return &errors{last: err, info: fmt.Sprintf(str, arg...), fileline: FileLine(3)}
}
func New(err error, str string, arg ...interface{}) error {
	return newError(err, str, arg...)
}
func NewCreateFail(err error) error { return newError(err, "storage create failure") }
func NewRemoveFail(err error) error { return newError(err, "storage remove failure") }
func NewInfoFail(err error) error   { return newError(err, "storage info failure") }
func NewListFail(err error) error   { return newError(err, "storage list failure") }

func FileLine(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	list := strings.Split(file, "/")
	if len(list) > 2 {
		list = list[len(list)-2:]
	}
	return fmt.Sprintf("%s:%d", strings.Join(list, "/"), line)
}
