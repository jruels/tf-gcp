package main

import (
	"os"
	"text/template"
)

func main() {
	inline := `{{len .}}`
	data := "World"
	t, _ := template.New("").Parse(inline)
	t.Execute(os.Stdout, data)
}
