package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}


func main() {

	nextInt := intSeq()   // nextInt присвоена функция тк intSeq() возвращает функцию

	fmt.Println(nextInt)    // 0x48f380
	fmt.Println(nextInt())  // 1
    fmt.Println(nextInt())  // 2

	newInt2 := intSeq()    // так выделена новая область в памяти, счетчик внутри функции intSeq обнуляется
	fmt.Println(newInt2)    // 0x48f360
	fmt.Println(newInt2())  // 1
	fmt.Println(nextInt())   // 3 - тут будет все продолжится, тк будет модифицировать другая область памяти
}

