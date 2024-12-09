package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	me := Person{}

	fileBytes, err4 := ioutil.ReadFile("hello3.bin")
	if err4 != nil {

	}

	json.Unmarshal(fileBytes, &me)
	fmt.Println(me)
}
