package main

import "fmt"

type MyStruct struct {
	data1 int
	data2 int
}

func (s MyStruct) Increment() {
	s.data1++
	s.data2++
}

func main() {

	var c = MyStruct{5, 10}
	c.Increment()
	fmt.Println(c)
}
