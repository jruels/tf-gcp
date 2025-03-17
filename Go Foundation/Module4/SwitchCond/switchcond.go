package main

import "fmt"

func main() {
	const (
		Critical = iota
		Error
		Warning
	)
	val := Critical

	switch val {
	case Critical:
		fmt.Println("Critical")
		fallthrough
	case Error:
		fmt.Println("Error")
		fallthrough
	case Warning:
		fmt.Println("Warning")
	}
}
