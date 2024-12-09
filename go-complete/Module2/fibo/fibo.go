package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib(start int, end int) []int {
	var a = 0
	var b = 1
	var answer []int

	for b <= end {
		if a >= start {
			answer = append(answer, a)
		}
		a, b = b, a+b
	}
	if a <= end {
		answer = append(answer, a)
	}
	return answer
}

func main() {

	// default values
	var start int = 100
	var end int = 300

	fmt.Printf("%d %d %T\n", start, end, end)
	if len(os.Args) == 3 {
		start, _ = strconv.Atoi(os.Args[1])
		end, _ = strconv.Atoi(os.Args[2])
	}

	// max end value
	if end > 100000 {
		end = 100000
	}

	// print called func
	fmt.Println(fib(start, end))
}
