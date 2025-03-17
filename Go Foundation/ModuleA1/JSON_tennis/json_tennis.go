package main

import "fmt"
import "encoding/json"

type TennisDoubles struct {
	Players  []string
	Nickname string `json:",omitempty"` // modifier
	W        int    `json:"Won"`
	L        int    `json:"Lost"`
	ties     int    `json:",string"` // it does  matter!
}

func main() {
	var teams = []TennisDoubles{
		{
			Players:  []string{"Chris Evert", "Pam Shriver", "Billie Jean"},
			Nickname: "Hot Shots",
			W:        7,
			L:        3,
			ties:     2,
		},
		{
			[]string{"Venus Williams", "Maria Sharapova"},
			"",
			9,
			1,
			1,
		},
	}

	jsonData, err := json.MarshalIndent(teams, "pretty", "\t")

	if err != nil {
		fmt.Println("Marshaling failed!")
		return
	}
	fmt.Println(teams)
	fmt.Printf("%s", jsonData)
}
