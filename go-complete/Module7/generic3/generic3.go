package main

// go get golang.org/x/exp/constraints

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// add method with generics
func add[T constraints.Ordered](a T, b T) T {
	return a + b
}
func main() {
	fmt.Println(add(4, 2))
}
