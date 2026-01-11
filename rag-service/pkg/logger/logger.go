package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init(level string, mode string) error {
	var config zap.Config
	if mode == "debug" {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionConfig()
	}

	// Set log level
	l, err := zapcore.ParseLevel(level)
	if err != nil {
		return err
	}
	config.Level = zap.NewAtomicLevelAt(l)

	// Build logger
	logger, err := config.Build()
	if err != nil {
		return err
	}

	Log = logger
	return nil
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}

// Helper functions for global access
func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}
