package log

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/natefinch/lumberjack"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
)

type Logger interface {
	Infof(str string, arg ...interface{})
	Warnf(str string, arg ...interface{})
	With(args ...interface{}) Logger
}
type logger struct{ *zap.SugaredLogger }

func (s *logger) format(str string, arg ...interface{}) string {
	if len(arg) > 0 {
		if ctx, ok := arg[len(arg)-1].(context.Context); ok {
			if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
				return fmt.Sprintf("traceID: %s ", span.TraceID().String()) + fmt.Sprintf(str, arg[:len(arg)-1]...)
			}
		}
	}
	return fmt.Sprintf(str, arg...)
}
func (s *logger) Infof(str string, arg ...interface{}) { s.SugaredLogger.Infof(s.format(str, arg...)) }
func (s *logger) Warnf(str string, arg ...interface{}) { s.SugaredLogger.Warnf(s.format(str, arg...)) }
func (s *logger) With(arg ...interface{}) Logger {
	return &logger{s.SugaredLogger.With(arg...).WithOptions(zap.AddCallerSkip(1))}
}

var l *logger

func Infof(str string, arg ...interface{})  { l.Infof(str, arg...) }
func Warnf(str string, arg ...interface{})  { l.Warnf(str, arg...) }
func Fatalf(str string, arg ...interface{}) { l.Fatalf(str, arg...) }
func Fatalln(arg ...interface{})            { l.Fatalln(arg...) }
func With(arg ...interface{}) Logger        { return l.With(arg...) }

func Debugf(str string, arg ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.Infof(fmt.Sprintf("%s:%d ", file, line) + fmt.Sprintf(str, arg...))
}

func New(config *config.Config) (Logger, error) {
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	l = &logger{zap.New(zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stderr), zap.InfoLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(&lumberjack.Logger{Filename: config.LogPath}), zap.InfoLevel),
	), zap.AddCaller(), zap.AddCallerSkip(2)).Sugar()}
	return l, nil
}
