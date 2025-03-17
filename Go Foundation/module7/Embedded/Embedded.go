package main

import "fmt"

type Address struct {
	street string
	city   string
	state  string
}
type Person struct {
	first string
	last  string
	Address
}

// Person.city

func main() {
	var bob = Person{"Bob", "Wilson",
		Address{"200 Broad St", "Phoenix",
			"AZ"}}

	fmt.Println(bob.first, bob.last,
		bob.street, bob.city,
		bob.state)
}
