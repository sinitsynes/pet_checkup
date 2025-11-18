package middleware

import (
	"log/slog"
	"net/http"
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

		slog.Info("Request",
			"method", r.Method,
			"url", r.URL.Path,
			"status", wrapped.statusCode)
	})
}
