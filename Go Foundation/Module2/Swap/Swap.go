package main

import "fmt"

func main() {
	a := 5
	b := 10
	a, b = swap1(a, b)
	fmt.Println(a, b)

	swap2(&a, &b)
	fmt.Println(a, b)
}

func swap1(x int, y int) (int, int) {
	return y, x
}

func swap2(x *int, y *int) {
	*x, *y = *y, *x
}
