package main

import (
	"fmt"
	"time"
)

func wait(from, to int) {
	for i := from; i <= to; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}
func main() {
	go wait(0, 5)
	go wait(6, 10)
	time.Sleep(500 * time.Millisecond) //because of waiting to go routines
}
