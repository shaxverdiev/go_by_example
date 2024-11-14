// https://gobyexample.com/command-line-arguments

package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[2]

	fmt.Println(argsWithProg)  // 0 аргумент - название функции 
	fmt.Println(argsWithoutProg)  // тут перечисление начнется со первого по индексу элемента
	fmt.Println(arg)  // вернется второй по индексу элемент
}

// $ go build command-line-arguments.go
// $ ./command-line-arguments a b c d
// [./command-line-arguments a b c d]       
// [a b c d]
// c