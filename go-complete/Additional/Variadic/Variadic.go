package main

import "fmt"

func main() {

	t1 := total(12, 5, 9, 3, 10)
	fmt.Println(t1)
}

func total(a int, values ...int) int {
	var result int
	for _, item := range values {
		result += item
	}
	return result
}
