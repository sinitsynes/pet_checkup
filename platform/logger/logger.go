package logger

import (
	"log/slog"
	"os"
)

func Setup(stand string) {
	var handler slog.Handler
	if stand != "local" {
		handler = slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo})
	} else {
		handler = slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug})
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
