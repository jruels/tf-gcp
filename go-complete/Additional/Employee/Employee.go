package main

import "fmt"

type employee struct {
	firstName string
	lastName  string
	age       int
	weight    int
}

func main() {
	fred := employee{age: 45, firstName: "Fred",
		lastName: "Wilson", weight: 156}
	fmt.Println(fred)
}
