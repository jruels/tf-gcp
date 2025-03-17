package main

import (
	"fmt"
)

func main() {

	{
		defer fmt.Println("test")
	}
	fmt.Println("test2")
}
