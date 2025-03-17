package main

import "fmt"
import "reflect"

func main() {

	a, b := swap(10, 5.4)
	c, d := swap("dog", "cat")

	fmt.Println("a", a, "b", b)
	fmt.Println("c", c, "d", d)
}

func swap(first interface{},
	second interface{}) (interface{},
	interface{}) {
	ft := reflect.TypeOf(first).Kind()
	st := reflect.TypeOf(second).Kind()

	if ft != st {
		return nil, nil
	}

	return second, first
}
