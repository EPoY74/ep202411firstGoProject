// Автор: Евгений Петров
// Почта: p174@mail.ru
// Цель: Обучение языку go
// Урок: https://go.dev/doc/tutorial/handle-errors

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"example.com/greetings"
)

func main() {

	// Открываю файл, если не получается - выдаю ошибку
	const fileName = "hello.log"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_EXCL|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Файл %v уже существует.\nЕсли файл не существует - произошла другая ошибка", fileName)
	} else {
		timeNow := time.Now()
		formattedString := fmt.Sprintf("Файл создан %v  в  %v \n",
			timeNow.Format(time.DateOnly),
			timeNow.Format(time.TimeOnly))
		if _, err := file.WriteString((formattedString)); err != nil {
			log.Fatalf("Не удалось записать в файл %v", err)
		}
		defer file.Close()
	}

	file1, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)
	if err != nil {
		log.Fatalf("Не удалось открыть файл %v", err)
	}
	defer file1.Close()

	multiWriter := io.MultiWriter(os.Stdout, file1)
	logger := log.New(multiWriter, "Ошибка: ", log.Ldate|log.Ltime|log.Lshortfile)

	message, err := greetings.Hello("")
	if err != nil {
		logger.Println(err)
		os.Exit(1)
		// log.Fatal(err)
	}

	fmt.Println(message)
	// fmt.Println(err)
}
