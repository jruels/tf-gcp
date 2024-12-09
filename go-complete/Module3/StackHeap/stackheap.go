package main

import "fmt"

//import "os"

func main() {
	pa := new(int)
	fmt.Printf("Type (pa) = %T\n", pa)
	fmt.Printf("a = %v pa=%v\n", *pa, pa)

	// b := 5
	// fmt.Printf("pa = %v pb=%v", pa, &b)

	pa = nil
}
