package main

import (
	"fmt"
)

func main() {
	const pi = 3.14
	var x int = 1
	var pi32 float32 = 3.14159265           //Output 3.1415927
	var pi64 float64 = 3.141592653589793238 //Output 3.141592653589793
	var name string = "Osman Kocak"
	var isActive bool = true

	fmt.Println(pi)
	fmt.Println(x)
	fmt.Println(pi32)
	fmt.Println(pi64)
	fmt.Println(name)
	fmt.Println(isActive)
}
