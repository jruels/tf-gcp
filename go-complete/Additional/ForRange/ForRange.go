package main

import "fmt"

func main() {

	//    make(type, size, initial capacity)
	a1 := make([]int, 5, 10)

	//a1 := []int{11, 12, 13, 14, 15}

	fmt.Println(cap(a1))
	for a, b := range a1 {
		fmt.Println(a, b)
	}
}
