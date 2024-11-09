package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0  // тут происходит разыменовываение "открытие" после чего в него кладется значение 0, это значит что все значения аргументов будут изменены на 0
}


func main() {
	i := 1333
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)  // что бы не значит i после прохождения через функцию zeroptr они превратится в 0

	fmt.Println("pointer:", &i) // выведется адрес перменной i
}


