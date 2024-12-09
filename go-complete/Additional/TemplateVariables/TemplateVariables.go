package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("")
	t.Parse(`{{ $F := .FirstName }}
	  {{ $L := .LastName }}
	  Normal: {{$F}} {{$L}}
             Reverse: {{$L}} {{$F}}`)
	t.Execute(os.Stdout, struct {
		FirstName string
		LastName  string
	}{
		"Gigi",
		"Sayfan",
	})
}
