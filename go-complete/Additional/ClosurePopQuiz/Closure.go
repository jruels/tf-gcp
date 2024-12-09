package main

import "fmt"

func squares() func() int {
	var x int = 4  // outer function
	var z int = 5
	return func() int {
		x++   // capture the variable
		z++   //  captured also
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
