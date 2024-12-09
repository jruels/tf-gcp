package main

import (
	"fmt"
)

func main() {

	for {
		var s string
		fmt.Print("Enter text: ")
		fmt.Scanln(&s)
		fmt.Println(s)
		if len(s) == 0 {
			print("done")
			break
		}
	}
}
