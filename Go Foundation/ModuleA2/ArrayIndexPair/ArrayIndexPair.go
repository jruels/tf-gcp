package main

import "fmt"

func main() {

	example2 := [...]int{10: 4}
	for index, value := range example2 {
		fmt.Println(index, value)
	}

	// example3 := []int{8: 7, 10: 4, 20: 12}
	// for index, value := range example3 {
	// 	fmt.Println(index, value)
	// }

	// example4 := [15]int{9, 10, 10: 4}
	// for index, value := range example4 {
	// 	fmt.Println(index, value)
	// }

	// example5 := [...]string{1: "a", "b", 10: "j"}
	// for index, value := range example5 {
	// 	fmt.Println(index, value)
	// }
}
