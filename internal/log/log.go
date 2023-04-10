package log

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var sugaredLogger *zap.SugaredLogger

func InitZapLogger() {
	env := viper.GetString("misc.env")
	logger, _ := zap.NewDevelopment()
	if env == "prod" {
		logger, _ = zap.NewProduction()
	}
	sugaredLogger = logger.Sugar()
}

func Info(args ...interface{}) {
	sugaredLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	sugaredLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	sugaredLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	sugaredLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	sugaredLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	sugaredLogger.Errorf(template, args...)
}

func Debug(args ...interface{}) {
	sugaredLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	sugaredLogger.Debugf(template, args...)
}

func Fatal(args ...interface{}) {
	sugaredLogger.Debug(args...)
}

func Fatalf(template string, args ...interface{}) {
	sugaredLogger.Debugf(template, args...)
}
