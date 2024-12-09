package main

import (
	"bytes"
	"fmt"
)

func main() {

	strings := []string{"one", "two", "three", "four"}
	var buffer bytes.Buffer

	for _, valuestring := range strings {
		buffer.WriteString(valuestring)
	}
	fmt.Println(buffer.String())
}
