package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}

	//slices
	var a []int = x[1:3]
	fmt.Println(a) //[2 3]
	var b []int = x[:3]
	fmt.Println(b) //[1 2 3]
	var c []int = x[1:]
	fmt.Println(c) //[2 3 4 5]
	c[0] = 6
	fmt.Println(x) //[1 6 3 4 5]

	//Slice with make
	y := make([]int, 3)
	fmt.Println(y)      //[0 0 0]
	y = append(y, 4, 5) //can add more than 1 with comma
	fmt.Println(y)      //[0 0 0 4 5]
}
