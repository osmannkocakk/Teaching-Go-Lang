package main

import "fmt"

func sumAll(nums ...int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
}
func main() {
	sumAll(1)
	sumAll(1, 2, 3)
	sumAll(4, 5)
	sumAll(1, 2, 3, 4, 5, 6)
	s := []int{1, 2, 3, 4}
	sumAll(s...)
}
