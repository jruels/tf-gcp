package main

import "fmt"

type MyStruct struct {
	data1 int
	data2 int
}

func (s MyStruct) Add() int {
	return s.data1 + s.data2
}

func (s MyStruct) Multiply() int {
	return s.data1 * s.data2
}

func main() {

	var c = MyStruct{5, 10}
	fmt.Println(c.Add())
	fmt.Println(c.Multiply())
}
