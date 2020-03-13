package log

import "go.uber.org/zap"

type Logger struct {
	*zap.SugaredLogger
}

func New() *Logger {
	c := zap.NewProductionConfig()
	c.DisableStacktrace = true
	l, _ := c.Build()
	return &Logger{
		SugaredLogger: l.Sugar(),
	}
}
