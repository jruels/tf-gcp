package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func main() {
	inline := "Hello, {{.}}"
	data := "World"
	t, err := template.New("t").Parse(inline)
	var buffer bytes.Buffer
	err = t.Execute(&buffer, data)
	fmt.Print(buffer.String())
	if err == nil {
	}
}
