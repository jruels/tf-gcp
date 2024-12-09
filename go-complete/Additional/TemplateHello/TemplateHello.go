package main

import (
	"os"
	"text/template"
)

type Data struct {
	A int
	B int
}

func main() {
	inline := "Hello, {{.A}} and {{.B}}"
	data := Data{1, 2}
	t := template.New("")
	t.Parse(inline)
	t.Execute(os.Stdout, data)
}
