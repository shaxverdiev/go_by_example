/*
Этот код демонстрирует использование итераторов и генераторов в Go с помощью пакетов iter и slices. 
Это позволяет создавать итерируемые последовательности данных и работать с ними в функциональном стиле. 
Особенность кода в том, что с помощью дженериков (T any) и интерфейсов создаются универсальные структуры и функции, 
которые могут работать с любыми типами данных.

Итератор здесь используется для абстракции перебора элементов списка.
Вместо того чтобы напрямую использовать цикл для перебора, вы создаёте функцию, 
которая скрывает логику этого перебора и предоставляет интерфейс для последовательного получения элементов.
*/



package main


import (
	"fmt"
	"iter"
	"slices"
)

type List1[T any] struct {
	head, tail *element1[T]
}

type element1[T any] struct {
	next *element1[T]
	val T
}

func (lst *List1[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element1[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element1[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List1[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}





func main() {
	lst := List1[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}