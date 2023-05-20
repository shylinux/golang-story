package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
)

type Logger interface {
	Write(buf []byte) (int, error)
	Infof(str string, arg ...interface{})
	Printf(str string, arg ...interface{})
}
type logger struct {
	*zap.SugaredLogger
}

func (s *logger) Write(buf []byte) (int, error) {
	s.Info(string(buf))
	return len(buf), nil
}
func (s *logger) Printf(str string, arg ...interface{}) { s.Infof(str, arg...) }

var l logger

func Infof(str string, arg ...interface{}) { l.Infof(str, arg...) }

func NewLogger(config *config.Config) (Logger, error) {
	consoleErrors := zapcore.Lock(os.Stderr)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return true })
	return &logger{zap.New(zapcore.NewCore(consoleEncoder, consoleErrors, highPriority)).Sugar()}, nil
}
