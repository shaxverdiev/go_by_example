package main

import (
	"fmt"
	"time"
)

func foo(from string) {
	for i := 0; i < 59; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	foo("direct")
	
	go foo("goroutine")
	time.Sleep(time.Second)

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Println("done")
}