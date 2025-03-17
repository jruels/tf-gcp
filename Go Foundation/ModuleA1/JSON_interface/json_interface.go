package main

import "encoding/json"
import "fmt"

func main() {

	json_binary := []byte(`{"First":"Sam", "Last":"Wilson", "Salary":12000}`)

	var json_interface any
	err := json.Unmarshal(json_binary, &json_interface)

	if err != nil {
		fmt.Println(err)
		return
	}

	json_map := json_interface.(map[string]any)

	for key, value := range json_map {
		switch value.(type) {
		case string:
			fmt.Println(key, "string: ", value)
		case float64:
			fmt.Println(key, "float64: ", value)
		default:
			fmt.Printf("%v\n", value)
		}
	}
}
