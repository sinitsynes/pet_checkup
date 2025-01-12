package utils

import (
	"errors"
	"log/slog"
	"net/http"
)

// https://boldlygo.tech/posts/2024-01-08-error-handling/
type statusError struct {
	error
	status int
}

func (e statusError) Unwrap() error   { return e.error }
func (e statusError) HTTPStatus() int { return e.status }

func WithHTTPStatus(err error, status int) error {
	return statusError{
		error:  err,
		status: status,
	}
}

func HTTPStatus(err error) int {
	if err == nil {
		return 0
	}
	var statusErr interface {
		error
		HTTPStatus() int
	}
	if errors.As(err, &statusErr) {
		return statusErr.HTTPStatus()
	}
	return http.StatusInternalServerError
}

type RequestValidationError struct {
	Err error
}

func (e RequestValidationError) Error() string {
	slog.Error("Ошибка парсинга запроса",
		slog.Any("error", e.Err.Error()))
	return e.Err.Error()
}

func (e RequestValidationError) Unwrap() error {
	return e.Err
}

func NewRequestValidationError() RequestValidationError {
	return RequestValidationError{}
}
