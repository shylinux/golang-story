package logs

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
)

type Logger interface {
	With(args ...interface{}) Logger
	Infof(str string, arg ...interface{})
	Warnf(str string, arg ...interface{})
	Errorf(str string, arg ...interface{})
	Debugf(str string, arg ...interface{})
}
type logger struct{ *zap.SugaredLogger }

func (s *logger) With(arg ...interface{}) Logger {
	return &logger{s.SugaredLogger.With(arg...).WithOptions(zap.AddCallerSkip(1))}
}
func (s *logger) Infof(str string, arg ...interface{}) {
	s.SugaredLogger.Infof(s.format(str, arg...))
}
func (s *logger) Warnf(str string, arg ...interface{}) {
	s.SugaredLogger.Warnf(s.format(str, arg...))
}
func (s *logger) Errorf(str string, arg ...interface{}) {
	s.SugaredLogger.Errorf(s.format(str, arg...))
}
func (s *logger) Debugf(str string, arg ...interface{}) {
	s.SugaredLogger.Debugf(s.format(str, arg...))
}
func (s *logger) format(str string, arg ...interface{}) string {
	if len(arg) > 0 {
		if ctx, ok := arg[len(arg)-1].(context.Context); ok {
			if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
				return fmt.Sprintf(" %s-%s ", span.TraceID().String(), span.SpanID()) + fmt.Sprintf(str, arg[:len(arg)-1]...)
			}
		}
	}
	return fmt.Sprintf(str, arg...)
}

var l *logger
var log *logger

func New(config *config.Config) (Logger, error) {
	conf := config.Log
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	tees := []zapcore.Core{}
	if conf.Stdout || conf.Path == "" {
		tees = append(tees, zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stderr), zap.InfoLevel))
	}
	if conf.Path != "" {
		tees = append(tees, zapcore.NewCore(consoleEncoder, zapcore.AddSync(&lumberjack.Logger{
			Filename: conf.Path, MaxSize: conf.MaxSize, MaxAge: conf.MaxAge, LocalTime: true,
		}), zap.InfoLevel))
	}
	log = &logger{zap.New(zapcore.NewTee(tees...), zap.AddCaller()).Sugar()}
	l = &logger{log.SugaredLogger.WithOptions(zap.AddCallerSkip(2))}
	if conf.Pid != "" {
		ioutil.WriteFile(conf.Pid, []byte(fmt.Sprintf("%d", os.Getpid())), 0755)
	}
	return l, nil
}
func With(arg ...interface{}) Logger        { return log.With(arg...) }
func Infof(str string, arg ...interface{})  { l.Infof(str, arg...) }
func Warnf(str string, arg ...interface{})  { l.Warnf(str, arg...) }
func Errorf(str string, arg ...interface{}) { l.Errorf(str, arg...) }
func Fatalf(str string, arg ...interface{}) { l.Fatalf(str, arg...) }
func Fatalln(arg ...interface{})            { l.Fatalln(arg...) }
func Debugf(str string, arg ...interface{}) { l.Debugf(str, arg) }

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
