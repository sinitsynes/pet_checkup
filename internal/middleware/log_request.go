package middleware

import (
	"log/slog"
	"net/http"
	"os"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // Default to 200 if WriteHeader is not called
		}

		next.ServeHTTP(wrapped, r)
		handler := slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug})
		logger := slog.New(handler)

		logger.Info("Request",
			"method", r.Method,
			"url", r.URL.Path,
			"status", wrapped.statusCode)
	})
}
