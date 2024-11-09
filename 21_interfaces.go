package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}


type rect2 struct {
	width, height float64
}

type circle struct {
	radius float64
}

// meth1
func (r rect2) area() float64 {
	return r.width * r.height
}

func (r rect2) perim() float64 {
	return 2*r.width + 2*r.height
}

// meth2
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}


func measure(g geometry) {  // аргументом в эту функции подойдет только тип имеющий 2 метода определенных выше в интерфейсе geometry
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect2{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}