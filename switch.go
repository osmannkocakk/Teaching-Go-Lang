package main

import "fmt"

func main() {
	num := 2
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Zero")
	}

	age := 17
	switch {
	case age >= 0 && age < 18:
		fmt.Println("Not Allowed")
	case age >= 18:
		fmt.Println("Allowed")
	default:
		fmt.Println("N/A")
	}
}
