package main

import "fmt"
import "encoding/json"

type data struct {
	First  string
	Second string
}

func main() { // key:T
	dataJson := `{"first": "data1","second": "data2"}` // string:string
	var mydata data
	err := json.Unmarshal([]byte(dataJson), &mydata)
	if err != nil {
		fmt.Print(err)
		return
	}
}
