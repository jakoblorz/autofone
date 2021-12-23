package log

import (
	"go.uber.org/zap"
)

var (
	DefaultLogger, _ = zap.NewDevelopment()
)

func Verbosef(format string, args ...interface{}) {
	DefaultLogger.Sugar().Debugf(format, args...)
}

func Printf(format string, args ...interface{}) {
	DefaultLogger.Sugar().Infof(format, args...)
}

func Print(args ...interface{}) {
	DefaultLogger.Sugar().Info(args...)
}
