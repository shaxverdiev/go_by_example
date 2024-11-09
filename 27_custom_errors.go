package main

import (
    "errors"
    "fmt"
)

type argError struct {
    arg     int
    message string
}

// если у типа реализован метод Error (по аналогии с String()) то считается что он тип реализует интерфейс error
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f1(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "can't work with it!!!!!"}
    }
    return arg + 3, nil
}

func main() {
    _, err := f1(42)
    var ae *argError
    if errors.As(err, &ae) {
        fmt.Println(ae.arg)
        fmt.Println(ae.message)
    } else {
        fmt.Println("err doesn't match argError")
    }
}