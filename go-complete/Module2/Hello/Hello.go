package main // executable binary

import (
	"fmt"     // println
	"os"      // command-line args
	"strconv" // convert from string to integer
)

var hello = []string{"Hello", "Hola",
	"Bon Jour", "Ciao", "こんにちは", // UNICODE UTF-8
}

func main( /* command line args */ ) {
	//_ := 5

	//var index int = 1
	index := 1 // short declaration
	if len(os.Args) > 1 {
		index, _ = strconv.Atoi(os.Args[1])
	}

	if index < 1 || index > len(hello) {
		index = 1
	}

	fmt.Println(hello[index-1])
}
