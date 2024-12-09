package main

import (
	"os"
	"text/template"
)

func main() {
	inline := "Hello, {{.}}"
	data := "World"
	t, _ := template.New("t").Parse(inline)
	f, _ := os.Create("dat2")
	defer f.Close()
	t.Execute(f, data)
}
