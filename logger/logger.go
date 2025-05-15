package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var Log *zap.Logger

// InitializeLogger initializes the Zap logger with the specified level.
// If no level is specified, it defaults to "info".  Uses environment variable for level.
func InitializeLogger() {
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "info" // Default log level
	}

	lvl, err := zapcore.ParseLevel(level)
	if err != nil {
		log.Fatalf("Invalid log level: %s", level)
	}

	config := zap.NewProductionConfig() // Or NewDevelopmentConfig for development
	config.Level = zap.NewAtomicLevelAt(lvl)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Optional: Customize time format

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	Log = logger
	Log.Info("Logger initialized", zap.String("level", level))
}

// CloseLogger syncs the logger before exiting.
func CloseLogger() {
	if Log != nil {
		Log.Sync() // Flushes buffer, if any
	}
}
