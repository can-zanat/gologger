package logger

import (
	"log"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level string
}

func DefaultConfig() Config {
	config := Config{}
	config.Level = "info"
	return config
}

func getLogLevel(level string) zapcore.Level {
	switch levelFromConfig := strings.TrimSpace(level); {
	case strings.EqualFold(levelFromConfig, "debug"):
		return zapcore.DebugLevel
	case strings.EqualFold(levelFromConfig, "error"):
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func NewWith(config Config) *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(getLogLevel(config.Level))
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	return logger
}

func New() *zap.Logger {
	return NewWith(DefaultConfig())
}

func NewWithLogLevel(level string) *zap.Logger {
	if level != "" {
		return NewWith(Config{Level: level})
	}
	return New()
}
