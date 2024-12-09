package main

import (
	"fmt"
	"strings"
)

func main() {

	test := "String conversion"
	test2 := []string{"one", "two", "three"}
	fmt.Printf("1 - %d\n", strings.Count(test, "i"))
	fmt.Printf("2 - %t\n", strings.HasPrefix(test, "Str"))
	fmt.Printf("3 - %s\n", strings.Join(test2, " "))
	fmt.Printf("4 - %s\n", strings.ToUpper(test))
}
