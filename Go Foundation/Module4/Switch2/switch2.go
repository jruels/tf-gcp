package main

import (
	"fmt"
)

func main() {

	a := 5
	b := 10

	val := 9
	b--
	switch val {
	case a:
		fmt.Println(a)
	case b:
		fmt.Println(b)
	case a + b:
		fmt.Println(a + b)
	}
}
