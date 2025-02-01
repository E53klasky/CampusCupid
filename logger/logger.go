package logger

import (
	"fmt"

	"github.com/achintya-7/dating-app/constants"
	"go.uber.org/zap"
)

type LoggerContext interface {
	Value(key any) any
}

// LoadLogger initializes the logger
func LoadLogger() {
	logger, err := zap.NewDevelopment(
		zap.AddStacktrace(zap.ErrorLevel),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		panic(err)
	}

	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	zap.L().Info("Logger initialized")
}

func Info(ctx LoggerContext, msg ...any) {
	if ctx == nil {
		zap.L().Info(fmt.Sprintf("%v", msg...))
		return
	}

	correlationID, ok := ctx.Value(constants.CORRELATION_ID).(string)
	if !ok {
		zap.L().Info(fmt.Sprintf("%v", msg...))
		return
	}

	zap.L().Info(fmt.Sprintf("[Correlation ID: %s] %v", correlationID, fmt.Sprint(msg...)))
}

func Error(ctx LoggerContext, msg ...any) {
	if ctx == nil {
		zap.L().Error(fmt.Sprintf("%v", msg...))
		return
	}

	correlationID, ok := ctx.Value(constants.CORRELATION_ID).(string)
	if !ok {
		zap.L().Error(fmt.Sprintf("%v", msg...))
		return
	}

	zap.L().Error(fmt.Sprintf("[Correlation ID: %s] %v", correlationID, fmt.Sprint(msg...)))
}

func Debug(ctx LoggerContext, msg ...any) {
	if ctx == nil {
		zap.L().Debug(fmt.Sprintf("%v", msg...))
		return
	}

	correlationID, ok := ctx.Value(constants.CORRELATION_ID).(string)
	if !ok {
		zap.L().Debug(fmt.Sprintf("%v", msg...))
		return
	}

	zap.L().Debug(fmt.Sprintf("[Correlation ID: %s] %v", correlationID, fmt.Sprint(msg...)))
}

func Warn(ctx LoggerContext, msg ...any) {
	if ctx == nil {
		zap.L().Warn(fmt.Sprintf("%v", msg...))
		return
	}

	correlationID, ok := ctx.Value(constants.CORRELATION_ID).(string)
	if !ok {
		zap.L().Warn(fmt.Sprintf("%v", msg...))
		return
	}

	zap.L().Warn(fmt.Sprintf("[Correlation ID: %s] %v", correlationID, fmt.Sprint(msg...)))
}

func Fatal(ctx LoggerContext, msg ...any) {
	if ctx == nil {
		zap.L().Fatal(fmt.Sprintf("%s", msg...))
		return
	}

	correlationID, ok := ctx.Value(constants.CORRELATION_ID).(string)
	if !ok {
		zap.L().Fatal(fmt.Sprintf("%v", msg...))
		return
	}

	zap.L().Fatal(fmt.Sprintf("[Correlation ID: %s] %v", correlationID, fmt.Sprint(msg...)))
}
