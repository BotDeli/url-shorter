package logger

import (
	slog "golang.org/x/exp/slog"
	"log"
	"os"
	"url-shorter/internal/config"
)

const (
	errInvalidLogLevel  = "invalid log level %s"
	errInvalidLogFormat = "invalid log format %s"
)

func MustStartLogger(config config.LoggerConfig) *slog.Logger {
	handlerOptions := getHandlerOptions(config.Level)
	handler := getHandler(config.Format, handlerOptions)
	logger := slog.New(handler)
	return logger
}

func getHandlerOptions(level string) *slog.HandlerOptions {
	slogLevel := getSlogLevel(level)
	return &slog.HandlerOptions{Level: slogLevel}
}

func getSlogLevel(level string) (slogLevel slog.Level) {
	switch level {
	case "debug":
		slogLevel = slog.LevelDebug
	case "info":
		slogLevel = slog.LevelInfo
	case "warn":
		slogLevel = slog.LevelWarn
	case "error":
		slogLevel = slog.LevelError
	default:
		log.Fatalf(errInvalidLogLevel, level)
	}
	return
}

func getHandler(format string, handlerOptions *slog.HandlerOptions) (handler slog.Handler) {
	switch format {
	case "text":
		handler = slog.NewTextHandler(os.Stdout, handlerOptions)
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, handlerOptions)
	default:
		log.Fatalf(errInvalidLogFormat, format)
	}
	return
}
