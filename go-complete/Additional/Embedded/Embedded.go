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
	var c Person = Person{"Bob", "Wilson",
		Address{"200 Broad St", "Phoenix",
			"AZ"}}

	fmt.Println(c.street, c.city,
		c.state)
}
