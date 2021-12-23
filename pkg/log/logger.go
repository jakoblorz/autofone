package log

import "go.uber.org/zap"

var (
	DefaultLogger, _ = zap.NewDevelopment()
)
