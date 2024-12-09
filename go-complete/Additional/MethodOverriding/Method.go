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

func (a Address) get() string {
	return fmt.Sprintf("%s\n%s\n%s",
		a.street, a.city, a.state)
}

func (a Person) get() (string, string) {
	return a.first + " " + a.last + "\n",
		a.Address.get()
}

func main() {

	var c = Person{"Bob", "Wilson",
		Address{"200 Broad St", "Phoenix",
			"AZ"}}

	fmt.Println(c.get())
}
