package main

import (
	"fmt"
	// "time"
)

func main() {
	// i := 2
	// fmt.Print("Write ", i, " as ")
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// }

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it's the weekend", time.Now().Weekday())
	// default:
	// 	fmt.Println("It's a weekday", time.Now().Weekday())
	// }

	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("It's before noon")
	// default:
	// 	fmt.Println("It's after noon")
	// }

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	whatAmI2(true)
	whatAmI2(1)
	whatAmI2("hey")
}

func whatAmI2(i any) {
	switch t := i.(type) {
	case bool:
		println("bool")
	case int:
		println("int")
	default:
		fmt.Printf("%T\n", t)
	}
}
