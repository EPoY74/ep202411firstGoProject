// Автор: Евгений Петров
// Почта: p174@mail.ru
// Цель: Обучение языку go
// Урок: https://go.dev/doc/tutorial/handle-errors

package main

import (
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func main() {

	// Открываю файл, если не получается - выдаю ошибку
	file, err := os.OpenFile("helo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetPrefix("Greetings: ")
	//log.SetFlags(0)
	if err != nil {
		log.Fatalf("Не удалось открыть файл %v", err)
	}
	defer file.Close()

	message, err := greetings.Hello("Milya")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	// fmt.Println(err)
}
