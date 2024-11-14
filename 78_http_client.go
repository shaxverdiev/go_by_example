// https://gobyexample.com/http-client

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	// Мы создаём сканер для чтения данных из resp.Body.
	// bufio.NewScanner позволяет читать данные построчно.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}


// $ go run http-clients.go
// Response status: 200 OK
// <!DOCTYPE html>
// <html>
//   <head>
//     <meta charset="utf-8">
//     <title>Go by Example</title>