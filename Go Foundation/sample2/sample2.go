package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var fib = []int(generate_fibonacci(100000))
	// fmt.Println(fib)
	var fib_from, fib_to int
	fib_from, _ = strconv.Atoi(os.Args[1])
	fib_to, _ = strconv.Atoi(os.Args[2])

	for i, _ := range fib {
		if fib[i] > fib_from {
			if fib[i] < fib_to {
				fmt.Println(fib[i])
			}
		}
	}

}

func generate_fibonacci(fib_max_value int) []int {
	var fib []int

	var a int = 0
	fib = append(fib, a)
	var b int = 1
	fib = append(fib, b)

	for {
		a, b = swap(a, b)
		b = a + b
		if b < fib_max_value {
			fib = append(fib, b)
		} else {
			return fib
		}
	}
}

func swap(x int, y int) (int, int) {
	return y, x
}
