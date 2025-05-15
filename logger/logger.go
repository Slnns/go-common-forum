package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var Log *zap.Logger

func InitializeLogger(level string) {
	lvl, err := zapcore.ParseLevel(level)
	if err != nil {
		log.Fatalf("Invalid log level: %s", level)
	}

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(lvl)

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	Log = logger
	Log.Info("Logger initialized", zap.String("level", level))
}

func CloseLogger() {
	if Log != nil {
		Log.Sync()
	}
}
