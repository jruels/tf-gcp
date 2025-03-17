package main

import "fmt"

func main() {
	var xyz int = 5
	var pxyz *int

	pxyz = &xyz
	fmt.Println(pxyz, *pxyz)

	fmt.Printf("%T %T", xyz, pxyz) // C/C++ format specifiers
}
