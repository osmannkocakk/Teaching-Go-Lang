package main

import "fmt"

type Customer struct {
	name string
	age  int
}

func (ptrC *Customer) addValue(val int) {
	ptrC.age += val
}

func main() {
	x := Customer{"Osman Kocak", 36}
	x.addValue(1)
	fmt.Println(x.age)
}
