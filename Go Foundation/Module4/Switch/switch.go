package main

import "fmt"

func main() {

	val := "str"
	switch val {
	case "test":
		fmt.Println("no")
	case "str":
		fmt.Println("good")
	}

	val2 := true
	switch val2 {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	default:
		fmt.Println("impossible!")
	}
}
