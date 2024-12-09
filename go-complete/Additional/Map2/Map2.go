package main

import (
	"fmt"
)

func main() {

	map1 := map[string]int{"a": 1, "b": 2}

	value, exist := map1["c"]

	fmt.Println(value)
	if !exist {
		fmt.Println("not okay!")
	}
}
