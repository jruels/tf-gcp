package main

import "fmt"

func main() {

	fmt.Println(funcA())
	fmt.Println(funcB())
	fmt.Println(funcC())
	fmt.Println(funcD())
	fmt.Println(funcE())
}

func funcA() int {
	return 5
}

func funcB() (int, int) {
	return 5, 10
}

func funcC() (a int, b int) {
	a = 12
	b = 18
	return a, b
}

func funcD() (a int, b int) {
	a = 20
	b = 30
	return
}

func funcE() (a int, b int) {
	a = 12
	b = 18
	return b, a
}
