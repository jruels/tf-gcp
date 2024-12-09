package main

import (
	"fmt"
)

type fptr func(int, int) int

func main() {
	mathFunctions := []fptr{a, d, m, s}
	mathfunctionNames := []string{"Addition", "Division", "Multiplication", "Subtraction"}
	//Test functions
	//for _, v := range mathFunctions {
	//    k := v(10,4)
	//    fmt.Println(k)
	//}

	mymap := make(map[string]fptr, len(mathfunctionNames))
	//build the map with functions
	for i, name := range mathfunctionNames {
		mymap[name] = mathFunctions[i]
	}
	for k, v := range mymap {
		fmt.Printf("%s : %d\n", k, v(10, 5))
	}
}
func a(i int, j int) int {
	fmt.Printf("a(%d, %d): ", i, j)
	return i + j
}
func d(i int, j int) int {
	fmt.Printf("d(%d, %d): ", i, j)
	return i / j
}
func m(i int, j int) int {
	fmt.Printf("m(%d, %d): ", i, j)
	return i * j
}
func s(i int, j int) int {
	fmt.Printf("s(%d, %d): ", i, j)
	return i - j
}
