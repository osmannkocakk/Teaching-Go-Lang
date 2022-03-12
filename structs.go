package main

import (
	"fmt"
)

type Customer struct {
	name    string
	surname string
	age     int
}

func main() {
	customer := Customer{name: "Osman", surname: "Kocak", age: 36}
	fmt.Print(customer.name + " " + customer.surname)
	if customer.age < 18 {
		fmt.Println(" is junior.")
	} else if customer.age >= 18 && customer.age <= 55 {
		fmt.Println(" is young.")
	} else {
		fmt.Println(" is senior")
	}

}
