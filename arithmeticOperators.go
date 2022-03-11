package main

import "fmt"

func main() {
	x := 27
	y := 5

	name := "Osman "
	surname := "Kocak"

	fmt.Println(name + surname)

	z := x + y
	fmt.Println(z)

	z = x - y
	fmt.Println(z)

	z = x * y
	fmt.Println(z)

	z = x / y // because of z:int the value is: 5 not 5.4
	fmt.Println(z)

	z = x % y
	fmt.Println(z)
}
