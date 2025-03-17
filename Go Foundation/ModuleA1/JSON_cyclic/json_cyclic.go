package main

import _ "encoding/json"
import _ "fmt"

type node struct {
	Previous *node
	Next     *node
	Location int
}

func main() {

	node1 := node{nil, nil, 1}
	//	node1.Previous = &node1
	b, err := json.Marshal(node1)

	if err == nil {
		fmt.Printf("%s", b)
	} else {
		fmt.Println(err)
	}
}
