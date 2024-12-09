package main

import (
	"fmt"
	"os"
)

type fptr func(int) int // fptr is function ptr type

func main() {

	if len(os.Args) == 1 {
		funcExecute(double)
	} else {
		funcExecute(triple)
	}

}

func double(a int) int {
	return a * a
}

func triple(a int) int {
	return a * a * a
}

func test(a, b int) int {
	return a * a * a
}

func funcExecute(f fptr) {
	fmt.Println(f(5))
}
