package main

import (
	"os"
)

func main() {

	filename := "hello.txt"
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	text := "42"
	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}
