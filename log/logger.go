package log

import (
	"go.uber.org/zap"
)

var mLogger *zap.SugaredLogger

func Init() {
	logger, _ := zap.NewProduction()
	mLogger = logger.Sugar()
	defer logger.Sync()
}

func Debug(args ...interface{}) {
	mLogger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	mLogger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	mLogger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	mLogger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	mLogger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	mLogger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	mLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	mLogger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	mLogger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	mLogger.Fatalf(format, args...)
}
