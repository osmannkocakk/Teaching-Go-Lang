package main

import "fmt"

func SumEven(from, to int, ch chan int) {
	sum := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	ch <- sum
}
func SumSquare(from, to int, ch chan int) {
	sum := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			sum += i * i
		}
	}
	ch <- sum
}

func main() {
	chEven := make(chan int)
	chSquare := make(chan int)

	go SumEven(0, 100, chEven)
	go SumSquare(0, 100, chSquare)

	//select only one that comes first
	select {
	case x := <-chEven:
		fmt.Println(x)
	case y := <-chSquare:
		fmt.Println(y)
	}
}
