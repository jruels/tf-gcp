package main

import (
	"fmt"
	"strconv"
)

func main() {

	x, _ := strconv.ParseFloat("120.21", 64)
	fmt.Printf("%8T\t%5v\n", x, x)
	y, _ := strconv.ParseInt("42", 10, 0)
	fmt.Printf("%8T\t%5v\n", y, y)
	z, _ := strconv.Atoi("42")
	fmt.Printf("%8T\t%5v\n", z, z)
}
