package main

import (
	"fmt"
)

func main() {

	slice1 := []int{1, 2, 3}   // slice1 is 1,2,3
	slice2 := slice1[1:3]	   // slice2 is 2,3       were the elements or "references" copied

	slice2[0] = 5

	fmt.Println(slice1)
}
