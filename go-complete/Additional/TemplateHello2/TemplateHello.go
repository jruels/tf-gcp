package main

import (
	"os"
	"text/template"
)

type Greeting struct {
	Greeting string
	Culture  string
}

func main() {
	g := Greeting{"Hello", "enUS"}
	inline := `"{{.Greeting}}" in this culture: {{.Culture}}`
	t := template.New("")
	t.Parse(inline)
	err := t.Execute(os.Stdout, g)
	if err != nil {
		// do something
	}
}
