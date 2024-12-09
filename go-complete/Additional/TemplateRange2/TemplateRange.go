package main

import (
	"os"
	"text/template"
)

type Array struct {
	Data  []string
	Count uint
}

func main() {
	info := Array{[]string{"1", "2", "5", "7"}, 5}
	t := template.New("")

	t.Parse(`{{range .Data}} {{.}} {{end}} 
	 Total {{.Count}}`)

	t.Execute(os.Stdout, info)
}
