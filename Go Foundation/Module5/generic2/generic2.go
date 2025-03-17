package main

import "fmt"

// add method with generics
func add[T any](a T, b T) T {
	return a + b
}
func main() {
	fmt.Println(add(4, 2))
}
