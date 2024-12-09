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

	var c = MyStruct{data2: 10, data1: 5}
	fmt.Println(c.Add())
	fmt.Println(c.Multiply())
}
