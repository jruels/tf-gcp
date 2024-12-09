package main

import "fmt"

func main() {

	type BitFlag int
	const (
		Bold BitFlag = 1 << iota
		Underline
		Italics
	)

	result := Underline | Italics

	fmt.Println("Flags", Bold, Underline, Italics)
	fmt.Println("Result", result)
}
