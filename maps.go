package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Osman"] = 36
	m["John"] = 34
	m["Amy"] = 30
	fmt.Println(m["Osman"])
	delete(m, "John")
	fmt.Println(m["John"]) //0
	fmt.Println(m)
}
