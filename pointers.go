package main

import "fmt"

func main() {
	x := 5
	p := &x
	*p = 10

	fmt.Println(x)
	fmt.Println(*p)
	fmt.Println(p)
}
