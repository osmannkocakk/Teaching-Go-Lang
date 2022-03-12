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
	wait(0, 5)
	wait(6, 10)
}
