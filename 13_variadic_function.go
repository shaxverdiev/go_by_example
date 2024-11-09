package main

import "fmt"

func sum(nums ...int) {  // фариадическая функция принимает слайс элементов
	fmt.Print(nums, " ")
	total := 0

	for i, num := range nums {
		fmt.Println(i, num)
		total += num
	}
	fmt.Println("two", total)	
}


func main() {
	sum(11, 22)
	sum(111, 222, 333)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}