package main

import "fmt"

func main() {
	a := make([]int, 4)
	a[1] = 1
	a[2] = 2

	for i, v := range a {
		fmt.Println(i, v)
	}
	for _, v := range a {
		fmt.Println(v)
	}
	for i, _ := range a {
		fmt.Println(i)
	}

	x := "hello"
	for _, c := range x {
		fmt.Println(c) //unicode output
	}
	for _, c := range x {
		fmt.Printf("%c \n", c) //actual characters
	}
}
