package main

import "fmt"

func main() {

	a1 := [...]int{11, 12, 13, 14, 15}
	for a, b := range a1 {
		fmt.Println(a, b)
	}
}
