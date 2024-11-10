// Автор: Евгений Петров
// Почта: p174@mail.ru
// Цель: Обучение языку go
// Урок: https://go.dev/doc/tutorial/handle-errors

package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
