package main

import (
	"fmt"
	"sync"
)

func funca() {
	defer wg.Done()
	fmt.Println(
		"In Goroutine funca")
}

func funcb() {
	defer wg.Done()
	fmt.Println(
		"In Goroutine funcb")
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go funca()
	go funcb()
	wg.Wait()
}
