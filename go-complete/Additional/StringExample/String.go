package main

import (
	"fmt"
	"strings"
)

var hello = "こんにちは "

func main() {
	indexb := strings.Index(hello, "ん")
	indexe := strings.LastIndex(hello, " ")
	fmt.Println(hello[indexb:indexe])
}
