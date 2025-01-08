package app

import "log/slog"

type BadRequestErr struct {
	Err error
}

func (e *BadRequestErr) Error() string {
	slog.Error("Ошибка парсинга запроса",
		slog.Any("error", e.Err.Error()))
	return e.Err.Error()
}

func (e *BadRequestErr) Unwrap() error {
	return e.Err
}
