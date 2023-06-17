package errors

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
)

type ErrorResp struct {
	Code     int64
	Info     string
	fileline string
	funcname string
	last     error
}

func (s *ErrorResp) Error() string {
	if s == nil {
		return ""
	}
	return fmt.Sprintf("%d: %s %s:%s %s", s.Code, s.Info, s.fileline, s.funcname, s.last.Error())
}
func (s *ErrorResp) ToGRPC() error {
	if s == nil {
		return nil
	}
	return status.Error(codes.Code(s.Code), fmt.Sprintf("%s %s:%s %s", s.Info, s.fileline, s.funcname, s.last.Error()))
}
func newResp(err error, code int64, str string, arg ...interface{}) *ErrorResp {
	switch err := err.(type) {
	case nil:
		return nil
	case *ErrorResp:
		return err
	}
	return &ErrorResp{Code: code, Info: fmt.Sprintf(str, arg...), fileline: FileLine(3), funcname: FuncName(3), last: err}
}
func ParseResp(err error, str string) *ErrorResp {
	switch err := err.(type) {
	case nil:
		return nil
	case *ErrorResp:
		return err
	}
	if ls := strings.SplitN(err.Error(), ":", 2); len(ls) > 1 {
		if code, _err := strconv.ParseInt(ls[0], 10, 64); _err == nil && code > 0 {
			return newResp(err, code, ls[1])
		}
	}
	return newResp(err, enums.Errors.Unknown, str)
}
func NewResp(err error, code int64, str string, arg ...interface{}) *ErrorResp {
	return newResp(err, code, str, arg...)
}
func NewNotAuth(err error) *ErrorResp {
	return newResp(err, enums.Errors.AuthFailure, "auth failure")
}
func NewNotFoundProxy(err error) *ErrorResp {
	return newResp(err, enums.Errors.NotFoundProxy, "not found proxy")
}
func NewInvalidParams(err error) *ErrorResp {
	return newResp(err, enums.Errors.InvalidParams, "invalid params")
}
func NewNotFoundUser(err error) *ErrorResp {
	return newResp(err, enums.Errors.NotFoundUser, "not found user")
}
func NewIncorrectPassword(err error) *ErrorResp {
	return newResp(err, enums.Errors.IncorrectPassword, "incorrect password")
}
func NewAlreadyExists(err error) *ErrorResp {
	return newResp(err, enums.Errors.AlreadyExists, "already exists")
}
func NewCreateFailResp(err error) *ErrorResp {
	return newResp(err, enums.Errors.ModelCreate, "service create failure")
}
func NewRemoveFailResp(err error) *ErrorResp {
	return newResp(err, enums.Errors.ModelRemove, "service remove failure")
}
func NewModifyFailResp(err error) *ErrorResp {
	return newResp(err, enums.Errors.ModelModify, "service modify failure")
}
func NewSearchFailResp(err error) *ErrorResp {
	return newResp(err, enums.Errors.ModelSearch, "service search failure")
}
func NewInfoFailResp(err error) *ErrorResp {
	return newResp(err, enums.Errors.ModelInfo, "service info failure")
}
func NewListFailResp(err error) *ErrorResp {
	return newResp(err, enums.Errors.ModelList, "service list failure")
}

type errors struct {
	info     string
	fileline string
	funcname string
	last     error
}

func (s *errors) Error() string {
	return fmt.Sprintf("%s %s:%s %s", s.info, s.fileline, s.funcname, s.last.Error())
}
func newError(err error, str string, arg ...interface{}) error {
	switch err.(type) {
	case nil:
		return nil
	case *ErrorResp:
		return err
	}
	return &errors{info: fmt.Sprintf(str, arg...), fileline: FileLine(3), funcname: FuncName(3), last: err}
}
func New(err error, str string, arg ...interface{}) error {
	return newError(err, str, arg...)
}
func NewCreateFail(err error) error { return newError(err, "storage create failure") }
func NewRemoveFail(err error) error { return newError(err, "storage remove failure") }
func NewModifyFail(err error) error { return newError(err, "storage modify failure") }
func NewSearchFail(err error) error { return newError(err, "storage search failure") }
func NewInfoFail(err error) error   { return newError(err, "storage info failure") }
func NewListFail(err error) error   { return newError(err, "storage list failure") }

func Assert(err error) {
	if err == nil {
		return
	}
	panic(err)
}
func FileLine(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	list := strings.Split(file, "/")
	if len(list) > 2 {
		list = list[len(list)-2:]
	}
	return fmt.Sprintf("%s:%d", path.Join(list[:]...), line)
}
func FuncName(skip int) string {
	fun, _, _, _ := runtime.Caller(skip)
	return path.Base(runtime.FuncForPC(fun).Name())
}
