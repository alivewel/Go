package sl

import (
	"golang.org/x/exp/slog"
)

// функция принимает ошибку и возвращает
// атрибут с ключом и текстовой ошибкой
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
