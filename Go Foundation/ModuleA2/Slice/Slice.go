package main

import "fmt"

var a1 = make([]int, 5, 10)
var a2 = []int{1, 2, 3, 4, 5}
var a3 = [][]int{{1, 2, 5}, {3, 8, 4}}

func main() {
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}
