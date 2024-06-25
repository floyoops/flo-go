package logger

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	sugarLogger *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
	logger, _ := zap.NewProduction()
	sugarLogger := logger.Sugar()
	return &ZapLogger{sugarLogger: sugarLogger}
}

func (l *ZapLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.sugarLogger.Infof(format, args...)
}

func (l *ZapLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *ZapLogger) Warnf(format string, args ...interface{}) {
	l.sugarLogger.Warnf(format, args...)
}

func (l *ZapLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.sugarLogger.Errorf(format, args...)
}
