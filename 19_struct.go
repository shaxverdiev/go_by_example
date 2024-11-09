package main

import "fmt"

type person struct {
	n string
	a int
}

func newPerson(name string) *person {

	p := person{n: name}
	p.a = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{n: "Alice", a: 30})

	fmt.Println(person{n: "Fred"})

	fmt.Println(&person{n: "Ann", a: 40})

	fmt.Println(newPerson("John"))


	s := person{n: "Sean", a: 50}
	fmt.Println(s.n)
	fmt.Println(s.a)

	sp := &s
	fmt.Println(sp.a)

	sp.a = 51
	fmt.Println(sp.a)

	dog := struct {
		name string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}