package main

import (
	"fmt"
)

type Int int

func (value Int) bool() bool {
	return value != 0
}

func int2bool(value int) bool {
	return value != 0
}

func main() {

	fmt.Println(int2bool(5))
	fmt.Println(int2bool(0))

	fmt.Println(Int(5).bool())
}
