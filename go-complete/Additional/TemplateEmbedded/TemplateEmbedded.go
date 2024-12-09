package main

import (
	"os"
	"text/template"
)

type Culture struct {
	Minor string
	Major string
}
type Data struct {
	Greeting string
	Culture
}

func main() {
	g := Data{"Hello", Culture{Minor: "en", Major: "US"}}
	inline := `"{{.Greeting}}" in this culture: {{.Minor}}-{{.Major}}`

	t := template.New("")
	t.Parse(inline)

	err := t.Execute(os.Stdout, g)

	if err != nil {
		// do something
	}
}
