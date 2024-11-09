package main

import "fmt"

// канал в который можно отправлять строки
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// канал из которого можно только читать строки
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings1 := make(chan string, 1)
	pongs1 := make(chan string, 1)
	ping(pings1, "passed message")
	pong(pings1, pongs1)
	fmt.Println(<-pongs1)
}

