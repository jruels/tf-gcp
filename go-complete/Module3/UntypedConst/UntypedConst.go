package main

import "fmt"

func main() {
	const a int = 5 // type constant
	const b = 5     // untyped const

	fmt.Printf("%T %T\n", a, b)

	//	var c float32 = a
	var d float32 = b
	fmt.Print(d)
}
