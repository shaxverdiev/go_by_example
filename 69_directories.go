// https://gobyexample.com/directories

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := os.Mkdir("/tmp/subdir", 0755)
	check(err)

	// defer os.RemoveAll("/tmp/subdir")


	createEmptyFile := func(name string) {
		// d := []byte("")
		d := make([]byte, 50000)
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("/tmp/subdir/file1")

	err = os.MkdirAll("/tmp/subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}


