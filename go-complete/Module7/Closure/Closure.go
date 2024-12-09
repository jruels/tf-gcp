package main

import "fmt"

func main() {

	a := 5
	b := 10

	x := func() int {
		return a * b
	}

	a++

	fmt.Println(x())

	myfunc:=func() int {
		return 42
	}

	fmt.Println(myfunc())
	fmt.Println(myfunc())
}
