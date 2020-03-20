package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func (l *Logger) With(context... interface{}) *Logger {
	return &Logger{
		SugaredLogger: l.SugaredLogger.With(context...),
	}
}

func NewWithContext(context... interface{}) *Logger {
	return New().With(context)
}

func New() *Logger {
	c := zap.NewProductionConfig()
	c.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	c.DisableStacktrace = true
	l, _ := c.Build()
	return &Logger{
		SugaredLogger: l.Sugar(),
	}
}
