// Автор: Евгений Петров
// Почта: p174@mail.ru
// Цель: Обучение языку go
// Урок: https://go.dev/doc/tutorial/handle-errors

package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("Greetings: ")
	//log.SetFlags(0)

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	// fmt.Println(err)
}
