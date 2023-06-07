package pulsar

import "github.com/apache/pulsar-client-go/pulsar/log"

type logger struct{}

func (l *logger) SubLogger(fields log.Fields) log.Logger             { return l }
func (l *logger) WithFields(fields log.Fields) log.Entry             { return l }
func (l *logger) WithField(name string, value interface{}) log.Entry { return l }
func (l *logger) WithError(err error) log.Entry                      { return l }
func (l *logger) Info(args ...interface{})                           {}
func (l *logger) Warn(args ...interface{})                           {}
func (l *logger) Error(args ...interface{})                          {}
func (l *logger) Debug(args ...interface{})                          {}
func (l *logger) Infof(format string, args ...interface{})           {}
func (l *logger) Warnf(format string, args ...interface{})           {}
func (l *logger) Errorf(format string, args ...interface{})          {}
func (l *logger) Debugf(format string, args ...interface{})          {}
