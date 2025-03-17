package main

import (
	"fmt"
)

func main() {

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 3) // empty slice
	slice2 = slice1          // copied into the slice

	slice2[0] = 5 // change is isolated
	slice1[2] = 12

	fmt.Println(slice1)
	fmt.Println(slice2)
}
