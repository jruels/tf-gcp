package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var start int
	var end int
	if len(os.Args) > 2 {
		start, _ = strconv.Atoi(os.Args[1])
		end, _ = strconv.Atoi(os.Args[2])
	}
	if end > 100000 {
		end = 100000
	}
	fmt.Println(fib(start, end))
}

func fib(start int, end int) []int {
	var a = 0
	var b = 1
	var temp int
	var result []int

	for b <= end {
		if a >= start {
			result = append(result, a)
		}
		temp = a + b
		a = b
		b = temp
	}
	if a <= end {
		result = append(result, a)
	}

	return result
}
