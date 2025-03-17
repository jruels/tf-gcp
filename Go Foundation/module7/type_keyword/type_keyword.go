package main

import "fmt"

type example1 = int
type example2 int

func main() {
	var vara example1
	var varb example2

	fmt.Printf("%T %T", vara, varb)
}
