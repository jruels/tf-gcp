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

func (Address) addressfunc(){

}

func (Person) addressfunc(){  // function overriding

}

func main() {
	var bob = Person{"Bob", "Wilson",
		Address{"200 Broad St", "Phoenix",
			"AZ"}}

	fmt.Println(bob.first, bob.last,
		bob.street, bob.city,
		bob.state)

	bob.addressfunc()  //  which addressFunc
}
