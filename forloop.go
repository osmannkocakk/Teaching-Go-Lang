package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	num := 1
	for num <= 10 {
		fmt.Println(num)
		num++
	}
}
