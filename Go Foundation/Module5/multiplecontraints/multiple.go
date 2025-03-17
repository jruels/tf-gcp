package main

import "fmt"

func myfunc[T int | float64](value T) {
	fmt.Println(value)
}

func main() {
	myfunc(42)
}
