package main

import (
	"fmt"
	"slices"
	// "slices"
)

func main() {

	var s[]string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	// uninit: [] true true

	s = make([]string, 3)  // [   ]
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	// Length (длина) — это количество элементов в коллекции.
    // Capacity (вместимость) — это максимальное количество элементов, которое коллекция может содержать, не требуя перераспределения памяти
	s[0] = "a"
	s[1] = "b"
	s[2] = "cc"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	

	// создание среза (динамического массива)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:4]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	
	t := []string{"g", "c", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		// вывод если срезы равны
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}	
	}

	lit := []string{"sss", "332"}
	fmt.Println(lit)

	twoStr := make([][]string, 3)
	fmt.Println("2d: ", twoD) 
	fmt.Println("strslice:", twoStr)
}
