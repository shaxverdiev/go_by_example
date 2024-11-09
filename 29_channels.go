package main

import "fmt"

func main() {
	messages := make(chan string)

	// отправили в канал
	go func() { messages <- "ping"}()

	// получили из канала
	msg := <-messages
	fmt.Println(msg)
}