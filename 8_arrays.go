package main

import "fmt"

func main() {
	// var a [5]int
	// fmt.Println("emp:", a)

	// a[4] = 100
	// fmt.Println("set:", a)
	// fmt.Println("get:", a[4])

	// fmt.Println("len:", len(a))

	// b := [5]int{1,2,3,4,5}
	// fmt.Println("dcl:", b)

	// b = [...]int{1,2,3,4,5}
	// fmt.Println("dcl:", b)

	// b = [...]int{100, 3: 400, 500}
	// fmt.Println("idx:", b)

	var twoD [2][3]int
	// или	
	// twoD = [2][3]int{
    // {0, 0, 0},  // Первая строка
    // {0, 0, 0},  // Вторая строка
	// }
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			println(i , j)
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	// вывод
	// 0 0
	// 0 1
	// 0 2
	// 1 0
	// 1 1
	// 1 2
	// 2d:  [[0 1 2] [1 2 3]]

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}