package log

import "go.uber.org/zap"

type Logger struct {
	*zap.SugaredLogger
}

func New() *Logger {
	logger, _ := zap.NewProduction()
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}
