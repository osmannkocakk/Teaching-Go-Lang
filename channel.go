package main

import (
	"fmt"
	"time"
)

func wait(from, to int, ch chan bool) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
	}
	ch <- false
}

func main() {
	ch := make(chan bool)

	go wait(0, 5, ch)
	go wait(6, 10, ch)

	<-ch //waits for channel before finish its job
}
