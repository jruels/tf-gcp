package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

/*
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}
*/

type IntStringMap map[int]string // create alias

func (m *IntStringMap) MarshalJSON() ([]byte, error) { // Marshaler interface
	ss := map[string]string{}
	for k, v := range *m {
		i := strconv.Itoa(k)
		ss[i] = v // convert to map[string]string
	}
	return json.Marshal(ss) // marshal map[string]string
}

func (m *IntStringMap) UnmarshalJSON(data []byte) error { // receiver is map[int]map
	ss := map[string]string{}
	err := json.Unmarshal(data, &ss) // get map[string]string
	if err != nil {
		return err
	}
	for k, v := range ss {
		i, err := strconv.Atoi(k) // convert string to an integer
		if err != nil {
			return err
		}
		(*m)[i] = v // add map[int]string to map slice
	}
	return nil
}

func main() {

	m := IntStringMap{4: "four", 5: "five"}
	data, err := m.MarshalJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("IntStringMap to JSON: ", string(data))

	m = IntStringMap{}

	jsonString := []byte("{\"1\": \"one\", \"2\": \"two\"}")
	m.UnmarshalJSON(jsonString)

	fmt.Printf("IntStringMap from JSON: %v\n", m)
	fmt.Println("m[1]:", m[1], "m[2]:", m[2])
}
