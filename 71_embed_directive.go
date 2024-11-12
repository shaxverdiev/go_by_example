// https://gobyexample.com/

package main

import (
	"embed"
)
// позволяет встраивать файлы и директории непосредственно в бинарный файл программы. 
// Это полезно, когда вы хотите включить ресурсы (например, HTML, изображения, конфигурации и т.д.) прямо в ваш Go-код, 
// чтобы не зависеть от внешних файлов в процессе выполнения программы


//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

    print(fileString)
    print(string(fileByte))

    content1, _ := folder.ReadFile("folder/file1.hash")
    print(string(content1))

    content2, _ := folder.ReadFile("folder/file2.hash")
    print(string(content2))
}