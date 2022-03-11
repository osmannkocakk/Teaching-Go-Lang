package main

import "fmt"

func main() {
	a := 5
	b := 4
	c := CalcFunc(a, b)
	fmt.Println(c)
	a, b = swap(a, b)
	fmt.Println(a)
	fmt.Println(b)

}
func CalcFunc(num1, num2 int) int {
	return num1 * num2
}
func swap(x, y int) (int, int) {
	return y, x
}
