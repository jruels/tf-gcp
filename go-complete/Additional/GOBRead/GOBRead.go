package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int32
}

func main() {
	filename := "buffer.gob"
	var bob = new(Person)
	file, err := os.Open(filename)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(bob)
	}
	file.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bob.Name, "\t", bob.Age)
	}

}
