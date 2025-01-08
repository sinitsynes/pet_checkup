package app

import (
	"log/slog"
	"os"
)

func LoggerSetup() *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, opts))
}
