package main

import (
	"fmt"
	"time"
)

func main() {
	// timer1 := time.NewTimer(2 * time.Second)

	// <-timer1.C   // Операция <-timer1.C блокирует выполнение программы до тех пор, пока таймер не сработает
	// fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		// ничего не будет выведено. Так как основная функция main() завершится раньше, чем истечёт 1 секунда, горутина не успеет вывести "Timer 1 fired".
		<-timer2.C
		fmt.Println("Timer 1 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}