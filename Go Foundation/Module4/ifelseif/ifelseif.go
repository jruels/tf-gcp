package main

import "fmt"

func main() {
	a := 5
	b := 10
	if a > b {
		fmt.Println("a")
	}
	if b > a {
		fmt.Println("b")
	}
	if b == a {
		fmt.Println("a equal b")
	}

	// Alternative
	if a > b {
		fmt.Println("a")
	} else if b > a {
		fmt.Println("b")
	} else {
		fmt.Println("a equal b")
	}
}
