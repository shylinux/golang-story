package errors

import (
	"fmt"

	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

type ErrorResp struct {
	Code     int64
	Info     string
	fileline string
	last     error
}

func (s *ErrorResp) Error() string {
	return fmt.Sprintf("%d: %s %s %s", s.Code, s.Info, s.fileline, s.last.Error())
}
func newResp(err error, code int64, str string, arg ...interface{}) error {
	switch err.(type) {
	case nil:
		return nil
	case *ErrorResp:
		return err
	}
	return &ErrorResp{Code: code, Info: fmt.Sprintf(str, arg...), fileline: logs.FileLine(3), last: err}
}
func NewResp(err error, code int64, str string, arg ...interface{}) error {
	return newResp(err, code, str, arg...)
}
func NewNotFoundProxy(err error) error {
	return newResp(err, enums.Errors.NotFoundProxy, "not found proxy")
}
func NewInvalidParams(err error) error {
	return newResp(err, enums.Errors.InvalidParams, "invalid params")
}
func NewCreateFailResp(err error) error {
	return newResp(err, enums.Errors.ModelCreate, "model create failure")
}
func NewRemoveFailResp(err error) error {
	return newResp(err, enums.Errors.ModelRemove, "model remove failure")
}
func NewModifyFailResp(err error) error {
	return newResp(err, enums.Errors.ModelModify, "model modify failure")
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
	case *ErrorResp:
		return err
	}
	return &errors{last: err, info: fmt.Sprintf(str, arg...), fileline: logs.FileLine(3)}
}
func New(err error, str string, arg ...interface{}) error {
	return newError(err, str, arg...)
}
func NewCreateFail(err error) error { return newError(err, "storage create failure") }
func NewRemoveFail(err error) error { return newError(err, "storage remove failure") }
func NewModifyFail(err error) error { return newError(err, "storage modify failure") }
func NewInfoFail(err error) error   { return newError(err, "storage info failure") }
func NewListFail(err error) error   { return newError(err, "storage list failure") }
