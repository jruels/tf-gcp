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

func main() {

	var c Person

	fmt.Println(c.get())
}
