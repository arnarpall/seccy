package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Console() *Logger {
	lc := zap.NewDevelopmentConfig()
	lc.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	lc.DisableStacktrace = true
	lc.DisableCaller = true
	lc.EncoderConfig = zapcore.EncoderConfig{
		MessageKey: "Y",
	}
	lg, _ := lc.Build()
	logger := &Logger{
		SugaredLogger: lg.Sugar(),
	}
	return logger
}
