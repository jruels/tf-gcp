package main

import "encoding/json"
import "fmt"

type Employee struct {
	First    string
	Second   string
	Role     string
	Salary   int
	internal int
}

func main() {
	sam := Employee{"Sam", "Wilson", "President", 120000, 123}
	result, err := json.Marshal(sam)

	if err == nil {
		fmt.Printf("%s\n", result)
	} else {
		fmt.Println(err)
		return
	}

	var sam2 Employee
	err = json.Unmarshal(result, &sam2)

	if err == nil {
		fmt.Println(sam2)
	} else {
		fmt.Println(err)
	}

}
