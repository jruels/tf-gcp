package main

import (
	"encoding/json"
	"fmt"
)

// 1 .Create a Go application
// 2. Create an Employee structure. It should have these

// members: name, role, age, insurance that
// correspond to the table columns.
type Employee struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	Age       int    `json:"age,omitempty"`
	Insurance string `json:"insurance,omitempty"`
}

// 5. Ignore data not shown in the table ==> omitempty

func main() {

	// 3. Initialize a slice of Employees
	allemployees := []Employee{
		Employee{"Bob Jones", "Manager", 55, ""},
		Employee{"Sally Wilson", "VP", 55, "No"},
		Employee{"Sam Smith", "Warehouse", 0, "Yes"},
		Employee{"Adam Freeman", "Warehouse", 55, "Yes"},
	}

	var serializedEmployees []byte

	// 4. Convert data to JSON
	serializedEmployees, mErr := json.MarshalIndent(allemployees, "", "\t")
	if mErr != nil {
		fmt.Println(mErr)
	} else {
		// 6. Display the final JSON

		//convert the byte[] to string
		contentAsString := string(serializedEmployees)

		fmt.Println(contentAsString)
	}

}
