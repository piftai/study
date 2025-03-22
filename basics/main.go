package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func(ch chan int) {
		ch <- 1 // за сколько времени выполнится программа если оставим эту строку здесь, а за сколько выполнится если перенесем на 14 строку?
		fmt.Println("hello from out goroutine")
		time.Sleep(5 * time.Second)
	}(ch)
	fmt.Println(<-ch)
}
