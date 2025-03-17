package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var i int = 1

	fmt.Printf("Size of i is: %d",
		unsafe.Sizeof(i))
}
