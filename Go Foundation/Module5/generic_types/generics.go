package main

import "fmt"

type dog struct{}
type cat struct{}

type herd[T any] struct {
	animals []T
}

func main() {
	myherd := herd[dog]{animals: []dog{dog{}, dog{}, dog{}}}
	fmt.Println(myherd)
}
