package main

import (
	"os"
	"text/template"
)

type Order struct {
	First   string
	Last    string
	Count   int
	Product string
}

	func main() {
		o1 := Order{"Bob", "Jones", 5, "Phone charger"}
		t, _ := template.ParseFiles("Action.txt")
		t.Execute(os.Stdout, o1)
	}
