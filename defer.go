package main

import "fmt"

func main() {
	defer welcome()
	fmt.Print("My Family,")
	fmt.Print("My Friends,")
	fmt.Print("My Guests ")
}
func welcome() {
	fmt.Println("Welcome!")
}
