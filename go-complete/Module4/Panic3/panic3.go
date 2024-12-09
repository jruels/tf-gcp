package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		err := recover() // any type
		var message = err.(error).Error()

		if message == "runtime error: index out of range" {
			fmt.Fprintf(os.Stderr, "Exception handled\n")
		} else if message != "" {
			fmt.Fprintf(os.Stderr, "Unplanned Exception: %v\n", message)
			panic(err)
		} else {
			fmt.Fprintf(os.Stderr, "Exception unknown\n")
			panic(err)
		}
	}()

	stuff := []int{1, 2, 3, 4}
	stuff[6] = 12
}
