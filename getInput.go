package main

import "fmt"

func main() {
	var userInput string
	var intInput int
	fmt.Scanln(&userInput)
	fmt.Println(userInput)

	fmt.Scanln(&intInput)
	fmt.Println(intInput * 2)
}
