package main

import (
	"encoding/gob"
	"os"
)

type Person struct {
	Name string
	Age  int32
}

func main() {
	filename := "buffer.gob"
	bob := Person{"Bob Johnson", 35}

	file, err := os.Create(filename)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(bob)
	}
	file.Close()
}
