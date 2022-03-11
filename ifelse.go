package main

import "fmt"

func main() {
	x := 18
	if x >= 18 {
		fmt.Println("Allowed")
	} else {
		fmt.Println("Not Allowed!")
	}

	if y := 10; y > 9 {
		fmt.Println(y)
	}
}
