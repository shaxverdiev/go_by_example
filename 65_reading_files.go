// https://gobyexample.com/reading-files

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("/tmp/dat")
	check(err)


	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))


	// Seek - позволяет перейти к определенному месту в файле и начать читать его оттуда
	// тут с 6ого байта
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	_, err = f.Seek(4, io.SeekCurrent)
	check(err)

	_, err = f.Seek(-10, io.SeekEnd)
	check(err)

	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	// %d — количество прочитанных байт (n3).
	// %d — позиция курсора, с которой началось чтение (o3).
	// %s — прочитанные байты, преобразованные в строку (string(b3)).

	// устанавливаем указатель на чтение на первый байт (перемотка, как с кассетой)
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	r4 := bufio.NewReader(f)
	// bufio.Reader используется для буферизированного чтения, что делает операции чтения более эффективными, особенно если вы часто читаете небольшие куски данных из большого файла.
	// Внутренний буфер (по умолчанию размером 4096 байт) позволяет загружать данные из файла в память, что уменьшает количество обращений к диску
	b4, err := r4.Peek(5)
	// Метод Peek позволяет заглянуть в следующие 5 байт из файла без перемещения курсора.
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))


	f.Close()
}