// https://gobyexample.com/context

package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello3(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10*time.Second):
		fmt.Fprint(w, "hello\n")
	case <-ctx.Done():  // это сработает если клиент отключился до истечения 10 секунд
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}	
}

func main() {
	http.HandleFunc("/hello", hello3)
	http.ListenAndServe(":8090", nil)
}