package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

func (b base) greeting() string {
	return "hello"
}

type container struct {
	base         // тут происходит встраивание типа
	str string
}

func (c container) cont_method() string {
	return "cont method"
}

func main() {
	
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("aslo num:", co.base.num)

	fmt.Println("describe:", co.describe())


	type describer interface {
		describe() string
		// greeting() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())  // тип реализует метод describe (это сработает даже если происходит встраивание типов)
	// fmt.Println("greeting:", d.greeting())
}

