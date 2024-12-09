package main

import "fmt"

type dog struct{}
type cat struct{}
type rodent struct{}

type herd[T dog | cat] struct {
	animals []T
}

func main() {
	myherd := herd[rodent]{animals: []cat{cat{}, cat{}, cat{}}}
	fmt.Println(myherd)
}
