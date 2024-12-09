package main

import (
	"fmt"
)

func main() {

	a := 5
	b := 10

	val := 9
	b--
	c := 9

	switch val {
	case a:
		fmt.Println(a)
	case b:
		fmt.Println("B", b)
		fallthrough
	case c:
		fmt.Println("C", c)
	case a + b:
		fmt.Println(a + b)
	}
}
