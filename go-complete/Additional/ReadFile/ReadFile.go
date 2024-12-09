package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	fileBytes, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileBytes)
	fileString := string(fileBytes)
	fmt.Println(fileString)
}
