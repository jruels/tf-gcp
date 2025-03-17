package main

import (
	"fmt"
	"os"
	"strconv"
)

var hello = []string{"Hello", "Hola",
	"Bon Jour", "Ciao", "こんにちは"}

func main() {
	var index = 1
	if len(os.Args) > 1 {
		index, _ = strconv.Atoi(os.Args[1])
	}
	if index < 1 || index > len(hello) {
		index = 1
	}

	fmt.Println(hello[index-1])
	temp := hello[index-1]
	fmt.Printf("string: \"%s\":\n", temp)
	fmt.Println("index rune char bytes")
	for index, char := range temp {
		fmt.Printf("%-2d %U '%c' %X\n",
			index, char, char,
			[]byte(string(char)))
	}
}
