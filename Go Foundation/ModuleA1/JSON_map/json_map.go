package main

import "encoding/json"
import "fmt"

func main() {
	data := map[string]int{"a": 1, "b": 2, "c": 3}
	//data := map[int]int{42: 1, 43: 2, 44: 3}

	b, err := json.Marshal(data)

	if err == nil {
		fmt.Printf("%s", b)
	} else {
		fmt.Println(err)
	}

}
