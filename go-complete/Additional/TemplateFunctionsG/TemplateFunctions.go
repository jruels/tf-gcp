package main

import (
	"os"
	"text/template"
)

type Data struct {
	First string
	Last  string
}

func (d Data) Hello() string {
	return "Hello " + d.First + " " + d.Last
}

func main() {
	t := template.New("")
	t.Parse(`{{.Hello}}`)
	t.Execute(os.Stdout, Data{"Sally", "Wilson"})
}
