package main

import "fmt"

func equal3[T comparable](a T, b T, c T) bool {
	return (a == b) == (b == c)
}

func main() {
	fmt.Println(equal3(4, 2, 5))
}
