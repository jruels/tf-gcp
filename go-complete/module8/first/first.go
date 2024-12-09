package main

import "fmt"

func funca() {
	fmt.Println(
		"In Goroutine funca")
}

func main() {
	go funca()
}
