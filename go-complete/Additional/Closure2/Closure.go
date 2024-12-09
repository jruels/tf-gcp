package main

import "fmt"

func AFunc() func() {
	a := 5
	return func() {
		a++
		fmt.Println(a)
	}
}

func main() {
	f := AFunc()
	f()
}
