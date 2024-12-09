package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	
	me := Person{
		Id:   1,
		Name: "Me",
		Age:  64}

	b, err := json.Marshal(me)
	fmt.Println(b, err)
	err2 := ioutil.WriteFile("hello3.bin",
		b, 0644)
	if err2 != nil {
		log.Fatal(err)
	}
}
