package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerInstance *zap.Logger
	once           sync.Once
)

// GetLogger initializes and returns a custom logger with colored output
func GetLogger() *zap.Logger {
	once.Do(func() {
		config := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder, // Enables colored levels
			EncodeTime:     zapcore.ISO8601TimeEncoder,       // Human-readable time format
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(config), // Console encoder with color
			zapcore.AddSync(os.Stdout),        // Output to stdout
			zapcore.DebugLevel,                // Log level (adjust as needed)
		)

		loggerInstance = zap.New(core, zap.AddCaller())
	})
	return loggerInstance
}

// Sync flushes any buffered log entries
func Sync() {
	if loggerInstance != nil {
		_ = loggerInstance.Sync()
	}
}
