package main

import "fmt"

func handler() {
	err := recover()
	if err != nil {
		fmt.Println("***", err.(error))
	}
}
func main() {
	defer handler()
	err := fmt.Errorf("my error")
	panic(err)
}
