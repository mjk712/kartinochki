package e

import (
	"fmt"
)

func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err) // %w используем только для ошибок
}

func WrapIfErr(msg string, err error) error {
	if err == nil {
		return nil
	}
	return Wrap(msg, err)
}
