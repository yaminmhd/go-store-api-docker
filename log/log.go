package log

import (
	"github.com/sirupsen/logrus"
	"github.com/yaminmhd/go-hardware-store/config"
	"net/http"
	"os"
)

type Logger struct {
	*logrus.Logger
}

var Log *Logger

type LoggerError struct {
	Error error
}

func panicIfError(err error) {
	if err != nil {
		panic(LoggerError{err})
	}
}

func SetupLogger() {
	level, err := logrus.ParseLevel(config.LogLevel())
	panicIfError(err)

	logrusVar := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}

	Log = &Logger{logrusVar}
}

func BuildContext(context string) logrus.Fields {
	return logrus.Fields{
		"context": context,
	}
}

func (logger *Logger) Errorf(r *http.Request, format string, args ...interface{}) {
	logger.httpRequestLogEntry(r).Errorf(format, args...)
}

func (logger *Logger) Infof(r *http.Request, format string, args ...interface{}) {
	logger.httpRequestLogEntry(r).Infof(format, args...)
}

func (logger *Logger) Warnf(r *http.Request, format string, args ...interface{}) {
	logger.httpRequestLogEntry(r).Warnf(format, args...)
}

func (logger *Logger) httpRequestLogEntry(r *http.Request) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"RequestMethod": r.Method,
		"Path":          r.URL.Path,
	})
}