package main

import "fmt"

// обычная функция
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))


	// анонимные функции тоже могут быть рекурсивными
	var fib func(n int) int  // определение переменной с определенным типом (в переменной может содержаться исключительно функции соответствующая той что в типе )

	fib = func(n int) int {  // присваивается функция соответствующая типу
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}