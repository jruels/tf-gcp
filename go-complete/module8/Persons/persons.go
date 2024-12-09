package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fredwilson := person{"Fred", 45}
	sallyjohnson := person{age: 45, name: "Sally"}
	fmt.Println(fredwilson)
	sallyjohnson.age += 15
	fmt.Println(sallyjohnson.name,
		sallyjohnson.age)
}
