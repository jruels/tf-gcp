package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		err := recover() // interface {}
		var message = err.(error).Error()

		if strings.HasPrefix(message, "runtime error: index out of range") {
			fmt.Fprintf(os.Stderr, "Exception handled\n")
		} else if message != "" {
			fmt.Fprintf(os.Stderr, "Unplanned Exception***: %v\n", message)
			panic(err)
		} else {
			fmt.Fprintf(os.Stderr, "Exception unknown\n")
			panic(err)
		}
	}()

	stuff := []int{1, 2, 3, 4}
	stuff[6] = 12
}
