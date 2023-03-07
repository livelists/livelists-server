package logger

import (
	"fmt"
	"github.com/go-logr/logr"
	"go.uber.org/zap/zapcore"
)

var (
	defaultLogger = logr.Discard()
	pkgLogger     = Logger(logr.Discard().V(10))
)

// GetLogger returns the logger that was set with SetLogger with an extra depth of 1
func GetLogger() logr.Logger {
	return defaultLogger
}

// GetDefaultLogger returns the logger that was set but with LiveKit wrappers
func GetDefaultLogger() Logger {
	return Logger(defaultLogger)
}

// SetLogger lets you use a custom logger. Pass in a logr.Logger with default depth
func SetLogger(l logr.Logger, name string) {
	defaultLogger = l.WithCallDepth(1).WithName(name)
	// pkg wrapper needs to drop two levels of depth
	pkgLogger = Logger(l.WithCallDepth(2).WithName(name))
}

func Debugw(msg string, keysAndValues ...interface{}) {
	pkgLogger.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	fmt.Println(pkgLogger)
	pkgLogger.Infow(msg, keysAndValues...)
}

func Warnw(msg string, err error, keysAndValues ...interface{}) {
	pkgLogger.Warnw(msg, err, keysAndValues...)
}

func Errorw(msg string, err error, keysAndValues ...interface{}) {
	pkgLogger.Errorw(msg, err, keysAndValues...)
}

func ParseZapLevel(level string) zapcore.Level {
	lvl := zapcore.InfoLevel
	if level != "" {
		_ = lvl.UnmarshalText([]byte(level))
	}
	return lvl
}

type Logger logr.Logger

func (l Logger) toLogr() logr.Logger {
	if logr.Logger(l).GetSink() == nil {
		return defaultLogger
	}
	return logr.Logger(l)
}

func (l Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.toLogr().V(1).Info(msg, keysAndValues...)
}

func (l Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.toLogr().Info(msg, keysAndValues...)
}

func (l Logger) Warnw(msg string, err error, keysAndValues ...interface{}) {
	if err != nil {
		keysAndValues = append(keysAndValues, "error", err)
	}
	l.toLogr().Info(msg, keysAndValues...)
}

func (l Logger) Errorw(msg string, err error, keysAndValues ...interface{}) {
	l.toLogr().Error(err, msg, keysAndValues...)
}
