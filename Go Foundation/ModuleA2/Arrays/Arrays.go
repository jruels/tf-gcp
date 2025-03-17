package main

import "fmt"

var a1 [5]int
var a2 = [5]int{1, 2, 3, 4, 5}
var a3 = [...]int{1, 2, 3, 4, 5}
var a4 = [2][3]int{{1, 2, 5}, {3, 8, 4}}

func main() {
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	//fmt.Println(a4[1][2])
	for i := 0; i < len(a2); i++ {
		fmt.Println(a2[i])
	}
}
