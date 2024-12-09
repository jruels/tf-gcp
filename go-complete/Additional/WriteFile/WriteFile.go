package main

import (
	"io/ioutil"
	"log"
)

func main() {
	s := "This is super cool!, File!"
	err := ioutil.WriteFile("hello.txt",
		[]byte(s), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
