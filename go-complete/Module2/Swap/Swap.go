package main

import "fmt"

func main() {
	a := 5
	b := 10
	a = swap1(a, b)
	fmt.Println(a, b)

	swap2(&a, &b)
	fmt.Println(a, b)
}

func swap1(_ int, y int) int {
	return y
}

func swap2(x *int, y *int) {
	*x, *y = *y, *x
}
